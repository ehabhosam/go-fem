// we are in the main package (which is the entry package of the module)
package main

import (
	"fmt"
	"frontendmasters/learn-go/ch1/pkg"
)

// global variable
var goat = "ehab hosam"

func init() {
	// this function runs when the package is imported
	// it's like a constructor
	// it runs before the main function
}

func init() {
	// you can have multiple init functions
	// they run in the order they appear
}

func main() {
	print("I found the goat!! ", goat, "\n")

	// Declaring variables
	var a int = 5 // this
	var x = 5     // is same as this
	y := 7        // as same as this
	print(a, x, y)

	// a constast should really be a constant
	// string, boolean, or number ONLY
	const f = 5 // constant

	// Data types
	// - string values: string
	// - integer values: int8, int16, int32, int64, uint8, uint16, uint32, uint64
	// - float values: float32, float64
	// - boolean values: bool

	// Boolean Operators: normal stuff

	// call function from same package
	// same folder files see each other
	printFunctions()

	// in production, don't use print
	// print is good for development debugging
	// use the logger instead from 'fmt' package
	fmt.Println("This is a good log.")

	// this comes from a differnt package
	fmt.Println(pkg.SomeData)

	// so the heirarchy is: (first is the highest)
	// - workspace
	// - module
	// - package
	// - file

	// Collections
	// - Arrays: fixed size, same type [5]int
	// - Slices: dynamic size, same type []int
	// - Maps: key-value pairs map[string]int
	// - Generics: from 1.18 (we have it now)
}
