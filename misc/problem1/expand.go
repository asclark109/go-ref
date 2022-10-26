package problem1

import (
	"errors"  // raising errors
	"regexp"  // regex to detect ranges
	"strconv" // basic string conversions (e.g. string to int)
	"strings" // basic string operations
)

func Expand(intList string) []int {
	// HIGH LEVEL ALGORITHM:
	// (1) split string by "," to obtain slice of strings.
	// (2) for each string in slice:
	//   (2a) remove left and right trailing white space.
	//   (2b) if string is a valid num, add it to a growing slice of ints
	//   (2c) if string is a valid range ("###-###"), generate representative []int
	//        and add []int to growing slice of ints
	// (3) generate new []int without duplicates

	// (1)
	var strElems []string = strings.Split(intList, ",")

	// (2)
	var rawInts []int
	for i := range strElems {

		// (2a)
		strTrimmed := strings.TrimSpace(strElems[i])

		// (2b, 2c)
		if isPostiveInt(strTrimmed) {
			intNum, _ := strconv.Atoi(strTrimmed)
			rawInts = append(rawInts, intNum)
		} else if min, max, err := getMinMax(strTrimmed); err == nil {
			newIntSlice := rangeToIntSlice(min, max)
			rawInts = append(rawInts, newIntSlice...)
		}
	}
	// (3)
	return toUniqueSlice(rawInts)
}

// isPositiveInt returns true if string can be interpretted as a positive int.
// e.g. "-5" -> false; "30" -> true; "3.04" -> false; " 3" -> false;
func isPostiveInt(str_word string) bool {
	if intNum, err := strconv.Atoi(str_word); err == nil {
		if intNum > 0 {
			return true
		}
	}
	return false
}

// getMinMax attempts to interpret a string as a range.
// If successful, returns the min and max of the range;
// otherwise, returns -1, -1, and an error.
// e.g. "4-5" -> 4,5,nil; "3 -30" -> -1,-1,error;
func getMinMax(strRange string) (int, int, error) {
	validRangeRegex := regexp.MustCompile("^[0-9]+-[0-9]+$")
	match := validRangeRegex.Match([]byte(strRange))
	if match {
		minAndMax := strings.Split(strRange, "-")
		minStr := minAndMax[0]
		maxStr := minAndMax[1]
		if isPostiveInt(minStr) && isPostiveInt(maxStr) {
			min, _ := strconv.Atoi(minStr)
			max, _ := strconv.Atoi(maxStr)
			if min < max {
				return min, max, nil
			}
		}
	}
	return -1, -1, errors.New("string can't be interpretted as a range")
}

// rangeToIntSlice takes two integers min and max (min < max assumed).
// returns a slice of all integers between min and max, inclusive.
// e.g. 5,9 -> [5 6 7 8 9]
func rangeToIntSlice(min, max int) []int {
	intSlice := make([]int, max-min+1)
	counter := min
	for i := 0; i < max-min+1; i++ {
		intSlice[i] = counter
		counter += 1
	}
	return intSlice
}

// toUniqueSlice accepts a slice of ints and returns a new slice with all duplicates removed.
func toUniqueSlice(nums []int) []int {
	uniqueSlice := make([]int, 0, len(nums))
	m := make(map[int]bool) // will indicate whether int already added to growing slice
	for _, num := range nums {
		if _, alreadyAdded := m[num]; !alreadyAdded {
			m[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}
	return uniqueSlice
}
