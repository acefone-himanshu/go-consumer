package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	config "webhook-consumer/internal/config"
	consumer "webhook-consumer/internal/consumer"
	"webhook-consumer/internal/logger"
)

func main() {

	reader := config.GetKafkaReader()
	defer reader.Close()

	var wg sync.WaitGroup

	const maxWorkers = 2000
	sem := make(chan struct{}, maxWorkers)

	var messageCount uint64 = 0

	intialGoRoutines := runtime.NumGoroutine()
	startTime := time.Now()

	logger.Logger.Info(intialGoRoutines)
	// Ticker to log the count every 1 second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Goroutine to handle graceful shutdown
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		logger.Logger.Warn("Received shutdown signal. Waiting for ongoing processes to complete...")
		cancel()
	}()

	// Goroutine to log message count
	go func() {
		for range ticker.C {
			goroutines := runtime.NumGoroutine()
			if goroutines == intialGoRoutines {
				endTime := time.Now()
				logger.Logger.Info("Total time taken: ", endTime.Sub(startTime))
				os.Exit(0)
			}
			message := fmt.Sprintf("Messages processed in last 1 second: %d, Total: %d\n", messageCount, goroutines)
			logger.Logger.Info(message)
			messageCount = 0
		}
	}()

	logger.Logger.Info("start consuming ... !!")

	for {
		select {
		case <-ctx.Done():
			close(sem)
			wg.Wait()
			return

		default:

			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				if err == context.Canceled {
					return
				}
				logger.Logger.Error("Error fetching message:", err)
				continue
			}

			messageCount++
			sem <- struct{}{}
			wg.Add(1)

			go consumer.ProcessMessage(&msg, &wg, sem)

		}
	}

}
