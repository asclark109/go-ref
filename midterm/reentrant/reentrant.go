package reentrant

import "sync"

type ReentrantLock struct {
	mutex     *sync.Mutex
	cond      *sync.Cond // use conditional var
	wgContext *sync.WaitGroup
	ownerGoId *int // go routine that currently holds lock
	holdCount *int // how many times current go routine holds the lock
}

func NewReentrantLock() *ReentrantLock {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	condVar := sync.NewCond(&mutex)
	var ownerGoId int
	var holdCount int
	return &ReentrantLock{&mutex, condVar, &wg, &ownerGoId, &holdCount}
}

func (lock *ReentrantLock) Lock(id int) {

	lock.mutex.Lock()

	if *lock.ownerGoId == id { // groutine asking to lock currently holds lock
		*lock.holdCount += 1
		lock.mutex.Unlock()
		return
	}

	// otherwise this goroutine does NOT currently hold lock, must go to sleep
	for *lock.holdCount != 0 {
		lock.cond.Wait()
	}

	// now this goroutine owns the lock
	*lock.ownerGoId = id
	*lock.holdCount = 1

	lock.mutex.Unlock()
}

func (lock *ReentrantLock) Unlock(id int) {

	lock.mutex.Lock()

	if *lock.ownerGoId != id || *lock.holdCount == 0 { // groutine asking to lock currently holds lock
		// this is an illegal call to unlock; for now just return
		lock.mutex.Unlock()
		return
	}

	// otherwise this is a valid scenario to unlock.
	// this goroutine holds lock and the holdCount > 0
	*lock.holdCount -= 1

	// if this goroutine has released all its locks (holdCount == 0),
	// then we can wake up a different goroutine waiting to acquire the lock
	if *lock.holdCount == 0 {
		lock.cond.Signal() // wake up one goroutine
	}

	lock.mutex.Unlock()

}
