# Worker Pool in Go

A simple worker pool implementation in Go that allows concurrent processing of jobs using goroutines and channels.

## Installation

```bash
go get -u github.com/maybeswapnil/grub
```

## Usage
```go
package main

import (
	"log"
	"time"
	workerpool "workergroup/src"
)

type MyJob struct {
	ID    int
	Value string
	key   string
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

	// Add custom jobs to the pool
	for i := 1; i <= 50; i++ {
		myJob := MyJob{ID: i, Value: "Custom Value", key: "key"}
		pool.AddJob(myJob)
	}
	// Wait for all jobs to complete
	pool.Wait()
	elapsed := time.Since(start)
	log.Printf("Package took %s", elapsed)
	log.Println("All jobs completed.")
}
```

### API
workerpool.Job Represents a unit of work. Users can define their own types that implement this interface.

type BatchProcessorFunc func(Job) Represents a function that processes a batch of jobs. Users should implement their own custom batch processor function.

type WorkerPool struct Represents a pool of workers.

NewWorkerPool(batchProcessor BatchProcessorFunc, numWorkers, bufferSize int) *WorkerPool: Initializes a new worker pool.

AddJob(job Job) error: Adds a job to the worker pool.

Wait(): Waits for all jobs to complete.

### Contributing
Feel free to contribute by opening issues or submitting pull requests. Bug reports, feature requests, and feedback are highly encouraged.

### License
This project is licensed under the MIT License - see the LICENSE file for details.

Feel free to customize it further based on the specific details of your package and any additional information you'd like to provide.
