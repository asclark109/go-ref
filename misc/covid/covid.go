package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const zipcodeCol = 0
const weekStart = 2
const casesWeek = 4
const testsWeek = 8
const deathsWeek = 14
const MMDDYYYY = "01/02/2006" // format of date

func main() {

	// parse arguments
	modeArgPtr := flag.String("mode", "sequential", "the mode, to be one of [sequential, static, map]")
	threadsArgPtr := flag.Int("threads", runtime.NumCPU(), "number of threads to use when solving")
	flag.Parse()

	// terminate early if program did not receive ZIPCODE YEAR MONTH
	otherArgs := flag.Args()
	if len(otherArgs) < 3 {
		fmt.Println(genUsageStr())
		return
	}
	zipcode := otherArgs[0]
	month, _ := strconv.Atoi(otherArgs[1]) // can assume int provided
	year, _ := strconv.Atoi(otherArgs[2])  // can assume int provided

	// assemble filenames to process
	filenames := make([]string, 500)
	for i := 1; i <= 500; i++ {
		filename := "data/covid_" + strconv.Itoa(i) + ".csv"
		filenames[i-1] = filename
	}

	// interpret arguments and execute program
	var cases, tests, deaths int
	group := &sync.WaitGroup{}
	var flag int64
	recordMap := make(map[string][3]int)
	context := &sharedContext{&cases, &tests, &deaths, &zipcode, &month, &year, &filenames, &recordMap, group, &flag}
	if *modeArgPtr == "sequential" {
		forkJoinSolve(context, 1) // reduced code duplication. forkJoinSolve with 1 thread same as sequentialSolve
	} else if *modeArgPtr == "static" {
		forkJoinSolve(context, *threadsArgPtr)
	} else if *modeArgPtr == "map" {
		mapSolve(context)
	} else { // did not recognize mode
		fmt.Println(genUsageStr())
		fmt.Println(*threadsArgPtr)
		return
	}

	displayResult(cases, tests, deaths)
}

// displayResult prints out the solution in desired format
func displayResult(cases, tests, deaths int) {
	casesStr := strconv.Itoa(cases)
	testsStr := strconv.Itoa(tests)
	deathsStr := strconv.Itoa(deaths)
	fmt.Printf(strings.Join([]string{casesStr, testsStr, deathsStr}, ",") + "\n")
}

// forkJoinSolve solves an instance of the covid data wrangling problem provided by the
// shared context. this will divide the total work (all files) into approximately
// equal portions among the threads to be spawned. each thread will process its set of
// files and update the tallied tests, cases, and deaths for the zipcode, month, year
func forkJoinSolve(context *sharedContext, threads int) {
	numFilesPerThread := len(*context.filenames) / threads                       // generic workload (int division)
	numFilesBigThread := len(*context.filenames) - numFilesPerThread*(threads-1) // workload for thread doing extra work
	var strt, end int
	if numFilesPerThread > 0 {
		for i := 0; i < threads-1; i++ { // create (threads-1) other goroutines
			// determine the file indices to include in this thread
			strt = i * numFilesPerThread
			end = strt + (numFilesPerThread - 1)
			context.group.Add(1)
			go worker(context, strt, end)
		}
	}
	// assign to main thread the remainder load
	context.group.Add(1)
	worker(context, end, end+numFilesBigThread-1)
	context.group.Wait()
}

// mapSolve solves an instance of the covid data wrangling problem provided by the
// shared context. this will spawn a goroutine for every file to be processed in the
// problem instance provided by context
func mapSolve(context *sharedContext) {
	threads := len(*context.filenames)
	for i := 0; i < threads-1; i++ {
		context.group.Add(1)
		go worker(context, i, i)
	}
	// assign to main thread the remainder load
	context.group.Add(1)
	worker(context, threads-1, threads-1) // threads-1 == last index
	context.group.Wait()
}

// sharedContext is struct for a problem instance of the covid data
// wrangling problem
type sharedContext struct {
	cases     *int
	tests     *int
	deaths    *int
	zipcode   *string
	month     *int
	year      *int
	filenames *[]string
	recordMap *map[string][3]int
	group     *sync.WaitGroup
	flag      *int64
}

