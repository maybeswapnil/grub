package grub

import (
	"errors"
	"sync"
)

// Job represents a unit of work
type Job interface{}

// BatchProcessorFunc represents a function that processes a job
type BatchProcessorFunc func(Job)

// WorkerPool represents a pool of workers
type WorkerPool struct {
	workers        []*worker
	jobQueue       chan Job
	batchProcessor BatchProcessorFunc
	wg             sync.WaitGroup
	mux            sync.Mutex
	closed         bool
}

// worker represents an individual worker
type worker struct {
	id         int
	workerPool *WorkerPool
}

// NewWorkerPool initializes a new worker pool with a job generator function
// Parameters:
// - batchProcessor: The function that processes a batch of jobs.
// - numWorkers: The number of workers in the pool.
// - bufferSize: The size of the job queue buffer.
// Returns:
// - A pointer to the initialized WorkerPool.
func NewWorkerPool(batchProcessor BatchProcessorFunc, numWorkers, bufferSize int) *WorkerPool {
	pool := &WorkerPool{
		jobQueue:       make(chan Job, bufferSize),
		batchProcessor: batchProcessor,
	}

	for i := 1; i <= numWorkers; i++ {
		worker := &worker{
			id:         i,
			workerPool: pool,
		}
		pool.workers = append(pool.workers, worker)
		worker.start()
	}

	return pool
}

// start initiates a worker to process jobs
func (w *worker) start() {
	go func() {
		for job := range w.workerPool.jobQueue {
			w.workerPool.batchProcessor(job)
			w.workerPool.wg.Done()
		}
	}()
}

// AddJob adds a job to the worker pool
// Parameters:
// - job: The job to be added to the pool.
// Returns:
// - An error if the pool is closed, nil otherwise.
func (pool *WorkerPool) AddJob(job Job) error {
	pool.mux.Lock()
	defer pool.mux.Unlock()

	if pool.closed {
		return errors.New("job not added, pool is closed")
	}

	pool.wg.Add(1)
	pool.jobQueue <- job
	return nil
}

// Wait waits for all jobs to complete
// Closes the job queue and sets the pool to closed to prevent further job additions.
func (pool *WorkerPool) Wait() {
	pool.mux.Lock()
	defer pool.mux.Unlock()

	if !pool.closed {
		close(pool.jobQueue)
		pool.closed = true
	}

	pool.wg.Wait()
}