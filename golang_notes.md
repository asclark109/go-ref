-

# Go (Golang)
https://www.youtube.com/watch?v=X4q1OM0voO0&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=1

Called "Golang" to avoid confusion caused by searching "Go" on search engines.

Features:
- Low Latency (if you are concerned about latency; adopted by twitch)
- Garbage Collection (if you want very fast and lots of support)
- GPU (if concerned about GPU power and utilization)
- Concurrency Support (built-in support)

### Why new language?

Python easy, C++ fast, Java garbage collection. 21st century C programming language. Takes syntax from languages like C, Java and makes it rediculously fast. C,C++ fast but compile time slow. Concurrency was added to these languages later. Golang takes advantage of high memory RAM and GPU power in today's computers.

Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines [golang docs].

### Installation

- `MOD`: for initializing packages...
- `ENV`: environment variables; many env vars set by go. `GoROOT` and `GoWorkSpace` are important. `GoROOT` is where all binaries to run the program are stored. `GoWorkSpace` is the folder where you put your code. Go tells you specifically organize your code in a strict directory structure. Though, this structure is only expected if the code is to go in production or go to a public repository.
- `RUN`:
- `GET`: to get things from the repository
- `FMT`:

Special Directory Structure.
```
project/
  bin/
  pkg/
  src/
    github.com/
      asclark/
        helloWorld  
```
`bin/` holds binaries. `pkg/` holds packages you get from online although that's not compulsory. `src/` where your code goes.

in terminal, see version
```
$ go version
go version go1.19.1 windows/amd64
```

see env variables
```
$ go env
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\ascla\AppData\Local\go-build
set GOENV=C:\Users\ascla\AppData\Roaming\go\env
set GOEXE=.exe
set GOEXPERIMENT=
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\Users\ascla\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\ascla\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=C:\Program Files\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=C:\Program Files\Go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.19.1
set GCCGO=gccgo
set GOAMD64=v1
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=NUL
set GOWORK=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=C:\Users\ascla\AppData\Local\Temp\go-build539226397=/tmp/go-build -gno-record-gcc-switches
```

installation of go located at `set GOROOT=C:\Program Files\Go`. Will help you when running files. GoWorkSpace is indicated by `set GOPATH=C:\Users\ascla\go`

```go
var name string = "hitesh"
var a, _ = fmt.Println(name) // use underscore if you dont want var
fmt.Println(a)
```

### input from user

reading string input from user
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)
// read input from user
func main() {
  // read a string
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter your full name: ")
  myname, _ := reader.ReadString('\n')
  fmt.Println(myname)
}

// if reading number from input, need another package
// because entering a number then hitting enter will
// append characters \n to the end of the number in the input; hence can't be parsed as float
```

Reading numeric input from user
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// to take input from user that is a numeric
// need to use some more built in 
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumrating + 2)
}
```

### example program

obtain a rating from a user and then announce that rating was logged.
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Front end
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your full name: ")
	name, _ = reader.ReadString('\n') // don't need : because type already assigned

	//take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Dosa center between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	userRatingNum, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64) // need : to assign type

	//Back end
	currTime := time.Now().Format(time.Stamp) // or time.Kitchen
	fmt.Printf("Hello %v, \n Thanks for rating our dosa center with %v star rating. \n\n Your rating was recorded in our system at %v\n\n", name, userRatingNum, currTime)

	if userRatingNum == 5 {
		fmt.Println("Bonus for team for 5 star service")
	} else if userRatingNum == 4 || userRatingNum == 3 {
		fmt.Println("we are always improving")
	} else if userRatingNum < 3 {
		fmt.Println("we need serious work on our side")
	}

}
```

### Pointers

```go
package main

import (
	"fmt"
)

func main() {
	var p *int // point that is not storing anything

	if p != nil {
		fmt.Println("P is having a value: ", *p) // print what pointer p points to, which is basically nil
	} else {
		fmt.Println("Watchout for nil values")
	}
}
```

```go
package main

import (
	"fmt"
)

func main() {
	var lifebooster float64 = 99.2
	lifeboosterRef := &lifebooster // reference to memory location
	//var lifeboosterRef = &lifebooster // can also be done like so
	fmt.Println(lifebooster)     // prints 99.2
	fmt.Println(lifeboosterRef)  // prints 0xc0000a6058
	fmt.Println(*lifeboosterRef) // prints 99.2, the thing located at 0xc0000a6058

  lifebooster = lifebooster * 2.2

  fmt.Println(lifebooster)     // prints 218.24000000000004
	fmt.Println(lifeboosterRef)  // prints 0xc0000a6058
	fmt.Println(*lifeboosterRef) // prints 218.24000000000004, the thing located at 0xc0000a6058
}
```
This prints `99.2` and `0xc000018098` (reference to location in memory)

### Arrays

Array type definition specifies and a length and element type (e.g. `[4]int`). The array's size is fixed and part of its type.

```go
var a [4]int
a[0] = 1
i := a[0]
// i == 1
```
Arrays do not need to be initialized explicitly; the zero value of an array is a ready-to-use array whose elements are themselves zeroed:
```go
// a[2] == 0, the zero value of the int type
```

arrays in Go are values. An array variable denotes the entire array (not a pointer to the first array element as in C). This means that when you assign or pass around an array value you will make a copy of its contents (to avoid the copy you could pass a `pointer` to the array, but then that's a pointer to an array, not an array).

The array is sort of like a struct but with indexed rather than named fields; a fixed-size composite value.

specifying array literal
```go
b := [2]string{"Penn", "Teller"}
c := [...]string{"Penn", "Teller"} // compiler counts the size
```
the type of `b` and `c` is `[2]string`

```go
package main

