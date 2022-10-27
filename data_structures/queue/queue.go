package queue

import "sync/atomic"

//// QUEUE ////
// Note: I borrowed some code I saw online for a basic implementation of a FIFO queue.
// I refactored it, added to it, and made it into a FIFO queue for goroutines.
// That is, the EBFLock owns a single internal queue with which it assigns ticket numbers
// to goroutines that ask to acquire the lock.
// Credit to Sub6Resources: https://gist.github.com/Sub6Resources/f93447ed9b288219df6992a487e62e57

// The implementation of the FIFO queue uses an underlying slice.
// a goroutine will ask the lock for the lock.
// the lock will acknowledge the request by issuing a ticket for the goroutine, enqueuing that ticket
// number into its internal queue.

type GoroutineQueue struct {
	queue []int
	label int    // acts as ticket number generator for priority
	flag  *int64 // prevents race conditions when goroutines want to enqueue themselves
}

func (q *GoroutineQueue) EnqueueAndIssueLabel() int {

	// CRITICAL SECTION START (threads take turns enqueuing themselves)
	for !atomic.CompareAndSwapInt64(q.flag, 0, 1) {
	}

	issuedLabel := q.label
	q.queue = append(q.queue, issuedLabel)
	q.label++

	atomic.StoreInt64(q.flag, 0)
	// CRITICAL SECTION END (threads take turns enqueuing themselves)

	return issuedLabel // this is the issued "label" for the int we just added

}

func (q *GoroutineQueue) Top() int {
	return q.queue[0]
}

func (q *GoroutineQueue) Dequeue() int {
	// CRITICAL SECTION START
	for !atomic.CompareAndSwapInt64(q.flag, 0, 1) {
	}

	temp := q.queue[0]
	q.queue = q.queue[1:]

	atomic.StoreInt64(q.flag, 0)
	// CRITICAL SECTION END

	return temp
}

func (q *GoroutineQueue) Empty() bool {
	return len(q.queue) == 0
}

func NewGoroutineQueue() *GoroutineQueue {
	flag := int64(0)
	return &GoroutineQueue{make([]int, 0), 0, &flag}
}

// EXAMPLE USAGE
// fmt.Println("Queues")
// queue := NewGoroutineQueue()
// _ = queue.EnqueueAndIssueLabel()
// _ = queue.EnqueueAndIssueLabel()
// fmt.Println("Top:", queue.Top())
// fmt.Println("Dequeued", queue.Dequeue())
// fmt.Println("New Top:", queue.Top())
// fmt.Println("Empty?", queue.Empty())
// fmt.Println("Dequeued", queue.Dequeue())
// fmt.Println("Empty now?", queue.Empty())

//// MISC ////
