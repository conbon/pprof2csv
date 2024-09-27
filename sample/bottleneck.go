package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // Import for enabling pprof profiling
	"runtime"
	"time"
)

func cpuIntensiveTask() {
	// fmt.Println("Starting CPU-intensive task...")
	for i := 0; i < 1000000000; i++ {
		_ = i * i // Waste CPU cycles
	}
	// fmt.Println("Finished CPU-intensive task.")
}

func memoryIntensiveTask() {
	// fmt.Println("Starting memory-intensive task...")
	leak := make([][]byte, 0)
	for i := 0; i < 1000; i++ {
		block := make([]byte, 1024*1024) // Allocate 1 MB in each iteration
		leak = append(leak, block)
		time.Sleep(10 * time.Millisecond) // Simulate some delay
	}
	// fmt.Println("Finished memory-intensive task. Leaking memory...")
}

func main() {
	// Start HTTP server for pprof
	go func() {
		// fmt.Println("Starting pprof server on :6111")
		if err := http.ListenAndServe(":6111", nil); err != nil {
			fmt.Println("Error starting pprof server:", err)
		}
	}()

	// Run CPU and memory intensive tasks
	go cpuIntensiveTask()
	go memoryIntensiveTask()

	// Prevent program from exiting
	runtime.Gosched() // Yield CPU
	select {}         // Block indefinitely
}
