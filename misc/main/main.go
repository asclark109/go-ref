package main

import (
	"fmt"
	"proj1/feed"
)

func main() {
	fmt.Println("hello")
	f := feed.NewFeed()
	fmt.Println(f.Display())
	f.Add("hello", float64(2))
	fmt.Println(f.Display())
	f.Add("hello", float64(3))
	fmt.Println(f.Display())
	f.Add("hello", float64(10))
	fmt.Println(f.Display())
	f.Add("hello", float64(3))
	fmt.Println(f.Display())

	f.Remove(float64(10))
	fmt.Println(f.Display())
	f.Remove(float64(3))
	fmt.Println(f.Display())
	f.Remove(float64(3))
	fmt.Println(f.Display())
	f.Remove(float64(3))
	fmt.Println(f.Display())
	f.Remove(float64(3))
	fmt.Println(f.Display())
	f.Remove(float64(2))
	fmt.Println(f.Display())

	// fmt.Println(f.Contains(float64(1)))
	// fmt.Println(f.Contains(float64(2)))
	//Check to make sure Contains returns False on empty feed
	// for i := 1; i <= rand.Intn(100); i++ {
	// 	if feed.Contains(float64(i)) {
	// 		t.Errorf("Feed is empty. Contains = %v", i)
	// 	}
	// }
}
