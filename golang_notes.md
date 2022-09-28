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

### Basics (Hello World, running binaries)

Hello World.
```go
package main

import "fmt"

func main(){
  fmt.Println("hello world") // print to stdout
}
```

To run a go program.
```
$ go run hello-world.go
```

Can build the code into binaries.
```
$ go build hello-world.go
$ ls
hello-world    hello-world.go
```

Can execute the binaries directly.
```
$ ./hello-world
hello world
```

Some Simplicities.
```go
var name string = "hitesh"    // basic declaration and initialization of variable
var a, _ = fmt.Println(name)  // use underscore to assign unused values
fmt.Println(a)
```

### Values

Go has various value types: strings, integers, floats, booleans with as expected behavior
for the operators `+`, `-`, `&&`, `||`, `!`.
```go
package main

import "fmt"

func main() {

  "go" + "lang"  // "golang"
  1+1            // 2
  7.0/3.0        // 2.3333333333333335
  true && false  // false
  true || false  // true
  !true          // false
}
```

### Variables

In Go, variables are explicitly declared and used by compiler to check type-correctness of function calls.

`var` declares 1 or more variables. \
Go supports multiple assignment. `var b, c int = 1, 2`\
Go will infer the type of initialized variables. `var d = true` \
Can declare variables but not initialize them. `var e int`\
`:=` shorthand will declare and initialize a variable (and do type inference). `f := "apple"` instead of `var f string = "apple"`

```go
package main

import "fmt"

func main() {

  var a = "initial" // can infer the type of initialized variables
  var b, c int = 1, 2 // supports multiple assignment
  var d = true 

  var e int // declares a variable (no initialization)
  fmt.Println(e) // 0; declared variables that aren't intialized are zero-valued

  f := "apple" // := syntax is shorthand for declaring and initializing a variable, 
  fmt.Println(f)
}
```

### Constants

Go supports *constants* of character, string, boolean, and numeric values.

`const` declares a constant value.
A `const` statement can appear anywhere a `var` statement can.
Constant expressions perform arithmetic with arbitrary precision.
A numeric constant has no type until it’s given one, such as by an explicit conversion.
A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here `math.Sin` expects a `float64`.

```go
package main

import (
  "fmt"
  "math"
)

const s string = "constant"

func main() {
  fmt.Println(s)           // constant
  const n = 500000000
  const d = 3e20 / n       // 6e+11
  fmt.Println(int64(d))    // 600000000000
  fmt.Println(math.Sin(n)) // -0.28470407323754404
}
```

### Getting input from user

Reading string input from user
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)    // declare and initialize (with type inference) bufio.Reader (reading from Stdin)
  fmt.Print("Enter your full name: ")
  myname, _ := reader.ReadString('\n')   // store result from input at the terminal into myname. 2nd assignment is ignored (returns an error as 2nd assignment)
  fmt.Println(myname)
}

// if reading number from input, need another package
// because entering a number then hitting enter will
// append characters \n to the end of the number in the input; hence can't be parsed as float
```

Reading numeric input from user requires more work.
Entering a number and then hitting enter will append \n to the end of the input; e.g. entering '34' at console results in '34\n' and hence can't be parsed as float.
```go
package main

import (
	"bufio"    // for reading input
	"fmt"      // for printing
	"os"       // for reading input from Stdin
	"strconv"  // to convert str to float
	"strings"  // to trim strings
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumrating + 2)
}
```


### Iteration statement (`for` loops)

`for` is the only looping construct in Go. 

`break`, `continue`, and `goto` are supported for for loops.

Use `break` to break directly outside of a loop.\
Use `continue` to skip to the next iteration of the loop.

```go
package main

import "fmt"

func main() {
  // loops follow C style structure but without parentheses

  // single condition for loop
  i := 1
  for i <= 3 { 
    fmt.Println(i)
    i = i + 1
  }

  // classic for loop; initial / condition / after
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  // for loop without a condition will loop repeatedly until you break out of the loop
  // or return from the enclosing function
  for {
    fmt.Println("loop")
    break            // breaks to just outside this for loop
  }

  // can also continue to the next iteration of the loop
  for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue     // skip to next iteration of loop (jump to top of loop to conidition statement)
    }
    fmt.Println(n)
  }

  //// MORE EXAMPLES ////
  start := 1
	things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(start)                   // 1
	fmt.Println(things)                  // [maleta bolso boleto vaso casa]

	for i := 0; i < 10; i++ {            // i is only in the scope of loop
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i := range things {              // for i in range things (i is the idx)
		fmt.Println(things[i])
	}

	for _, thing := range things {       // for item in items
		fmt.Println(thing)
	}

	for start < 100 {                    // kind of like a while loop
		start += start
		if start == 32 {
			// break    // breaks to just outside this for loop
			continue // skip to next iteration of loop (jump to top of loop to conidition statement)
		}
		fmt.Println("START is now:", start)
	}
}
```

### Branching Statement (If/Else)

Can use `if` and `else` for branching. 

```go
package main

