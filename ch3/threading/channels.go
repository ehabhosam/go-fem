package threading

import "time"

// channels are how goroutines communicate
// channels are typed
// you can send and receive values from channels

func PrintMessageNotifyWhenDone(msg string, done chan string) {
	for i := 0; i < 5; i++ {
		println(msg)
		time.Sleep(time.Second / 2)
	}
	done <- "done !" + msg
}
