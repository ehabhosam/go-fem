package functions

func Hello() {
	print("Hello from the functions package\n")
}

// simple go function
func Add(a int, b int) int {
	return a + b
}

// functions can return more than one value
func Swap(a, b string) (string, string) {
	return b, a
}

// functions can return named values
func NamedReturn(a, b int) (x int) {
	x = a + b
	return
} // not really useful

// functions can recieve reference to a variable
func Zero(x *int) {
	*x = 0
}