import "fmt"

func main() {

  // basic
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  // if statement
  if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  // A statement can precede conditionals;
  // any variables declared in this statement are available in all branches (num);
  // parentheses around conditions not needed. braces {} are needed.
  if num := 9; num < 0 { // initial statement; condition
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}
```



### Branching Statement (Switch)

Switch statements express conditionals across many branches.

```go 
package main

import (
    "fmt"
    "time"
)

func main() {

  // basic switch statement
  i := 2
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  // can use commas to separate multiple expressions in the same case statement.
  // this example uses the optional default case.
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("It's a weekday")
  }

  // switch without an expression is an alternate way to express if/else logic.
  // Here we also show how the case expressions can be non-constants.
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  // A type switch compares types instead of values. You can use this to discover
  // the type of an interface value. In this example, the variable t will have 
  // the type corresponding to its clause.
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int")
    default:
        fmt.Printf("Don't know type %T\n", t)
    }
  }
  whatAmI(true) // "I'm a bool"
  whatAmI(1) // "I'm an int"
  whatAmI("hey") // "Don't know type string"
  }
```

### Data Structures (Arrays)

In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.

```go
package main

import "fmt"

func main() {

  var a [5]int                 // declares an int array of size 5; note BOTH type of elements and length become part of the array's type
  fmt.Println("emp:", a)       // emp: [0 0 0 0 0]

  a[4] = 100 
  fmt.Println("set:", a)       // set: [0 0 0 0 100]
  fmt.Println("get:", a[4])    // get: 100
  fmt.Println("len:", len(a))  // len: 5

  b := [5]int{1, 2, 3, 4, 5}   // declare and initialize an array in one line. int array of size 5.
  fmt.Println("dcl:", b)       // dcl: [1 2 3 4 5]

  // Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)    // 2d:  [[0 1 2] [1 2 3]]
}
```

Array type definition specifies and a length and element type (e.g. `[4]int`). The array's size is fixed and part of its type.

```go
var a [4]int   // declare an int array of size 4
a[0] = 1       // sets first element in array to 1
i := a[0]      // gets first element in array
// i == 1
```
Arrays do not need to be initialized explicitly; the zero value of an array is a ready-to-use array whose elements are themselves zeroed:
```go
// a[2] == 0, the zero value of the int type
```

arrays in Go are values. An array variable denotes the entire array (not a pointer to the first array element as in C).
This means that when you assign or pass around an array value you will make a copy of its contents
(to avoid the copy you could pass a `pointer` to the array, but then that's a pointer to an array, not an array).

The array is sort of like a struct but with indexed rather than named fields; a fixed-size composite value.

specifying array literals.
```go
b := [2]string{"Penn", "Teller"}     // declare and initialize a string array of size 2 with predefined values "Penn", "Teller"
c := [...]string{"Penn", "Teller"}   // as above but compiler counts the size for you
```
Note, the type of `b` and `c` is `[2]string`

More example usage of arrays.
```go
package main

import "fmt"

func main() {
	var numbers [3]string   // declare a string array of size 3
	numbers[0] = "uno"
	numbers[1] = "dos"
	numbers[2] = "tres"
	fmt.Println(numbers)    // [uno dos tres]

	var colors = [4]string{"rojo", "gris", "azul", "verde"} // comma separated values
	fmt.Println(colors)                                     // [rojo gris azul verde]
	fmt.Println(colors[2])                                  // azul
	fmt.Println(len(colors))                                // 4
}
```

### Data Structures (Slices)
https://go.dev/blog/slices-intro

Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
Unlike arrays, slices are typed only by the elements they contain (not the number of elements). E.g. `[]int`.

To create an empty slice with non-zero length, use the builtin `make`.

grow Slices with `append`.

Copy Slices by creating a new empty slice and then using `copy()`.

Slice indexing works with left_idx inclusive, right_idx exclusive. E.g. `s[2:4]` includes elements at index 2,3.

make a slice of strings of length 3 (initially zero-valued). Grow it with `append`
```go
package main

import "fmt"