import "fmt"

func main() {
	var numbers [3]string // declare an array of strings of size 3
	numbers[0] = "uno"
	numbers[1] = "dos"
	numbers[2] = "tres"

	fmt.Println(numbers) /// [uno dos tres]

	var colors = [4]string{"rojo", "gris", "azul", "verde"} // comma separated values
	fmt.Println(colors)                                     // [rojo gris azul verde]
	fmt.Println(colors[2])                                  // azul
	fmt.Println(len(colors))                                // 4
}
```

### Slices

`Slices` give the benefits of the array but scale dynamically. They are an abstraction layer built on top of Go's array type. They are more common than arrays

Type specification for a slice is `[]T`, where `T` is the type of the elements of the slice. Slice type has no specified length.

A slice can be created with the built-in function `make`, which has signature
```go
func make([]T, len, cap) []T
```
where T stands for the element type of the slice to be created. make function takes a type, length, and optional capacity. When called, make allocates an array and returns a slice that refers to that array

```go
var s []byte // declare slice of bytes
s = make([]byte, 5, 5) // create slice object
// s == []byte{0, 0, 0, 0, 0}
```

can do this more succinctly by omitting capacity
```go
s := make([]byte, 5) // make slice of bytes of length 5
```

can inspect size and capacity with `len()` and `cap()`
```go
len(s) == 5
cap(s) == 5
```

default zero value for slice is `nil`. `len()` and `cap()` both return `0` for `nil` slice.

A slice can be formed by "slicing" an existing slice or array
```go
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
```

```go
// b[:2] == []byte{'g', 'o'}
// b[2:] == []byte{'l', 'a', 'n', 'g'}
// b[:] == b
```

can create a slice from an array
```go
x := [3]string{"Лайка", "Белка", "Стрелка"} // array of size 3
s := x[:] // a slice referencing the storage of x
```

### Slice internals

A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).

slice `length` refers to the number of elements referenced to by the slice. slice `capacity` is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).

slicing does not copy the slice's data. It creates a new slice value that points to the original array. Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice:

```go
d := []byte{'r', 'o', 'a', 'd'} // slice of bytes
e := d[2:]
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
```

can grow `s` to its capacity by slicing again:
```go
s := make([]byte, 5) // make slice of bytes of length 5
s = s[2:4] // shrink length. capacity unchanged
s = s[:cap(s)] // expand length. capacity unchanged
// s == []byte{'a', 'm'}
```

### growing a slice (`copy` and `append` functions)

increase the capacity of a slice by creating a new, larger slice and copying the contents of the original slice into it (note this can be done easier and is show)
```go
t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
for i := range s {
        t[i] = s[i]
}
s = t
```

use built-in `copy` function to grow slice. it returns the number of elements copied:
```go
func copy(dst, src []T) int
```
using `copy`, can double the size of the slice from above
```go
t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t
```

To append data to the end of the slice (e.g. a slice of bytes), use `AppendByte`, which gives great fine control over the way the slice is grown. Depending on the characteristics of the program, it may be desirable to allocate in smaller or larger chunks, or to put a ceiling on the size of a reallocation.
```go
p := []byte{2, 3, 5}
p = AppendByte(p, 7, 11, 13)
// p == []byte{2, 3, 5, 7, 11, 13}
```

Most programs do not need complete control, so `append` is used in most purposes:
```go
func append(s []T, x ...T) []T
```

`append` appends the elements `x` to the end of the slice `s`, and grows the slice if a greater capacity is needed.

To append multiple slices, use `...`:
```go
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
```

Since the zero value of a slice (`nil`) acts like a zero-length slice, you can declare a slice variable and then append to it in a loop:
```go
// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}
```

### A potential gotcha with copying slices
*potential gotcha* re-slicing a slice doesn't make a copy of the underlying array. The full array will be kept in memory until it is no longer referenced. Different from python. Be aware, the array you first declared still exists in memory.

As mentioned earlier, re-slicing a slice doesn’t make a copy of the underlying array. The full array will be kept in memory until it is no longer referenced. Occasionally this can cause the program to hold all the data in memory when only a small piece of it is needed.

For example, this `FindDigits` function loads a file into memory and searches it for the first group of consecutive numeric digits, returning them as a new slice.

```go
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}
```

This code behaves as advertised, but the returned `[]byte` points into an array containing the entire file. Since the slice references the original array, as long as the slice is kept around the garbage collector can’t release the array; the few useful bytes of the file keep the entire contents in memory.

To fix this problem one can copy the interesting data to a new slice before returning it:
```go 
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}
```

A more concise version of this function could be constructed by using `append`. This is left as an exercise for the reader.

Example usage of slices
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	// var things = [4]string{} // array
	var things = []string{"maleta", "ropa", "reloj"} // slice
	fmt.Println(things)                              // [maleta ropa reloj]

	things = append(things, "bolso") // appending item(s) to existing slice
	fmt.Println(things)              // [maleta ropa reloj bolso]

	things = append(things[1:]) // to shorten a slice and release memory properly
	// things = append(things[1 : len(things)-1]) // same as above (defaulting)
	fmt.Println(things) // [ropa reloj bolso]

	heros := make([]string, 3, 3) // initialize slice object (create underlying array), length 3, capacit 3
	// even though capacity is 3, slice capacity will grow dynamically as needed
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "spiderman"
	fmt.Println(heros) // [thor ironman spiderman]

	heros = append(heros, "deadpool") // slice capacity will grow dynamically by 2x
	fmt.Println(heros)                // [thor ironman spiderman deadpool]
	fmt.Println(len(heros))           // 4
	fmt.Println(cap(heros))           // 6

	myints := []int{4, 7, 3, 2, 9}
	isSorted := sort.IntsAreSorted(myints)
	fmt.Println("Are ints sorted: ", isSorted)

	fmt.Println("BEFORE: ", myints)
	sort.Ints(myints)
	fmt.Println("AFTER: ", myints)

}
```

