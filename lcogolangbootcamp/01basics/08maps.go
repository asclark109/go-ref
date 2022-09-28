package main

// import "fmt"

// func main() {
// 	// usually use make when creating map
// 	score := make(map[string]int) // [key]values. using automatically built-in new keyword; allocating memory but not initializing memory
// 	score["mine"] = 89
// 	fmt.Println(score) // map[mine:89]

// 	score["josh"] = 34
// 	score["ron"] = 23
// 	score["sam"] = 56
// 	score["vicky"] = 78
// 	fmt.Println(score) // map[josh:34 mine:89 ron:23 sam:56 vicky:78]

// 	getRonScore := score["ron"]
// 	fmt.Println(getRonScore) // 23

// 	delete(score, "vicky")
// 	fmt.Println(score) // map[josh:34 mine:89 ron:23 sam:56]

// 	// iteration through keys and values
// 	for k, v := range score { // k, v could be any var reference name e.g. a, b
// 		fmt.Printf("score of %v is %v\n", k, v)
// 	}
// 	// score of sam is 56
// 	// score of mine is 89
// 	// score of josh is 34
// 	// score of ron is 23
// }