func main() {

  s := make([]string, 3)      // declare and initialize zero-valued slice of length 3
  fmt.Println("emp:", s)      // emp: [  ]

  // can set and get just like arrays.
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)       // set: [a b c]
  fmt.Println("get:", s[2])    // get: c
  fmt.Println("len:", len(s))  // len: 3

  // In addition to these basic operations, slices support several more that make them richer than arrays.
  // One is the builtin append, which returns a slice containing one or more new values. Note that we need
  // to accept a return value from append as we may get a new slice value.
  s = append(s, "d")           // grow slice
  s = append(s, "e", "f")      // grow slice
  fmt.Println("apd:", s)       // apd: [a b c d e f]

  // slices can also be copy’d.
  // Here we create an empty slice c of the same length as s and copy into c from s.
  c := make([]string, len(s))  // declare and initialize new string slice c with same size as s
  copy(c, s)                   // copy s into c
  fmt.Println("cpy:", c)       // cpy: [a b c d e f]

  // Slices support a “slice” operator with the syntax slice[low:high].
  // For example, this gets a slice of the elements s[2], s[3], and s[4].
  l := s[2:5]
  fmt.Println("sl1:", l)       // sl1: [c d e]

  l = s[:5]                    // slice up to (but excluding) s[5].
  fmt.Println("sl2:", l)       // sl2: [a b c d e]

  l = s[2:]                    // slices up from (and including) s[2]
  fmt.Println("sl3:", l)       // sl3: [c d e f]

  t := []string{"g", "h", "i"} // declare and initialize a variable for slice in a single line.
  fmt.Println("dcl:", t)

  // Slices can be composed into multi-dimensional data structures.
  // The length of the inner slices can vary, unlike with multi-dimensional arrays.
  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
```

`Slices` give the benefits of the array but scale dynamically.\
They are an abstraction layer built on top of Go's array type. They are more common than arrays.

Type specification for a slice is `[]T`, where `T` is the type of the elements of the slice.\
Slice type has no specified length.

A slice can be created with the built-in function `make`, which has signature
```go
func make([]T, len, cap) []T
```
where `T` stands for the element type of the slice to be created.\
`make` function takes a type, length, and optional capacity.  \
When called, `make` allocates an array and returns a slice that refers to that array

create slice.
```go
var s []byte            // declare slice of bytes
s = make([]byte, 5, 5)  // create slice object
// s == []byte{0, 0, 0, 0, 0}
```

create slice succinctly (omitting capacity)
```go
s := make([]byte, 5)    // make slice of bytes of length 5
```

can inspect size and capacity with `len()` and `cap()`
```go
len(s) == 5
cap(s) == 5
```

default zero value for slice is `nil`.\
`len()` and `cap()` both return `0` for `nil` slice.

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
x := [3]string{"Лайка", "Белка", "Стрелка"} // declare and initialize a string array of size 3
s := x[:]                                   // a slice referencing the storage of x
```

#### Slice internals

A slice is a descriptor of an array segment.\
It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).

slice `length` refers to the number of elements referenced to by the slice.\
slice `capacity` is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).

slicing does not copy the slice's data.\
It creates a new slice value that points to the original array.\
Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice:

```go
d := []byte{'r', 'o', 'a', 'd'} // declare and initialize a slice of bytes
e := d[2:]                      // slice e now points to the same underlying array as d 
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
```

can grow `s` to its capacity by slicing again:
```go
s := make([]byte, 5)            // declare and initialize a slice of bytes of length 5
s = s[2:4]                      // shrink length of slice. capacity unchanged.
s = s[:cap(s)]                  // expand length of slice. capacity unchanged.
// s == []byte{'a', 'm'}
```

increase the capacity of a slice by creating a new, larger slice and copying the contents of the original slice into it (note this can be done easier than is shown below)
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

To append data to the end of the slice (e.g. a slice of bytes), use `AppendByte`, which gives great fine control over the way the slice is grown.
Depending on the characteristics of the program, it may be desirable to allocate in smaller or larger chunks, or to put a ceiling on the size of a reallocation.
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

__A potential gotcha with copying slices__

*potential gotcha* re-slicing a slice doesn't make a copy of the underlying array. The full array will be kept in memory until it is no longer referenced. Be aware, the array you first declared still exists in memory.

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

Example usage of slices.
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"maleta", "ropa", "reloj"} // declare an initialize a slice of strings
	fmt.Println(things)                              // [maleta ropa reloj]

	things = append(things, "bolso")                 // append item(s) to existing slice
	fmt.Println(things)                              // [maleta ropa reloj bolso]

	things = append(things[1:])                      // to shorten a slice and release memory properly
	// things = append(things[1 : len(things)-1])    // same as above (defaulting)
	fmt.Println(things)                              // [ropa reloj bolso]

	heros := make([]string, 3, 3)                    // declare and initialize slice object (create underlying array), length 3, capacity 3
	// even though capacity is 3, slice capacity will grow dynamically as needed
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "spiderman"
	fmt.Println(heros)                               // [thor ironman spiderman]

	heros = append(heros, "deadpool")                // slice capacity will grow dynamically by 2x
	fmt.Println(heros)                               // [thor ironman spiderman deadpool]
	fmt.Println(len(heros))                          // 4
	fmt.Println(cap(heros))                          // 6

	myints := []int{4, 7, 3, 2, 9}                   // declare and intialize slice of ints
	isSorted := sort.IntsAreSorted(myints)           // returns bool if ints are sorted
	fmt.Println("Are ints sorted: ", isSorted)

	fmt.Println("BEFORE: ", myints)
	sort.Ints(myints)                                // sorts slice of ints
	fmt.Println("AFTER: ", myints)

}
```


### Data Structures (Maps)

Maps are Go’s built-in *associative data type* (sometimes called *hashes* or *dicts* in other languages).

Use `make` keyword when declaring and initializing a `map`.

Use `delete` to remove key value pairs in the `map`.

`_` is known as the blank identifier.

```go
package main

