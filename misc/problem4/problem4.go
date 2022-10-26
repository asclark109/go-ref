package main

import (
	"bufio"
	"errors"
	"fmt"
	"hw1/set"
	"os"
	"strconv"
	"strings"
)

func main() {
	// HIGH LEVEL BEHAVIOR
	// (1) check for invalid arguments provided to program (terminate early on any issue)
	// (2) create IntSets from data in file, perform IntSet operation, display result

	// (1)
	// CHECK: number of args
	arguments := os.Args[1:]
	if !(len(arguments) == 2) {
		fmt.Println(genErrorMsg())
		return
	}

	// CHECK: understood action
	action := arguments[0]
	validActions := []string{"union", "intersect", "diff"}
	if !contains(validActions, action) {
		fmt.Println(genErrorMsg())
		return
	}

	// CHECK: can open file
	// borrowed logic from https://golangdocs.com/golang-read-file-line-by-line
	filename := arguments[1]
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(genErrorMsg())
		return
	}

	// CHECK: can read file; file has expected number of lines (2 lines of numbers)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	if len(lines) != 2 {
		fmt.Println(genErrorMsg())
		readFile.Close() // close reader before exiting program
		return
	}
	readFile.Close()

	// (2)
	// interpret lines and create IntSets from each line
	sets := [2]*set.IntSet{} // above check confirms two lines in file
	for i, line := range lines {
		intNums, _ := parseLine(line) // assume files well-structured
		intSet := set.NewIntSet()
		for _, num := range intNums {
			intSet.Add(num)
		}
		sets[i] = intSet
	}

	// perform set operation and print to console
	var resultantSet *set.IntSet
	switch {
	case action == "union":
		resultantSet = sets[0].Union(sets[1])
	case action == "intersect":
		resultantSet = sets[0].Intersect(sets[1])
	case action == "diff":
		resultantSet = sets[0].Diff(sets[1])
	}
	fmt.Println(resultantSet.Prettify())
}

// contains returns true if a query string is an element of a string slice
func contains(strSlice []string, strQuery string) bool {
	for _, strWord := range strSlice {
		if strQuery == strWord {
			return true
		}
	}
	return false
}

// parseLine accepts a string representing a line from file with format "**##**##**...",
// i.e. a sequence of integers with any amount of whitespace left or right trailing each integer,
// and returns those integers as an int slice. on error, returns []int{}, error.
func parseLine(line string) ([]int, error) {
	intNums := []int{}
	strElems := strings.Fields(line)
	for _, strElem := range strElems {
		if intNum, err := strconv.Atoi(strElem); err == nil {
			intNums = append(intNums, intNum)
		} else {
			return []int{}, errors.New(fmt.Sprintf("encountered error parsing line to slice of ints: %v", strElems))
		}
	}
	return intNums, nil
}

// genErrorMsg returns a string that should be printed in the case of program error.
func genErrorMsg() string {
	return "Usage: problem4  <union | intersect | diff> file"
}
