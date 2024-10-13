package threading

import (
	"fmt"
	"time"
)

// go routines
// kind of threads but lighter
// ofc threads runs in same process

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second / 2)
	}
}

func SpinGoRoutines() {
	go printMessage("ehab")     // runs in a seprarate goroutine (thread)
	go printMessage("is goat")  // runs in a seprarate goroutine
	printMessage("hahahahahah") // runs in the main goroutine

	// if we run all in go routines, the main goroutine will
	// finish before the other goroutines execute
	// and we will get no any output, so we will need
	//to wait by blocking with a loop or time.Sleep for example
}
