package main

import (
	"fmt"
	"hw2/utils"
	"math"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

func main() {

	// CHECK: number of args (not required)
	arguments := os.Args[1:]
	if !(len(arguments) == 2) {
		fmt.Println(genUsageStr())
		return
	}

	// CHECK: args are ints (not required)
	intervalsStr := arguments[0]
	threadsStr := arguments[1]
	if !(utils.IsPostiveInt(intervalsStr) && utils.IsPostiveInt(threadsStr)) {
		fmt.Println(genUsageStr())
		return
	}
	intervals, _ := strconv.Atoi(intervalsStr)
	threads, _ := strconv.Atoi(threadsStr)

	// solve problem instance
	group := &sync.WaitGroup{}
	var flag int64
	var piEstimate float64
	context := &sharedContext{&piEstimate, group, &flag}
	forkJoinSolvePi(context, threads, intervals)
	fmt.Printf("%.10f", *context.piEstimate)

}

// genUsageStr returns the basic command-line usage msg as a string
func genUsageStr() string {
	return "Usage: pi interval threads\n" +
		"    interval = the number of iterations to perform\n" +
		"    threads = the number of threads (i.e., goroutines to spawn)\n"
}

// sequentialSolvePi accepts two ints i,j denoting the term indices in the
// the infinite series sum for the arc_tan_estimation method; it computes all
// terms from i to j (inclusive), sums them, multiplies result by 4, and returns
// the result. NOTE: this is essentially the sequential algorithm (i.e.
// solveProblem()) decomposed into a solveTask(). Clarifying example:
// sequentialSolvePi(0,99) == sequentialSolvePi(0,49) + sequentialSolvePi(49,99)
// == 3.14159...
func sequentialSolvePi(termStart, termEnd int) float64 {
	var piFragment float64
	var numerator, denominator float64
	var x float64

	for i := termStart; i <= termEnd; i++ {
		x = float64(i)
		numerator = math.Pow(-1, x) * math.Pow(1, 2*x+1)
		denominator = 2*x + 1
		piFragment += numerator / denominator
	}

	return piFragment * 4
}

// forkJoinSolvePi accepts a context pertaining to a problem instance
// for the estimation of pi; threads is the number of threads to use
// for solving the problem; iterations is the number of terms to compute
// and use in the estimation of pi (i.e. more iterations for higher
// precision. This function mutates piEstimation, a field within the
// sharedContent that represents the estimation of pi. ASSUMES:
// context.piEstimate == 0 initially, threads and iterations >= 1
func forkJoinSolvePi(context *sharedContext, threads, iterations int) {
	// direct solve
	if iterations <= 1000 || threads == 1 {
		*context.piEstimate = sequentialSolvePi(0, iterations-1)
		return
	}
	// fork-join solve
	numTermsPerThread := iterations / threads                       // generic workload (int division)
	numTermsBigThread := iterations - numTermsPerThread*(threads-1) // workload for thread doing extra work
	var strt, end int
	for i := 0; i < threads-1; i++ { // create (threads-1) other go routines
		// determine term indices to include in this thread
		strt = i * numTermsPerThread
		end = strt + (numTermsPerThread - 1)
		context.group.Add(1)
		go worker(context, strt, end)
	}
	// assign to main thread the remainder load
	context.group.Add(1)
	worker(context, end+1, (end+1)+numTermsBigThread-1)
	context.group.Wait()
}

// shared memory struct for a problem instance of estimating pi
type sharedContext struct {
	piEstimate *float64
	group      *sync.WaitGroup
	flag       *int64
}

// subroutine for threads (goroutines); applies atomic operations when read/writing
// the current estimated value for pi in the sharedcontext.
func worker(ctx *sharedContext, termStart int, termEnd int) {

	piFragment := sequentialSolvePi(termStart, termEnd) // all threads do work independently here

	for !atomic.CompareAndSwapInt64(ctx.flag, 0, 1) { // all threads take turns updating the estimate of Pi
	}

	*ctx.piEstimate += piFragment

	atomic.StoreInt64(ctx.flag, 0)

	ctx.group.Done()

}
