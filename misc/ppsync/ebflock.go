// Package ppsync provides basic mutual exclusion lock primitives.
package ppsync

import (
	"hw3/utils"
	"sync"
)

// An EBFLock improves upon TTASLock by reducing bus traffic.
// EBFLock performance is sensitive to input parameters.
type EBFLock struct {
	state          *AtomicBool
	minDelay       *int // ms == milliseconds
	maxDelay       *int // ms
	goroutineQueue *utils.GoroutineQueue
}

// Creates, initializes and returns a new EBFLock object.
// minDelay, maxDelay are given in milliseconds
func NewEBFLock(minDelay, maxDelay int) sync.Locker {
	atomicBool := NewABool(false)
	goroutineQueue := utils.NewGoroutineQueue()
	return &EBFLock{&atomicBool, &minDelay, &maxDelay, goroutineQueue}
}

// Lock locks lock. If the lock is already in use, the calling goroutine
// blocks until the lock is available.
func (lock *EBFLock) Lock() {
	myTicket := lock.goroutineQueue.EnqueueAndIssueLabel() // this goroutine issued a ticket number
	backoff := NewBackoff(*lock.minDelay, *lock.maxDelay)
	for {
		for lock.state.Get() {
		} // spin while some thread has lock; thread goes to next line when it sees lock is free
		if myTicket == lock.goroutineQueue.Top() { // if it is this goroutine's turn to go
			if !lock.state.GetAndSet(true) { // thread attempts to acquire lock
				return // this thread succeeded to acquire key
			} // this should always succeed I think
		} else {
			backoff.backoff() // backoff if it's not this goroutine's turn
		}
	}
}

// Unlock unlocks lock.
// It is a run-time error if lock is not locked on entry to Unlock.
//
// A locked lock is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a lock and then
// arrange for another goroutine to unlock it.
func (lock *EBFLock) Unlock() {
	_ = lock.goroutineQueue.Dequeue()
	lock.state.Set(false)
}
