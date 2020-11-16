package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	go run(channel)

	fmt.Println("Waiting...")
	done := false
	for !done {
		select {
		case done = <- channel:
			fmt.Println("Completed!")
		// case <- time.After(2 * time.Second):
		case <- After(2 * time.Second):
			fmt.Println("Timeout!")
			done = true
		}
	}
}

func run(channel chan<- bool) {
	time.Sleep(5 * time.Second) // change here to `1` and there's no timeout
	channel <- true
}

// pseudo-implementation of `time.After()`:
func After(duration time.Duration) <-chan time.Time {
	c := make(chan time.Time)
	go runTimer(duration, c)
	return c
}

func runTimer(duration time.Duration, c chan time.Time) {
	time.Sleep(duration)
	c <- time.Now()
}
