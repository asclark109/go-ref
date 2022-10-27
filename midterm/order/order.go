package main

// Write a program that asks the user (via the command line)
// to specify a certain n number of goroutines to spawn.
// Each goroutine will need to be given a unique integer
// identification value starting from 0.

// The program will spawn the n goroutines to execute the same
// function named worker(goID int). Within the worker function,
// each goroutine will synchronously increment a shared counter
// variable and then immediately print its unique id and the value of the updated shared variable. Each goroutine only increments the shared variable once and then exits the function. All goroutines must increment the shared counter in order based on their id number. This means forcing the goroutine with goID = 0 to run first, followed by the thread with goID = 1, and so on.

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

func main() {

	// CHECK: number of args (not required)
	arguments := os.Args[1:]
	if !(len(arguments) == 1) {
		fmt.Println("wrong number of args")
		return
	}

	// assume args are ints
	numRoutinesStr := arguments[0]
	numRoutines, _ := strconv.Atoi(numRoutinesStr)

	// solve problem instance
	group := &sync.WaitGroup{}
	var sharedVar int64
	context := &SharedContext{group, &sharedVar}
	forkJoin(context, numRoutines)

}

func forkJoin(context *SharedContext, numRoutines int) {

	for i := 0; i < numRoutines-1; i++ {
		context.group.Add(1)
		go worker(i, context)
	}
	context.group.Add(1)
	worker(numRoutines-1, context)
	context.group.Wait()

}

func worker(goID int, context *SharedContext) {

	myValue := int64(goID)
	newValue := int64(goID + 1)

	for !atomic.CompareAndSwapInt64(context.shareVar, myValue, newValue) { // all threads take turns updating the sharedVar
	}

	// update completed
	fmt.Printf("Gorountine id = %v, sharedVar = %v\n", myValue, newValue)

	// NOTE! the updating occurs correctly in order, the printing may not
	// be in sync, but the updates are, which is the important part

	context.group.Done()

}

type SharedContext struct {
	group    *sync.WaitGroup
	shareVar *int64
}