import "fmt"

func main() {

  // To create an empty map, use the builtin make: make(map[key-type]val-type).
  m := make(map[string]int) // map[key]val

  // Set key/value pairs using typical name[key] = val syntax.
  m["k1"] = 7
  m["k2"] = 13

  // print all key value pairs
  fmt.Println("map:", m) // map: map[k1:7 k2:13]

  // Get a value for a key with name[key].
  v1 := m["k1"]
  fmt.Println("v1: ", v1) // v1:  7

  fmt.Println("len:", len(m)) // len: 2

  // The builtin delete removes key/value pairs from a map.
  delete(m, "k2")
  fmt.Println("map:", m) // map: map[k1:7]

  // The optional second return value when getting a value from a map indicates if the key was present in the map.
  // This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
  // Here we didn’t need the value itself, so we ignored it with the blank identifier _.
  _, prs := m["k2"]
  fmt.Println("prs:", prs) // prs: false

  n := map[string]int{"foo": 1, "bar": 2} // declare and initialize a new map in the same line with this syntax.
  fmt.Println("map:", n)
}
```

What is the difference between `make` and `new`?

- `new`  -- only allocates memory   - no initialization of memory
- `make` -- allocate and initialize - non zeroed storage

```go
package main

import "fmt"

func main() {
	var score map[string]int // declare map with string keys, int values. 
                           // automatically uses built-in new keyword;
                           // allocating memory but not initializing memory

	score["mine"] = 89       // panic: assignment to entry in nil map
	// memory was allocated for you but not initialized for you
}
```

creating a map, adding key value pairs, getting values at keys, deleting key value pairs. 
```go
package main

import "fmt"

func main() {
	// (typically) use make to create map
	score := make(map[string]int) // declare and initialize map that will have string keys, int values
	score["mine"] = 89
	fmt.Println(score)            // map[mine:89]

	score["josh"] = 34
	score["ron"] = 23
	score["sam"] = 56
	score["vicky"] = 78
	fmt.Println(score)            // map[josh:34 mine:89 ron:23 sam:56 vicky:78]

	getRonScore := score["ron"]
	fmt.Println(getRonScore)      // 23

	delete(score, "vicky")        // remove key value
	fmt.Println(score)            // map[josh:34 mine:89 ron:23 sam:56]
}

```

### Range

`range` iterates over elements in a variety of data structures.

```go
package main

import "fmt"

func main() {

  nums := []int{2, 3, 4} // slice of ints

  // for each thing in range
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  // for idx, thing in range
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}
  // for key, value in map
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  // for key in map
  for k := range kvs {
    fmt.Println("key:", k)
  }

  ///// for idx, char in string ////
  // range on strings iterates over Unicode code points.
  // The first value is the starting byte index of the rune and the second the rune itself.
  // See Strings and Runes for more details.
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
```

### Functions

Functions are central in Go. We’ll learn about functions with a few different examples.

Explicit `return` statements are required for all functions.

Defining functions with names that start with lower-case letters are different from those that start with upper-case letters (likely affects visibility)!

No function overloading in Go. Use `...` in argument signature to specify a variadic function (accepts variable number of arguments).

```go
package main

import "fmt"

// MULTIPLE ARGUMENTS
// accepts two ints, returns one int
func plus(a int, b int) int {
  return a + b 
}

func plusPlus(a, b, c int) int {
  return a + b + c
}

func main() {

  res := plus(1, 2)               // invoke function and store result in res (declare initialize and infer type for res)
  fmt.Println("1+2 =", res)

  res = plusPlus(1, 2, 3)
  fmt.Println("1+2+3 =", res)
}
```

More examples.
```go
package main

import (
	"fmt"
)

func superman() {
	fmt.Println("I am superman")
}

func multiplyme(v1 int, v2 int) int {
	return v1 * v2
}

func multiplyme2(v1, v2 int) int { // syntactic sugar
	return v1 * v2
}

func multiplyanddivide(v1, v2 int) (int, int) { // syntactic sugar; multiple return values
	return v1 * v2, v1 / v2
}

// VARIADIC FUNCTIONS
// no function overloading in go; must specify the number of args
func adder(values ...int) int {
	sum := 0
	for _, num := range values {
		sum += num
	}
	return sum
}

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

func main() {
	superman()                                   // invoke function

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
```

### Functions (multiple return values)

Go has built-in support for *multiple return values*. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

```go
package main

import "fmt"

func vals() (int, int) {
  return 3, 7
}

func main() {

  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)
}
```

### Variadic Functions

__Variadic Function__: a function of indefinite arity, i.e., one which accepts a variable number of arguments.
Variadic functions can be called with any number of trailing arguments. For example, `fmt.Println` is a common variadic function.

```go
package main

import "fmt"

