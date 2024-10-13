package error

import "fmt"

// there is no try-catch in go
// but you can use the defer keyword to run a function after the current function ends
// it will execute even if something goes wrong
func SomeFunction() {
	defer fmt.Println("I am done")
	fmt.Println("I am doing something")
	panic("I am panicking")
}
