package main

import (
	"fmt"
	"frontendmasters/learn-go/ch3/threading"
	my_types "frontendmasters/learn-go/ch3/types"
	"time"
)

func main() {
	// my_types.Test()
	// my_types.TestStructs()
	// test_structs()
	// test_instructor()
	goroutines()
}

func test_structs() {
	max := my_types.Instructor{}
	// max := new(my_types.Instructor) // same as above
	max.FirstName = "Max"
	max.LastName = "Mukendi"
	// max.id = 1 // does not work because id is private
	fmt.Println(max.String())
	fmt.Println(max.GetScore())
}

func printSomething(p my_types.Printable) {
	fmt.Println(p.String())
}

func test_instructor() {
	max := my_types.NewInstructor("Max", "Mukendi")
	kyle := my_types.Instructor{FirstName: "Kyle", LastName: "Simpson"}

	goCourse := my_types.Course{
		Id:         1,
		Name:       "Go Programming",
		Slug:       "go-programming",
		Legacy:     false,
		Duration:   3.5,
		Instructor: max,
	}

	reactCourse := my_types.Course{
		Id:         2,
		Name:       "React Programming",
		Slug:       "react-programming",
		Legacy:     false,
		Duration:   4.5,
		Instructor: kyle,
	}

	workshop := my_types.Workshop{
		Course: goCourse,
		Date:   time.Now(),
	}

	fmt.Println(goCourse.String())
	fmt.Println(reactCourse.String())
	fmt.Println(workshop.String()) // method from Course was embedded in Workshop

	// we can do it with interfaces
	printSomething(goCourse)
	printSomething(workshop) // workshop is not a Course but it implements the Printable interface
}

func goroutines() {
	// threading.SpinGoRoutines()
	var channel = make(chan string)
	go threading.PrintMessageNotifyWhenDone("ehab", channel)
	go threading.PrintMessageNotifyWhenDone("is goat", channel)

	res := <-channel // will block till message reaches the channel
	fmt.Println(res)
	res = <-channel // will block again
	fmt.Println(res)
}
