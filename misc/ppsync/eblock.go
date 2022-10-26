// Package ppsync provides basic mutual exclusion lock primitives.
package ppsync

import (
	"math/rand"
	"sync"
	"time"
)

// An EBLock improves upon TTASLock by reducing bus traffic.
// EBLock performance is sensitive to input parameters.
type EBLock struct {
	state    *AtomicBool
	minDelay *int // ms == milliseconds
	maxDelay *int // ms
}

type Backoff struct {
	minDelay int // ms == milliseconds
	maxDelay int // ms
	limit    int // ms
}

func (lock *Backoff) backoff() {
	delay := rand.Intn(lock.limit)
	lock.limit = Min(lock.maxDelay, 2*lock.limit)       // update limit
	time.Sleep(time.Duration(delay) * time.Millisecond) // sleep for XX milliseconds
}

// Creates, initializes and returns a new BackOff object.
func NewBackoff(minDelay, maxDelay int) *Backoff {
	limit := minDelay
	return &Backoff{minDelay, maxDelay, limit}
}

func Min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}

// Creates, initializes and returns a new EBLock object.
// minDelay, maxDelay are given in milliseconds
func NewEBLock(minDelay, maxDelay int) sync.Locker {
	atomicBool := NewABool(false)
	return &EBLock{&atomicBool, &minDelay, &maxDelay}
}

// Lock locks lock. If the lock is already in use, the calling goroutine
// blocks until the lock is available.
func (lock *EBLock) Lock() {
	backoff := NewBackoff(*lock.minDelay, *lock.maxDelay)
	for {
		for lock.state.Get() {
		} // spin while some thread has lock; thread goes to next line when it sees lock is free
		if !lock.state.GetAndSet(true) { // thread attempts to acquire lock
			return // this thread succeeded to acquire key
		} else {
			backoff.backoff()
		}
	}
}

// Unlock unlocks lock.
// It is a run-time error if lock is not locked on entry to Unlock.
//
// A locked lock is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a lock and then
// arrange for another goroutine to unlock it.
func (lock *EBLock) Unlock() {
	lock.state.Set(false)
}
