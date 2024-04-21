package main

// Queue is a simple implementation of a concurrent queue using a channel
type Queue chan Car

// NewQueue creates a new Queue with an unbuffered channel (blocking)
func NewQueue() Queue {
	return make(chan Car)
}

// Enqueue adds a Job to the back of the queue
func (q Queue) Enqueue(job Car) {
	q <- job // Blocking send until a consumer is ready
}

// Dequeue removes and returns the Job at the front of the queue
func (q Queue) Dequeue() (Car, bool) {
	var job Car
	ok := true

	// Non-blocking receive. If no job is available, ok will be false.
	select {
	case job = <-q:
	default:
		ok = false
	}

	return job, ok
}
