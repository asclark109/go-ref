package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	var name string
// 	var userRating string

// 	//Front end
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("please enter your full name: ")
// 	name, _ = reader.ReadString('\n') // don't need : because type already assigned

// 	//take rating from user and convert it to float
// 	reader = bufio.NewReader(os.Stdin)
// 	fmt.Println("Please rate our Dosa center between 1 and 5: ")
// 	userRating, _ = reader.ReadString('\n')
// 	userRatingNum, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64) // need : to assign type

// 	//Back end
// 	currTime := time.Now().Format(time.Stamp) // or time.Kitchen
// 	fmt.Printf("Hello %v, \n Thanks for rating our dosa center with %v star rating. \n\n Your rating was recorded in our system at %v\n\n", name, userRatingNum, currTime)

// 	if userRatingNum == 5 {
// 		fmt.Println("Bonus for team for 5 star service")
// 	} else if userRatingNum == 4 || userRatingNum == 3 {
// 		fmt.Println("we are always improving")
// 	} else if userRatingNum < 3 {
// 		fmt.Println("we need serious work on our side")
// 	}

// }
