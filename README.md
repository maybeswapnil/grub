# Worker Pool in Go

The "grub" package in the provided Go code defines a flexible and concurrent worker pool for processing jobs. The package introduces the WorkerPool type, allowing users to efficiently distribute and execute a collection of jobs across a specified number of workers. The design encapsulates individual workers in the worker type, each running concurrently in a goroutine to process jobs from a shared job queue. The package ensures thread safety using a mutex and provides a clean interface for users to add jobs to the pool. The Wait method plays a crucial role in managing the pool's lifecycle, closing the job queue and waiting for all jobs to complete before allowing the program to proceed. Overall, this implementation facilitates parallel job processing, making it suitable for scenarios where concurrent execution of tasks is beneficial, such as in concurrent data processing or other parallelizable workloads.

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
	workerpool "github.com/maybeswapnil/grub"
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
	pool := workerpool.NewWorkerPool(myBatchProcessor, 10)
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

This Go code defines a simple worker pool package named "grub" that enables concurrent processing of jobs using a specified number of workers. Below is a breakdown of the components and their functionalities:

**Job Interface:**
The `Job` interface is an empty interface, indicating that any type can be used as a job.

**BatchProcessorFunc Type:**
The `BatchProcessorFunc` type is a function signature for processing a batch of jobs.

**WorkerPool Struct:**
Represents a pool of workers capable of processing jobs concurrently.
- Fields:
  - `workers`: A slice of worker instances.
  - `jobQueue`: A buffered channel for queuing jobs.
  - `batchProcessor`: The function that processes a batch of jobs.
  - `wg`: A WaitGroup to wait for all jobs to complete.
  - `mux`: A mutex to synchronize access to shared data.
  - `closed`: A flag indicating whether the pool is closed or not.

**Worker Struct:**
Represents an individual worker.
- Fields:
  - `id`: Unique identifier for the worker.
  - `workerPool`: Reference to the parent worker pool.

**NewWorkerPool Function:**
Initializes a new worker pool.
- Parameters:
  - `batchProcessor`: The function that processes a batch of jobs.
  - `numWorkers`: The number of workers in the pool.
- Creates worker instances, initializes the pool, and starts each worker.

**start Method (Worker):**
Initiates a worker to process jobs.
- A goroutine is spawned for each worker, continuously processing jobs from the job queue.

**AddJob Method (WorkerPool):**
Adds a job to the worker pool.
- If the pool is closed, it returns an error; otherwise, it adds the job to the queue and increments the WaitGroup.

**Wait Method (WorkerPool):**
Waits for all jobs to complete.
- Closes the job queue and sets the pool to closed, preventing further job additions.
- Uses a mutex to ensure thread safety.

Overall, this worker pool implementation provides a simple way to process a batch of jobs concurrently using a specified number of workers. Users can add jobs to the pool, and the pool will manage the execution of these jobs efficiently. The `Wait` method is crucial for ensuring that the program waits until all jobs are completed before proceeding.

### Contributing
Feel free to contribute by opening issues or submitting pull requests. Bug reports, feature requests, and feedback are highly encouraged.

### License
This project is licensed under the MIT License - see the LICENSE file for details.

Feel free to customize it further based on the specific details of your package and any additional information you'd like to provide.