// accept a variable number of arguments
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {

  sum(1, 2)
  sum(1, 2, 3)

  // If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this
  nums := []int{1, 2, 3, 4}
  sum(nums...) 
}
```

### Functions (closures)

Go supports *anonymous functions*, which can form *closures*. Anonymous functions are useful when you want to define a function inline without having to name it.

```go
package main

import "fmt"

// function below 'intSeq()' returns another function, which is defined
// anonymously in the body of intSeq. The returned function closes over
// the variable i to form a closure
func intSeq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {

  nextInt := intSeq() // returns a function and stores it in nextInt
  // This function value captures its own i value, which will be updated
  // each time we call nextInt.

  fmt.Println(nextInt()) // 1
  fmt.Println(nextInt()) // 2
  fmt.Println(nextInt()) // 3

  newInt2 := intSeq()
  fmt.Println(newInt2()) // 1
}
```

### Functions (recursion)

```go
package main

import "fmt"

func fact(n int) int {
  if n == 0 { // base case
    return 1
  }
  return n * fact(n-1) // recursive case
}

func main() {
    fmt.Println(fact(7))

    // Closures can also be recursive, but this requires the closure
    // to be declared with a typed var explicitly before it’s defined.
    var fib func(n int) int

    fib = func(n int) int { 
      if n < 2 {
        return n
      }
      // Since fib was previously declared in main, Go knows which 
      // function to call with fib here.
      return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7)) // 13
}
```

### Pointers

Go supports *pointers*, allowing you to pass references to values and records within your program.

We’ll show how pointers work in contrast to values with 2 functions: `zeroval` and `zeroptr`.

`zeroval` has an `int` parameter, so arguments will be passed to it by value.
`zeroval` will get a copy of `ival` distinct from the one in the calling function.

`zeroptr` in contrast has an `*int` parameter, meaning that it takes an `int` pointer.
The `*iptr` code in the function body then dereferences the pointer from its memory 
address to the current value at that address. Assigning a value to a dereferenced pointer
changes the value at the referenced address.

`&i` gives the memory address of variable i, i.e. a pointer to i.

`*p` means give me the thing that pointer p points to.

```go
package main

import "fmt"

// argument passed in will be by value; i.e., distinct from the var in the calling function
func zeroval(ival int) {
  ival = 0
}

// argument passed in will be a pointer; *iptr defrerences the pointer, meaning *iptr refers
// to "the thing located at memory location iptr". This body of code will set the var
// at location iptr to 0 (will change the var in the calling function)
func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i)    // 1

  zeroval(i)
  fmt.Println("zeroval:", i)    // 1

  zeroptr(&i)                   // &i gives the memory address of i, i.e. a pointer to i
  fmt.Println("zeroptr:", i)    // 0

  fmt.Println("pointer:", &i)   // prints the memory address: 0x42131100
}
```

More code exemplifying pointers
```go
package main

import (
	"fmt"
)

func main() {
	var p *int // declare a pointer to an int (pointer is not storing anything)
             // default value (zero-d value) is nil

	if p != nil { // evals to false
		fmt.Println("P is having a value: ", *p) // *p refers to what pointer p points to, which is nil
	} else {
		fmt.Println("Watchout for nil values")
	}
}
```

More code exemplifying pointers
```go
package main

import (
	"fmt"
)

func main() {
	var lifebooster float64 = 99.2    // declare and initialize a float64
	lifeboosterRef := &lifebooster    // lifeboosterRef holds reference to memory location
	                                  // can also be done like so
                                    // var lifeboosterRef = &lifebooster 

	fmt.Println(lifebooster)          // prints 99.2
	fmt.Println(lifeboosterRef)       // prints 0xc0000a6058
	fmt.Println(*lifeboosterRef)      // prints 99.2, the thing located at 0xc0000a6058

  lifebooster = lifebooster * 2.2

  fmt.Println(lifebooster)          // prints 218.24000000000004
	fmt.Println(lifeboosterRef)       // prints 0xc0000a6058
	fmt.Println(*lifeboosterRef)      // prints 218.24000000000004, the thing located at 0xc0000a6058
}
```

### Strings and Runes

A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a `rune` - it’s an integer that represents a Unicode code point. This Go blog post (https://go.dev/blog/strings) is a good introduction to the topic.

`rune` literals are enclosed with single quotes `''`; strings, with double quotes `""`

```go
package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {

  const s = "สวัสดี" //  Go string literals are UTF-8 encoded text.

  fmt.Println("Len:", len(s)) // Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
                              // Len: 18

  // Indexing into a string produces the raw byte values at each index.
  // This loop generates the hex values of all the bytes that constitute the code points in s.
  for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i]) // e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
  }
  fmt.Println()

  // To count how many runes are in a string, we can use the utf8 package. 
  // Note that the run-time of RuneCountInString dependes on the size of the string,
  // because it has to decode each UTF-8 rune sequentially. 
  // Some Thai characters are represented by multiple UTF-8 code points, so the result of this count may be surprising.
  fmt.Println("Rune count:", utf8.RuneCountInString(s)) // Rune count: 6

  // A range loop handles strings specially and decodes each rune along with its offset in the string.
  for idx, runeValue := range s {
    fmt.Printf("%#U starts at %d\n", runeValue, idx)
  }
  // U+0E2A 'ส' starts at 0
  // U+0E27 'ว' starts at 3
  // U+0E31 'ั' starts at 6
  // U+0E2A 'ส' starts at 9
  // U+0E14 'ด' starts at 12
  // U+0E35 'ี' starts at 15

  // can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
  fmt.Println("\nUsing DecodeRuneInString")
  for i, w := 0, 0; i < len(s); i += w {
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at %d\n", runeValue, i)
    w = width

    examineRune(runeValue) // e.g. passing a runevalue to a function
  }
}