### Maps

What is the difference between `make` and `new`?

- `new` -- only allocates memory - no initialization of memory
- `make` -- allocate and initialize - non zeroed storage

```go
package main

import "fmt"

func main() {
	var score map[string]int // [key]values. using automatically built-in new keyword; allocating memory but not initializing memory
	score["mine"] = 89 // panic: assignment to entry in nil map
	// memory was allocated for you but not initialized for you
}
```

creating a map, adding key value pairs, getting values at keys, deleting key value pairs. 
```go
package main

import "fmt"

func main() {
	// usually use make when creating map
	score := make(map[string]int) // [key]values. using automatically built-in new keyword; allocating memory but not initializing memory
	score["mine"] = 89
	fmt.Println(score) // map[mine:89]

	score["josh"] = 34
	score["ron"] = 23
	score["sam"] = 56
	score["vicky"] = 78
	fmt.Println(score) // map[josh:34 mine:89 ron:23 sam:56 vicky:78]

	getRonScore := score["ron"]
	fmt.Println(getRonScore) // 23

	delete(score, "vicky")
	fmt.Println(score) // map[josh:34 mine:89 ron:23 sam:56]
}

```

### structs

```go
package main

import (
	"fmt"
)

// define a struct with the word "type"
type User struct { // name of struct is upper case
	Name  string // attributes are upper case too
	Email string
	age   int
}

func main() {
	rob := User{"rob", "rob@lco.dev", 34} // creating a struct
	fmt.Printf("%v\n", rob) // {rob rob@lco.dev 34}
	// use %+v to obtain more details on the object
	fmt.Printf("%+v\n", rob) // {Name:rob Email:rob@lco.dev age:34}

	var sam = new(User) // remember only allocates memory. doesn't initialize
	sam.Name = "sam"
	sam.Email = "sam@lco.dev"
	sam.age = 22
	fmt.Printf("%v\n", sam)  // &{sam sam@lco.dev 22}
	fmt.Printf("%+v\n", sam) // &{Name:sam Email:sam@lco.dev age:22}

	// remember the 'new' only allocates memory. it doesn't initialize.
	// The & in &{sam sam@lco.dev 22} indicates a reference you are getting back.
	// The & symbol points to the address of the stored value

	var tobby = &User{"tobby", "tobby@lco.dev", 22} // less common
	fmt.Printf("%+v\n", tobby)                      // &{Name:tobby Email:tobby@lco.dev age:22}

}

```

### loops

loops follow C style structure but without parentheses
```go
package main

import "fmt"

func main() {
	start := 1
	things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(start) // 1
	fmt.Println(things) // [maleta bolso boleto vaso casa]

	for i := 0; i < 10; i++ {
		fmt.Println(i + start)
	}
}
```

Use `break` and `continue` in loops to break out of loops or skip to the next iteration of the loop. There is also a `goto` keyword.

```go
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
			// break    // breaks to just outside this for loop
			continue // skip to next iteration of loop (jump to top of loop to conidition statement)
		}
		fmt.Println("START is now:", start)
	}
}
```

### functions

Defining functions with names that start with lower-case letters are different from those that start with upper-case letters!

```go
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
```
