// Package lock provides an implementation of a read-write lock
// that uses condition variables and mutexes.
package lock

import (
	"proj1/semaphore"
)

const READ_LIMIT = 32

type RWLocker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

// RWLock is my implementation of a Read-Write Lock
type RWLock struct {
	read_sema        *semaphore.Semaphore
	write_sema_mutex *semaphore.Semaphore
}

// NewRWMutex creates and returns a new RWMutex value given its body and timestamp
func NewRWLock() *RWLock {
	read_sema := semaphore.NewSemaphore(READ_LIMIT)
	write_sema_mutex := semaphore.NewSemaphore(1)
	return &RWLock{read_sema, write_sema_mutex}
}

// Lock() will lock rw. If the lock is already in use,
// the calling goroutine blocks until the mutex is available.
func (rw *RWLock) Lock() {
	rw.write_sema_mutex.Down()
}

// Unlock() will unlock rw. It is a run-time error if rw
// is not locked on entry to Unlock. A locked Mutex is not
// associated with a particular goroutine. It is allowed for one
// goroutine to lock a Mutex and then arrange for another
// goroutine to unlock it.
func (rw *RWLock) Unlock() {
	rw.write_sema_mutex.Up()
}

// RLock() locks rw for reading.
// It should not be used for recursive read locking; a blocked
// Lock call excludes new readers from acquiring the lock.
func (rw *RWLock) RLock() {
	rw.read_sema.Down()
}

// RUnlock() undoes a single RLock() call; it does not affect other
// simultaneous readers. It is a run-time error if rw is not
// locked for reading on entry to RUnlock.
func (rw *RWLock) RUnlock() {
	rw.read_sema.Up()
}
