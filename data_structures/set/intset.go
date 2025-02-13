package set

import (
	"sort"
	"strconv"
	"strings"
)

// IntSet represents a set of integers
type IntSet struct {
	intMap map[int]interface{} // Use a map with only keys to hold the collection of integers that are part of the set.
	// NOTE: change to map approved by Lamont Samuels
}

// NewIntSet returns a newly allocated IntSet with its contents initilaized to zero
func NewIntSet() *IntSet {
	return &IntSet{make(map[int]interface{}, 0)}
}

// Add adds a number to this IntSet
func (recv *IntSet) Add(num int) {
	// if num is not already a key in underlying map, add it as new key
	if _, ok := recv.intMap[num]; !ok {
		recv.intMap[num] = nil
	}
}

// Prettify generates a prettily formated string representation of this IntSet
func (recv *IntSet) Prettify() string {
	numsAsStrs := make([]string, len(recv.intMap)) // will hold string repr of nums in set
	numsAsSlice := toIntSlice(&recv.intMap)        // puts keys from map[int] into an []int so we can sort it
	sort.Ints(*numsAsSlice)
	// generate slice of strs representing sorted list of nums
	for i, intNum := range *numsAsSlice {
		numsAsStrs[i] = strconv.Itoa(intNum)
	}
	return "{" + strings.Join(numsAsStrs, ",") + "}"
}

// Union generates a new IntSet from unioning this IntSet with another IntSet
func (recv *IntSet) Union(other *IntSet) *IntSet {
	// O(N+M) expected time algorithm; (N, M == number of elems in each IntSet)
	// (1) iterate thru keys in first map, add each key to a new map.
	// (2) iterate thru keys in second map, if key not already in new map, add it.
	// (3) create new IntSet from new map.

	intMapUnion := make(map[int]interface{}) // create new map to store result of operation
	// (1)
	for num := range recv.intMap {
		intMapUnion[num] = nil
	}
	// (2)
	for num := range other.intMap {
		if _, ok := intMapUnion[num]; !ok {
			intMapUnion[num] = nil
		}
	}
	// (3)
	return &IntSet{intMapUnion}
}

// Intersect generates a new IntSet from intersecting this IntSet with another IntSet
func (recv *IntSet) Intersect(other *IntSet) *IntSet {
	// O(N+M) expected time algorithm (N, M == number of elems in each IntSet)
	// (1) iterate through nums in 1st IntSet map
	//        if num also in 2nd IntSet map, add num to a new map
	// (2) return new IntSet holding the map containing the intersection

	intMapIntersection := make(map[int]interface{}) // store result of operation in new map
	for num := range recv.intMap {
		if _, ok := other.intMap[num]; ok {
			intMapIntersection[num] = nil
		}
	}
	return &IntSet{intMapIntersection}
}

// Diff generates a new IntSet from the set subtraction operation: this IntSet - other IntSet
func (recv *IntSet) Diff(other *IntSet) *IntSet {
	// O(N+M) expected time algorithm (N, M == number of elems in each IntSet)
	// (1) iterate through nums in 1st IntSet map
	//        if num NOT in 2nd IntSet map, add num to a new map
	// (2) return new IntSet holding the map containing the set difference

	intMapDiff := make(map[int]interface{}) // store result of operation in new map
	for num := range recv.intMap {          // in set 1
		if _, ok := other.intMap[num]; !ok { // not in set 2
			intMapDiff[num] = nil
		}
	}
	return &IntSet{intMapDiff}
}

// toIntSlice generates a new int slice from an existing map with integer keys
func toIntSlice(intMap *map[int]interface{}) *[]int {
	ints := make([]int, len(*intMap))
	idx := 0
	for num := range *intMap {
		ints[idx] = num
		idx++
	}
	return &ints
}
