package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // to take input from user; can create a reader that can read from standard input
// // string converter will need for project

// func main() {
// 	// var myString string
// 	// fmt.Scanln(&myString) // & means store result in variable myString
// 	// // note, weirdness happens if you enter a space eg Austin Clark
// 	// // 1 2 3 4
// 	// // 1
// 	// fmt.Println(myString)

// 	// var name string = "hitesh"
// 	// var a, _ = fmt.Println(name)
// 	// fmt.Println(a)

// 	// reader := bufio.NewReader(os.Stdin) // infer a value
// 	// fmt.Print("Enter your full name: ")
// 	// myname, _ := reader.ReadString('\n')
// 	// fmt.Println(myname)

// 	// reader := bufio.NewReader(os.Stdin)
// 	// fmt.Print("enter your rating")
// 	// myrating, _ := reader.ReadString('\n')
// 	// fmt.Println(myrating)

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("enter your rating: ")
// 	myrating, _ := reader.ReadString('\n')
// 	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
// 	fmt.Println(mynumrating + 2)
// }
