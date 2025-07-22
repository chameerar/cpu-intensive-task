package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
)

var interrupted atomic.Bool

func cpuIntensiveTask(duration time.Duration) {
	end := time.Now().Add(duration)
	iteration := 1
	iterStart := time.Now()

	for time.Now().Before(end) && !interrupted.Load() {

		// Busy loop for ~400ms
		busyStart := time.Now()
		for time.Since(busyStart) < 400*time.Millisecond {
			_ = math.Sqrt(12345.6789)
		}

		// Sleep for ~600ms
		time.Sleep(600 * time.Millisecond)

		duration := time.Since(iterStart)
		fmt.Printf("Iteration %d complete (total elapsed: %s)\n", iteration, duration.Truncate(time.Millisecond))
		iteration++
	}

	fmt.Println("CPU task stopped.")
}

func main() {
	runtime.GOMAXPROCS(1)

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\nInterrupt received. Shutting down gracefully...")
		interrupted.Store(true)
	}()

	fmt.Println("Starting moderately CPU-intensive task for 5 minutes...")
	cpuIntensiveTask(5 * time.Minute)
	fmt.Println("Done.")
}
