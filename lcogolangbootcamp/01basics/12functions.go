package main

import (
	"fmt"
)

func main() {
	superman() // invoke function

	result := multiplyme(1, 3)
	fmt.Println(result)

	result2, result3 := multiplyanddivide(1, 3)
	fmt.Printf("%v %T\n", result2, result2)
	fmt.Printf("%v %T\n", result3, result3)

	res := adder(1, 2, 3, 4)
	fmt.Println(res)

	myres, mylen, myname := adder2(1, 2, 3, 4)
	fmt.Println(myres, mylen, myname)

}

func superman() {
	fmt.Println("I am superman")
}

func multiplyme(v1 int, v2 int) int {
	return v1 * v2
}

func multiplyme2(v1, v2 int) int { // syntactic sugar
	return v1 * v2
}

func multiplyanddivide(v1, v2 int) (int, int) { // syntactic sugar
	return v1 * v2, v1 / v2
}

// no function overloading in go; i.e., must specify the number of args
func adder(values ...int) int {
	sum := 0
	for _, num := range values {
		sum += num
	}
	return sum
}

// no function overloading in go; i.e., must specify the number of args
func adder2(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}

func multiplier(values ...int) (int, string) {
	product := 0
	for _, num := range values {
		product *= num
	}
	return product, "lco"
}
