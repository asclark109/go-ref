package waitgroup

import (
	"sync/atomic"
)

// // From the Go Documentation: A WaitGroup waits for a collection of
// // goroutines to finish. The main goroutine calls Add to set the number
// // of goroutines to wait for.  Then each of the goroutines runs and calls
// // Done when finished.  At the same time, Wait can be used to block until
// // all goroutines have finished.
type WaitGroup interface {
	Add(amount uint)
	Done()
	Wait()
}

// NewWaitGroup returns a instance of a waitgroup
// This instance must be a pointer and should not
// be copied after creation.
func NewWaitGroup() WaitGroup {
	flag := int64(0)
	runningThreads := int64(0)
	return &ConcreteWaitGroup{&flag, &runningThreads}
}

// ConcreteWaitGroup is a struct implementing the WaitGroup interface
type ConcreteWaitGroup struct {
	runningThreads *int64
	flag           *int64
}

// Add signals to this ConcreteWaitGroup that a certain amount
// of new goroutines are to be spawned.
func (wg *ConcreteWaitGroup) Add(amount uint) {

	for !atomic.CompareAndSwapInt64(wg.flag, 0, 1) {
	}

	*wg.runningThreads += int64(amount)

	atomic.StoreInt64(wg.flag, 0)

}

// Done signals to this ConcreteWaitGroup that a goroutine
// belonging to this ConcreteWaitGroup has finished.
func (wg *ConcreteWaitGroup) Done() {

	for !atomic.CompareAndSwapInt64(wg.flag, 0, 1) {
	}

	*wg.runningThreads += int64(-1)

	atomic.StoreInt64(wg.flag, 0)

}

// Wait causes this ConcreteWaitGroup to block the calling scope
// (initiate a halt in the calling scope) until all goroutines
// belonging to this ConcreteWaitGroup have completed their tasks.
func (wg *ConcreteWaitGroup) Wait() {
	for !(atomic.LoadInt64(wg.runningThreads) == 0) {
	}
}
