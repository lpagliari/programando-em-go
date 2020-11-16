package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	// if we don't have the `go` here there would be a deadlock. Why?
	go produce(channel)

	// if there's another call to produce values, it is ignored -- probably because the buffer is full?
	// go produce(channel)

	value := <- channel
	fmt.Println(value)
}

func produce(channel chan int) {
	channel <- 33
}
