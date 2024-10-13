package main

func printFunctions() {
	print("hello from the functions.go file\n")
	print("I also know the goat!! ", goat, "\n")
}

func ExportedFunction() {
	// so if you import the package somewhere else, you can access this funnction
	// it's gonna be "pkgname.ExportedFunction()"
	print("I am an exported function, I start with a capital letter.\n")
}
