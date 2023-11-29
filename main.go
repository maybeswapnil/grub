package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
	grub "grub/src"
)

type MyJob struct {
	ID    int
	Value string
	key   string
}

func printSystemInfo() {
	numGoroutine := runtime.NumGoroutine()
	fmt.Printf("Number of Goroutines: %d\n", numGoroutine)
}

func myBatchProcessor(job workerpool.Job) {
	myJob := job.(MyJob)
	log.Println("Processing job:", myJob.ID)
	time.Sleep(time.Second)
	log.Println("Processed job:", myJob.ID, myJob.key)
}

func main() {
	// Create a worker pool with custom job type
	pool := workerpool.NewWorkerPool(myBatchProcessor, 5, 10)
	start := time.Now()

	printSystemInfo()

	// Add custom jobs to the pool
	for i := 1; i <= 50; i++ {
		myJob := MyJob{ID: i, Value: "Custom Value", key: "key"}
		pool.AddJob(myJob)
	}

	// Wait for all jobs to complete
	pool.Wait()

	elapsed := time.Since(start)
	printSystemInfo()
	log.Printf("Package took %s", elapsed)
	log.Println("All jobs completed.")
}