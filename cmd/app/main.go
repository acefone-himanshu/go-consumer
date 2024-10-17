package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
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

	const maxWorkers = 10000
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	var messageCount uint64 = 0
	// Ticker to log the count every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Graceful shutdown handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle graceful shutdown
	go func() {
		<-sigs
		fmt.Println("\nReceived shutdown signal. Waiting for ongoing processes to complete...")
		ticker.Stop()
		reader.Close() // Stop fetching new messages
		close(sem)     // Ensure all slots in the semaphore are freed
		wg.Wait()      // Wait for all workers to finish
		os.Exit(0)
	}()

	go func() {
		for range ticker.C {
			fmt.Printf("Messages processed in last 1 second: %d\n", messageCount)
			// Reset the counter after logging
			messageCount = 0
		}
	}()

	for {
		m, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Println("Error fetching message:", err)
			break
		}

		wg.Add(1)

		// Acquire a slot in the semaphore
		sem <- struct{}{}
		messageCount++
		// Process the message in a goroutine
		go consumer.ProcessMessage(m, &wg, sem)
	}
}
