package main

import (
	"fmt"
	"frontendmasters/learn-go/ch2/calculator"
	"frontendmasters/learn-go/ch2/files"
	"os"
	"slices"
)

func run_calculator() {
	var operation string
	var n1, n2 int
	fmt.Println("Welcome to go calculator.")
	fmt.Println("-------------------------")
	fmt.Println("what operation do u wanna perform?")
	fmt.Println("allowed operations: +, -, *, /")

	defer fmt.Println("---- program ended ----")

	fmt.Scanf("%s", &operation)

	fmt.Println("enter the first number")
	fmt.Scanf("%d", &n1)
	fmt.Println("enter the second number")
	fmt.Scanf("%d", &n2)

	var allowdOperation = []string{"+", "-", "*", "/"}

	if slices.Contains(allowdOperation, operation) {
		result := calculator.Calc(operation, n1, n2)
		fmt.Println("The result is: ", result)
	} else {
		fmt.Println("You entered a wrong operation.")
	}
}

func read_file() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("something went wrong, while reading the current directory")
	}
	content, err := files.ReadTextFile(rootDir + "/content.txt")
	if err != nil {
		fmt.Println("something went wrong, while reading the file")
	} else {
		fmt.Println(content)
	}
}

func write_file() {
	rootDir, err := os.Getwd()
	// let's double the current content of the file
	if err != nil {
		fmt.Println("something went wrong, while reading the current directory")
	}
	content, err := files.ReadTextFile(rootDir + "/content.txt")
	new_content := fmt.Sprintf("content:\n %s\n new content:\n %s%s", content, content, content)
	if err != nil {
		fmt.Println("something went wrong, while reading the current directory")
	}
	err = files.WriteToTextFile(rootDir+"/content.txt", new_content)
	if err != nil {
		fmt.Println("something went wrong, while writing to the file")
	}
}

func main() {
	// run_calculator()
	// read_file()
	write_file()
}
