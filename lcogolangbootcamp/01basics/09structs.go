package main

// import (
// 	"fmt"
// )

// // define a struct with the word "type"
// type User struct { // name of struct is upper case
// 	Name  string // attributes are upper case too
// 	Email string
// 	age   int
// }

// func main() {
// 	rob := User{"rob", "rob@lco.dev", 34}
// 	fmt.Printf("%v\n", rob) // {rob rob@lco.dev 34}
// 	// use %+v to obtain more details on the object
// 	fmt.Printf("%+v\n", rob) // {Name:rob Email:rob@lco.dev age:34}

// 	var sam = new(User) // remember only allocates memory. doesn't initialize
// 	sam.Name = "sam"
// 	sam.Email = "sam@lco.dev"
// 	sam.age = 22
// 	fmt.Printf("%v\n", sam)  // &{sam sam@lco.dev 22}
// 	fmt.Printf("%+v\n", sam) // &{Name:sam Email:sam@lco.dev age:22}

// 	// remember the 'new' only allocates memory. it doesn't initialize.
// 	// The & in &{sam sam@lco.dev 22} indicates a reference you are getting back.
// 	// The & symbol points to the address of the stored value

// 	var tobby = &User{"tobby", "tobby@lco.dev", 22} // less common
// 	fmt.Printf("%+v\n", tobby)                      // &{Name:tobby Email:tobby@lco.dev age:22}

// }