func examineRune(r rune) {

  if r == 't' {
      fmt.Println("found tee")
  } else if r == 'ส' {
      fmt.Println("found so sua")
  }
}
```

### Structs

Go’s *structs* are typed collections of fields. They’re useful for grouping data together to form records.

Structs are defined with the keyword `type`.

It is idiomatic to define constructor functions that create structs for you (rather than create structs inline).

Upper and lower casing of structs affects the visibility. Upper case (capitalized) structs are __Capitalized Identifiers__. The capital letter indicates that this is an exported identifier and is available outside the package. __Non-capitalized identifiers__ are not exported. The lowercase indicates that the identifier is not exported and will only be accessed from within the same package.

Upper and lower casing of struct fields affects visibility of those fields in the same way. struct fields that begin with capital letters are exported and those that begin with lower case letters are not exported.

```go
package main

import "fmt"

// STRUCT
// This person struct type has name and age fields
// lowercase lettering (person) indicates the struct will only be visible in package main
type person struct {
  name string // lower case field indicates it is not exported
  age  int
}

// CONSTRUCTOR FUNCTION FOR STRUCT
// newPerson constructs a new person struct with the given name.
// You can safely return a pointer to a local variable as a local variable will survive the scope of the function.
func newPerson(name string) *person {
  p := person{name: name}
  p.age = 42
  return &p // return pointer
}

func main() {

  fmt.Println(person{"Bob", 20})              // creates a new struct
                                              // {Bob 20}

  fmt.Println(person{name: "Alice", age: 30}) // can name the fields when initializing a struct
                                              // {Alice 30}

  fmt.Println(person{name: "Fred"})           // omitted fields will be zero-valued
                                              // {Fred 0}

  fmt.Println(&person{name: "Ann", age: 40})  // An & prefix yields a pointer to the struct.
                                              // &{Ann 40}

  fmt.Println(newPerson("Jon"))               // It’s idiomatic to encapsulate new struct creation in constructor functions
                                              // &{Jon 42}

  s := person{name: "Sean", age: 50} 
  fmt.Println(s.name)                         // Access struct fields with a dot. 
                                              // Sean

  sp := &s // You can also use dots with struct pointers - the pointers are automatically dereferenced.
  fmt.Println(sp.age)                         // 50

  sp.age = 51 // Structs are mutable.
  fmt.Println(sp.age)                         // 51
}
```

More example usage of structs.
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


### Methods (functions for structs)

Go supports *methods* defined on struct types.

```go
package main

import "fmt"

