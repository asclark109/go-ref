package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	// var things = [4]string{} // array
// 	var things = []string{"maleta", "ropa", "reloj"} // slice
// 	fmt.Println(things)                              // [maleta ropa reloj]

// 	things = append(things, "bolso") // appending item(s) to existing slice
// 	fmt.Println(things)              // [maleta ropa reloj bolso]

// 	things = append(things[1:]) // to shorten a slice and release memory properly
// 	// things = append(things[1 : len(things)-1]) // same as above (defaulting)
// 	fmt.Println(things) // [ropa reloj bolso]

// 	heros := make([]string, 3, 3) // initialize slice object (create underlying array), length 3, capacit 3
// 	// even though capacity is 3, slice capacity will grow dynamically as needed
// 	heros[0] = "thor"
// 	heros[1] = "ironman"
// 	heros[2] = "spiderman"
// 	fmt.Println(heros) // [thor ironman spiderman]

// 	heros = append(heros, "deadpool") // slice capacity will grow dynamically by 2x
// 	fmt.Println(heros)                // [thor ironman spiderman deadpool]
// 	fmt.Println(len(heros))           // 4
// 	fmt.Println(cap(heros))           // 6

// 	myints := []int{4, 7, 3, 2, 9}
// 	isSorted := sort.IntsAreSorted(myints)
// 	fmt.Println("Are ints sorted: ", isSorted) // Are ints sorted:  false

// 	fmt.Println("BEFORE: ", myints) // BEFORE:  [4 7 3 2 9]
// 	sort.Ints(myints)
// 	fmt.Println("AFTER: ", myints) // AFTER:  [2 3 4 7 9]

// }
