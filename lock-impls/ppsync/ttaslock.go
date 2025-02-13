// Package ppsync provides basic mutual exclusion lock primitives.
package ppsync

import (
	"sync"
)

// A TTASLock is a mutual exclusion lock that represents a test-and-test-and-set lock.
// The zero value for a TTASLock is an unlocked mutex.
type TTASLock struct {
	state *AtomicBool
}

// Creates, initializes and returns a new TTASLock object.
func NewTTASLock() sync.Locker {
	atomicBool := NewABool(false)
	return &TTASLock{&atomicBool}
}

// Lock locks lock. If the lock is already in use, the calling goroutine
// blocks until the lock is available.
func (lock *TTASLock) Lock() {
	for {
		for lock.state.Get() {
		} // spin while some thread has lock; thread goes to next line when it sees lock is free
		if !lock.state.GetAndSet(true) { // thread attempts to acquire lock
			return // this thread succeeded to acquire key
		}
	}
}

// Unlock unlocks lock.
// It is a run-time error if lock is not locked on entry to Unlock.
//
// A locked lock is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a lock and then
// arrange for another goroutine to unlock it.
func (lock *TTASLock) Unlock() {
	lock.state.Set(false)
}
