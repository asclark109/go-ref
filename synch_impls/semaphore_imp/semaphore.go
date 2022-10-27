package semaphore

import "sync"

type Semaphore struct {
	capacity  int
	mutex     *sync.Mutex
	condition *sync.Cond
}

func NewSemaphore(capacity int) *Semaphore {
	var mutex sync.Mutex
	condition := sync.NewCond(&mutex)
	return &Semaphore{capacity, &mutex, condition}
}

func (s *Semaphore) Up() {
	s.mutex.Lock()
	s.capacity++
	s.condition.Broadcast()
	s.mutex.Unlock()
}

func (s *Semaphore) Down() {
	s.mutex.Lock()
	for s.capacity == 0 {
		s.condition.Wait()
	}
	s.capacity--
	s.mutex.Unlock()
}