// worker goroutine is the goroutine to be spawn when solving the covid data wrangling
// problem with a parallelized mechanism; two indices i,j are provided, indicating the
// contiguous range of files to be processed by this goroutine (e.g. 3,6 means process
// file 3,4,5,6. This goroutine calls another goroutine to facilitate
func worker(ctx *sharedContext, i, j int) {

	var err error
	var uniqueRows map[string][3]int = map[string][3]int{}

	zipcode := ctx.zipcode
	year := ctx.year
	month := ctx.month

	// this thread tasked to process files i,i+1,...,j-1,j
	filesToProcess := *ctx.filenames
	for _, filename := range filesToProcess[i : j+1] {
		err = readAndParseRows(&uniqueRows, &filename, zipcode, year, month)
		if err != nil {
			fmt.Printf("encountered error reading '%s'\n", filename)
			fmt.Println(err)
		}
	}

	// (CRITICAL SECTION END) THREADS TAKE TURNS MERGING THEIR UNIQUE RESULTS TO TALLIES / UPDATING RECORDMAP
	for !atomic.CompareAndSwapInt64(ctx.flag, 0, 1) {
	}

	recordMap := *ctx.recordMap
	// short-circuit ability: if sharedcontext has empty map just hand over this thread's map
	if len(recordMap) == 0 {
		ctx.recordMap = &uniqueRows
	}
	// else copy in results
	for weekStart, tallies := range uniqueRows {
		if _, ok := recordMap[weekStart]; !ok { // this row is globally new! add its contribute to global tallies
			recordMap[weekStart] = [3]int{}
			*ctx.cases += tallies[0]
			*ctx.tests += tallies[1]
			*ctx.deaths += tallies[2]
		}
	}

	atomic.StoreInt64(ctx.flag, 0)
	// (CRITICAL SECTION END) THREADS TAKE TURNS MERGING THEIR UNIQUE RESULTS TO TALLIES / UPDATING RECORDMAP

	ctx.group.Done()

}

// readAndParseRows will read a csv file and will aggregate into a map the tests,cases, and deaths,
// for the zipcode, year, month specified. returned map holds the week_start as keys, and a length-3
// int array as values holding respectively tests,cases, and deaths
// (details of the problem instance are inside sharedContext). It will read the file, parse and
// interpret each row, and will add any new entries to a growing map of unique rows (uniqueRows)
func readAndParseRows(uniqueRows *map[string][3]int, filename *string, zipcode *string, year *int, month *int) error {
	// var cases, tests, deaths int
	var deltaCases, deltaTests, deltaDeaths int
	var correctZip, correctYear, correctMonth, missingData bool
	var weekStartStr string

	alreadySeenRows := *uniqueRows

	// open file
	readFile, err := os.Open(*filename)
	if err != nil {
		return errors.New("had trouble reading file")
	}
	defer readFile.Close()

	// read and parse row by row
	// code idea borrowed from: https://linuxhint.com/golang-read-csv/
	csvReader := csv.NewReader(readFile)
	for {
		row, err := csvReader.Read()
		if err != nil || err == io.EOF {
			break
		}

		// check zipcode; skip if no match
		correctZip = row[zipcodeCol] == *zipcode
		if !correctZip {
			continue
		}

		// check month and year; skip if no match
		weekStartStr = row[weekStart] // format "MM/DD/YYYY"
		t, _ := time.Parse(MMDDYYYY, weekStartStr)
		correctYear = t.Year() == *year
		correctMonth = int(t.Month()) == *month
		if !(correctMonth && correctYear) {
			continue
		}

		// check if this row has missing data; skip if so
		missingData = strings.Trim(row[casesWeek], " ") == ""
		missingData = missingData || strings.Trim(row[testsWeek], " ") == ""
		missingData = missingData || strings.Trim(row[deathsWeek], " ") == ""
		if missingData {
			continue
		}

		if _, ok := alreadySeenRows[weekStartStr]; !ok { // if this row is new, add it.
			deltaCases, _ = strconv.Atoi(row[casesWeek])
			deltaTests, _ = strconv.Atoi(row[testsWeek])
			deltaDeaths, _ = strconv.Atoi(row[deathsWeek])
			alreadySeenRows[weekStartStr] = [3]int{deltaCases, deltaTests, deltaDeaths}
		}
	}

	return nil
}

// genUsageStr returns the usage string for this commandline program
func genUsageStr() string {
	return "Usage: covid [-mode=value] [-threads=int] zipcode month year\n" +
		"    zipcode = a possible Chicago zipcode\n" +
		"    month = the month to display for that zipcode \n" +
		"    year  = the year to display for that zipcode \n" +
		"    -threads = optional the number of threads (i.e., goroutines to spawn) set to number of logical cores by default\n" +
		"    -mode value = optional values: static= static dist (fork-join), map = map dist \n"
}
