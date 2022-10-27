package reentrant

import (
	"fmt"
	"sync"
	"testing"
)

func recursiveFun(n, goID int, counter *int64, lock *ReentrantLock) {

	lock.Lock(goID)
	if n != 0 {
		for i := 0; i < 1e6; i++ {
			*counter = *counter + 1
		}
		recursiveFun(n-1, goID, counter, lock)
	}
	lock.Unlock(goID)

}

func runExperiment(goID int, counter *int64, wg *sync.WaitGroup, lock *ReentrantLock) {
	recursiveFun(3, goID, counter, lock)
	wg.Done()
}
func runTest(n int, t *testing.T) {

	var counter int64
	var wg sync.WaitGroup

	lock := NewReentrantLock()

	for goID := 0; goID < n; goID++ {
		wg.Add(1)
		go runExperiment(goID, &counter, &wg, lock)
	}
	wg.Wait()
	if counter != int64(n*3*1e6) {
		t.Errorf("Expected = %v Got = %v", int64(n*3*1e6), counter)
	}
}

func TestLock(t *testing.T) {

	var tests = []struct {
		threads int
	}{
		{2},
		{3},
		{5},
		{8},
	}
	for num, test := range tests {
		testname := fmt.Sprintf("T=%v", num)
		t.Run(testname, func(t *testing.T) {
			runTest(test.threads, t)
		})
	}
}
