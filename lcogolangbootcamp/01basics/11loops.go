package main

import "fmt"

func main() {
	start := 1
	things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(start)  // 1
	fmt.Println(things) // [maleta bolso boleto vaso casa]

	for i := 0; i < 10; i++ { // C style; i is only in the scope of loop
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i := range things { // for i in range things
		fmt.Println(things[i])
	}

	for _, thing := range things { // for item in items
		fmt.Println(thing)
	}

	for start < 100 { // kind of like a while loop
		start += start
		if start == 32 {
			break // breaks to just outside this for loop
		}
		fmt.Println("START is now:", start)
	}

	for start < 100 {
		start += start
		if start == 32 {
			continue // skip to next iteration of loop (jump to top of loop to conidition statement)
		}
		fmt.Println("START is now:", start)
	}

	for start < 100 { // kind of like a while loop
		start += start
		if start == 32 {
			goto lcolabel // jumps to the label 'lcolabel'
		}
		fmt.Println("START is now:", start)
	}

lcolabel:
	fmt.Println("print first line")
	fmt.Println("print another line")

}
