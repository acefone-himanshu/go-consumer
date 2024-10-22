package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	config "webhook-consumer/internal/config"
	consumer "webhook-consumer/internal/consumer"
)

func main() {

	reader := config.GetKafkaReader()
	defer reader.Close()

	fmt.Println("start consuming ... !!")

	var wg sync.WaitGroup

	const maxWorkers = 2
	sem := make(chan struct{}, maxWorkers)

	var messageCount uint64 = 0

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
		fmt.Println("\nReceived shutdown signal. Waiting for ongoing processes to complete...")
		cancel()
	}()

	// Goroutine to log message count
	go func() {
		for range ticker.C {
			goroutines := runtime.NumGoroutine()
			fmt.Printf("Messages processed in last 1 second: %d, Total: %d\n", messageCount, goroutines)
			messageCount = 0
		}
	}()

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
				log.Println("Error fetching message:", err)
				continue
			}

			messageCount++
			sem <- struct{}{}
			wg.Add(1)

			go consumer.ProcessMessage(&msg, &wg, sem)

		}
	}

}
