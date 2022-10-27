package main

import (
	"fmt"
	"hw3/ppsync"
	"hw3/utils"
	"os"
	"strconv"
	"sync"
	"time"
)

// acceptable program actions
const (
	TAS  = "tas"
	TTAS = "ttas"
	EB   = "eb"
	EBF  = "ebf"
	SYNC = "sync"
)

var ValidCommands = [5]string{TAS, TTAS, EB, EBF, SYNC}

// printUsage prints the usage statement for the program
func printUsage() {
	commands := "command:\n\ttas - Runs the program using a TASLock lock.\n" +
		"\tttas - Runs the program using a TTASLock lock.\n" +
		"\teb MIN_DELAY MAX_DELAY - Runs the promgram using a EBLock lock. MIN_DELAY and MAX_DELAY are both *STATE YOUR NUMERIC TYPE* values " +
		"that represent time in milliseconds.\n" +
		"\tebf - Runs the program using an ALock lock.\n" +
		"\tsync - Runs the syc.Mutex lock.\n"

	incrementAmount := "incrementAmount: the number of times the goroutines should increment the counter\n"
	threads := "threads: the number of threads to spawn\n"

	fmt.Printf("Usage: experiment command incrementAmount threads\n" + commands + incrementAmount + threads)
}

func main() {

	// [1] VALIDATE ARGUMENTS
	// [2] RUN PROGRAM

	// [1]
	// CHECK: number of args is at least 1 (name of command)
	arguments := os.Args[1:]
	// fmt.Println(arguments)
	if !(len(arguments) >= 1) {
		fmt.Println("did not provide enough arguments")
		printUsage()
		return
	}

	// CHECK: first arg is a recognized command
	hasValidCmd := false
	command := arguments[0]
	for _, validCmd := range ValidCommands {
		if command == validCmd {
			hasValidCmd = true
			break
		}
	}
	if !hasValidCmd {
		fmt.Println("did not recognize command")
		printUsage()
		return
	}

	// CHECK: valid remaining args

	var minDelay, maxDelay, incrementAmount, threads int

	if command == TAS || command == TTAS || command == SYNC {

		// CHECK: exactly two more args
		if len(arguments)-1 != 2 {
			fmt.Println("wrong num args")
			printUsage()
			return
		}

		// CHECK: last two args are ints
		incrementAmountStr := arguments[1]
		threadsStr := arguments[2]
		if !(utils.IsPostiveInt(incrementAmountStr) && utils.IsPostiveInt(threadsStr)) {
			fmt.Println("did not get positive ints for all remaining args")
			printUsage()
			return
		}
		incrementAmount, _ = strconv.Atoi(incrementAmountStr)
		threads, _ = strconv.Atoi(threadsStr)

	} else if command == EB || command == EBF {

		// CHECK: exactly four more args
		if len(arguments)-1 != 4 {
			fmt.Println("wrong num args")
			printUsage()
			return
		}

		// CHECK: last four args are ints
		// fmt.Println("args: ", arguments)
		minDelayStr := arguments[1]
		maxDelayStr := arguments[2]
		incrementAmountStr := arguments[3]
		threadsStr := arguments[4]

		if !(utils.IsPostiveInt(minDelayStr) && utils.IsPostiveInt(maxDelayStr) &&
			utils.IsPostiveInt(incrementAmountStr) && utils.IsPostiveInt(threadsStr)) {
			fmt.Println("did not get positive ints for all remaining args")
			printUsage()
			return
		}
		minDelay, _ = strconv.Atoi(minDelayStr)
		maxDelay, _ = strconv.Atoi(maxDelayStr)
		incrementAmount, _ = strconv.Atoi(incrementAmountStr)
		threads, _ = strconv.Atoi(threadsStr)
	}

	// fmt.Println("got good args.")
	// fmt.Println(minDelay, maxDelay, incrementAmount, threads)

	// [2] create lock based on program args. then run counter program with that lock

	var mutex sync.Locker

	switch {
	case command == TAS:
		mutex = ppsync.NewTASLock()
	case command == TTAS:
		mutex = ppsync.NewTTASLock()
	case command == EB:
		mutex = ppsync.NewEBLock(minDelay, maxDelay)
	case command == EBF:
		mutex = ppsync.NewEBFLock(minDelay, maxDelay)
	case command == SYNC:
		mutex = &sync.Mutex{}
	}

	// fmt.Println("calling runProgram with: ", incrementAmount, threads, &mutex)
	result := runProgram(incrementAmount, threads, &mutex)
	fmt.Println(result)

}

// below code borrowed from hw2/waitgroup_test.go

func counter(gId int, sharedCount *int, amount int, delayAmount time.Duration, group *sync.WaitGroup, mutex *sync.Locker) {

	//Critical section use synchronization before proceeding
	(*mutex).Lock()
	for i := 0; i < amount; i++ {
		*sharedCount = *sharedCount + 1
		if delayAmount > 0 {
			time.Sleep(delayAmount)
		}
	}
	(*mutex).Unlock()
	group.Done()
}

// below code borrowed from hw2/waitgroup_test.go

func runProgram(incrementAmount, threads int, mutex *sync.Locker) int {

	group := &sync.WaitGroup{}
	var sharedCount int

	for i := 0; i < threads; i++ {
		group.Add(1)
		go counter(i, &sharedCount, incrementAmount, 0, group, mutex)
	}
	group.Wait()

	// fmt.Printf("got:%v \texpected:%v", sharedCount, incrementAmount*threads)
	return sharedCount
}