type rect struct {
  width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
  return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver
func (r rect) perim() int {
  return 2*r.width + 2*r.height
}

func main() {
  r := rect{width: 10, height: 5}

  fmt.Println("area: ", r.area())  // area:  50
  fmt.Println("perim:", r.perim()) // perim: 30

  // Go automatically handles conversion between values and pointers for method calls.
  // use a pointer receiver type to 
  //    (1) avoid copying on method calls or
  //    (2) to allow the method to mutate the receiving struct.
  rp := &r
  fmt.Println("area: ", rp.area())  // area:  50
  fmt.Println("perim:", rp.perim()) // perim: 30
  // in other words, coder ought to pass 
}
```

### Interfaces

*Interfaces* are named collections of method signatures.

```go
package main

import (
    "fmt"
    "math"
)

// define basic interface
type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// implement an interface in Go by just implementing all the methods in the interface. Here we implement geometry on rects.
// have rect implement geometry interface
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// have circle implement geometry interface
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// can call methods in the named interface
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
```

### Struct Embedding (Inheritance?)

Go supports embedding of structs and interfaces to express a more seamless composition of types.
This is not to be confused with //go:embed which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.

a *container* struct *embeds* a *base* struct. 

```go
package main

import "fmt"

type base struct {
  num int
}

func (b base) describe() string {
  return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name.
type container struct {
  base
  str string
}

func main() {

  // When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
  co := container{
    base: base{
      num: 1,
    },
    str: "some name",
  }

  fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // We can access the base’s fields directly on co, e.g. co.num.

  fmt.Println("also num:", co.base.num) // Alternatively, we can spell out the full path using the embedded type name.

  // Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was embedded from base directly on co.
  fmt.Println("describe:", co.describe()) // i.e. this is a super method

  type describer interface {
    describe() string
  }

  // Embedding structs with methods may be used to bestow interface implementations
  // onto other structs. Here we see that a container now implements the describer
  // interface because it embeds base.
  var d describer = co
  fmt.Println("describer:", d.describe())
}
```

### Generics

Starting with version 1.18, Go has added support for *generics*, also known as *type parameters*.
Note, the `interface{}` constraint (type) is equivalent to the `any` constraint (type); `interface{}`
is an empty interface (interface with no methods defined), and all objects trivially implement 
this interface. The `any` constraint in Go is an alias for the `interface{}` constraint. 

The `comparable` constraint means values of this type can be compared with the `==` and `!=` operators.

```go
package main

import "fmt"

// As an example of a generic function, MapKeys takes a map of any type and returns a slice of
// its keys. This function has two type parameters - K and V; K has the comparable constraint,
// meaning that we can compare values of this type with the == and != operators. This is required
// for map keys in Go. V has the any constraint, meaning that it’s not restricted in any way
// (any is an alias for interface{}).
func MapKeys[K comparable, V any](m map[K]V) []K { // returns slice of keys in map
  r := make([]K, 0, len(m))
  for k := range m {
      r = append(r, k)
  }
  return r
}

// As an example of a generic type, List is a singly-linked list with values of any type.
type List[T any] struct {
  head, tail *element[T]
}

type element[T any] struct {
  next *element[T]
  val  T
}

// We can define methods on generic types just like we do on regular types, but we have 
// to keep the type parameters in place. The type is List[T], not List.
func (lst *List[T]) Push(v T) {
  if lst.tail == nil {
      lst.head = &element[T]{val: v} // creates element struct, then assigns pointer to element to lst.head
      lst.tail = lst.head // have the pointer to tail be the same as pointer to the head
  } else {
      lst.tail.next = &element[T]{val: v} // create element struct, assign it to be the tail's next
      lst.tail = lst.tail.next // update tail to now be the newly appended element.
  }
}

func (lst *List[T]) GetAll() []T {
  var elems []T // declares slice of generic Type
  for e := lst.head; e != nil; e = e.next { // initial statement / condition statement / update
      elems = append(elems, e.val) // grow slice
  }
  return elems
}

func main() {
  var m = map[int]string{1: "2", 2: "4", 4: "8"}

  // When invoking generic functions, we can often rely on type inference.
  // Note that we don’t have to specify the types for K and V when calling
  // MapKeys - the compiler infers them automatically.
  fmt.Println("keys:", MapKeys(m)) // keys: [4 1 2]

  // … though could also specify them explicitly.
  _ = MapKeys[int, string](m) // returns slice of keys; not used though

  lst := List[int]{} // empty linked list
  lst.Push(10)
  lst.Push(13)
  lst.Push(23)
  fmt.Println("list:", lst.GetAll()) // list: [10 13 23]
}
```

### Errors
https://go.dev/blog/error-handling-and-go

In Go it’s idiomatic to communicate errors via an explicit, separate return value.
This contrasts with the exceptions used in languages like Java and Ruby and the 
overloaded single result / error value sometimes used in C. Go’s approach makes 
it easy to see which functions return errors and to handle them using the same 
language constructs employed for any other, non-error tasks.

```go
package main

import (
  "errors"
  "fmt"
)

func f1(arg int) (int, error) { // error is last return value by convention; have type error (a built-in interface)
  if arg == 42 {

    return -1, errors.New("can't work with 42") // errors.New constructs a basic error value with the given error message.

  }

  return arg + 3, nil // return nil on good results
}

// It’s possible to use custom types as errors by implementing the Error() method on them.
// Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
type argError struct {
  arg  int
  prob string
}

func (e *argError) Error() string { // Error() method needs to be implemented when making custom error
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
  if arg == 42 {

    return -1, &argError{arg, "can't work with it"} // In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
  }
  return arg + 3, nil
}

func main() {

  // The two loops below test out each of our error-returning functions. 
  // Note that the use of an inline error check on the if line is a common idiom in Go code.
  for _, i := range []int{7, 42} { // iterate through int slice of two elements
    if r, e := f1(i); e != nil {  // inline error check on same line as if statement
      fmt.Println("f1 failed:", e)
    } else {
      fmt.Println("f1 worked:", r)
    }
  }
  for _, i := range []int{7, 42} {
    if r, e := f2(i); e != nil {
      fmt.Println("f2 failed:", e)
    } else {
      fmt.Println("f2 worked:", r)
    }
  }

  // If you want to programmatically use the data in a custom error,
  // you’ll need to get the error as an instance of the custom error type via type assertion.

  // e.(*argError) is a type assertion that casts the value e to *argError type. This is the 
  // type f2() returns on error - it is a pointer to an argError struct, which implements the
  // error interface.

  // This type assertion evaluates multi-valued to (ae,ok), where ae is the *argError typed value, 
  // if successful, and ok is a boolean letting you know if it was successful.

  // if statements in go can be separated into an initial assignment part, then a semicolon, 
  // then a boolean condition to evaluate to determine the branch.

  _, e := f2(42)
  if ae, ok := e.(*argError); ok { // means: try to cast e to *argError, if that is successful (do if block).
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
  }
}
```

### String Functions

The standard library’s `strings` package provides many useful string-related functions. Here are some examples to give you a sense of the package.

```go
package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

  p("Contains:  ", s.Contains("test", "es"))        // Contains:   true
  p("Count:     ", s.Count("test", "t"))            // Count:      2
  p("HasPrefix: ", s.HasPrefix("test", "te"))       // HasPrefix:  true
  p("HasSuffix: ", s.HasSuffix("test", "st"))       // HasSuffix:  true
  p("Index:     ", s.Index("test", "e"))            // Index:      1
  p("Join:      ", s.Join([]string{"a", "b"}, "-")) // Join:       a-b
  p("Repeat:    ", s.Repeat("a", 5))                // Repeat:     aaaaa
  p("Replace:   ", s.Replace("foo", "o", "0", -1))  // Replace:    f00
  p("Replace:   ", s.Replace("foo", "o", "0", 1))   // Replace:    f0o
  p("Split:     ", s.Split("a-b-c-d-e", "-"))       // Split:      [a b c d e]
  p("ToLower:   ", s.ToLower("TEST"))               // ToLower:    test
  p("ToUpper:   ", s.ToUpper("test"))               // ToUpper:    TEST
}
```

### String Formatting

Go offers excellent support for string formatting in the `printf` tradition.
Here are some examples of common string formatting tasks.

Go offers several printing *verbs* designed to format general Go values.

```go
package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

  p := point{1, 2}
  fmt.Printf("struct1: %v\n", p)  // %v prints instance of struct (or the value)
                                  // struct1: {1 2}

  fmt.Printf("struct2: %+v\n", p) // %+v prints instance of struct with the struct’s field names (if a struct).
                                  // struct2: {x:1 y:2}

  fmt.Printf("struct3: %#v\n", p) // %#v variant prints a Go syntax representation of the value,
                                  // i.e. the source code snippet that would produce that value.
                                  // struct3: main.point{x:1, y:2}

  fmt.Printf("type: %T\n", p)     // %T prints type
                                  // type: main.point

  fmt.Printf("bool: %t\n", true)  // %t prints formatted booleans
                                  // bool: true

  fmt.Printf("int: %d\n", 123)    // %d prints standard, base-10 formatting.
                                  // int: 123

  fmt.Printf("bin: %b\n", 14)     // %b prints binary representation
                                  // bin: 1110

  fmt.Printf("char: %c\n", 33)    // %c prints character corresponding to the given integer.
                                  // char: !
  
  fmt.Printf("hex: %x\n", 456)    // %x prints hex encoding
                                  // hex: 1c8

  fmt.Printf("float1: %f\n", 78.9)         // %f prints basic decimal formatting
                                           // float1: 78.900000

  fmt.Printf("float2: %e\n", 123400000.0) // %e prints float in scientific notation.
                                          // float2: 1.234000e+08

  fmt.Printf("float3: %E\n", 123400000.0) // %E prints float in scientific notation.
                                          // float3: 1.234000E+08

  fmt.Printf("str1: %s\n", "\"string\"")  // %s prints basic string
                                          // str1: "string"

  fmt.Printf("str2: %q\n", "\"string\"")  // %q to double-quote strings as in Go source.
                                          // str2: "\"string\""

  fmt.Printf("str3: %x\n", "hex this")    // %x renders the string in base-16, with two output characters per byte of input
                                          // str3: 6865782074686973

  fmt.Printf("pointer: %p\n", &p)         // %p prints pointer representation
                                          // pointer: 0xc0000ba000

  fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // %6d prints number with width 6 (right justified, padded w spaces).
                                             // width1: |    12|   345|

  fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) // %6.2f prints number with width 6, decimal precision of 2 (width.precision syntax).
                                                   // width2: |  1.20|  3.45|

  fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // - flag does a left-justify
                                                     // width3: |1.20  |3.45  |

  fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")  // for basic right-justified width, e.g. to ensure they align in table-like output.
                                                 // width4: |   foo|     b|

  fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b") // - flag does a left-justify
                                                  // width5: |foo   |b     |

  // So far we’ve seen Printf, which prints the formatted string to os.Stdout.
  // Sprintf formats and returns a string without printing it anywhere.
  s := fmt.Sprintf("sprintf: a %s", "string")
  fmt.Println(s) // sprintf: a string

  // You can format+print to io.Writers other than os.Stdout using Fprintf.
  fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // io: an error
}
```

### EXAMPLE PROGRAM

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





