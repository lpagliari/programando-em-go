package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 3)
	go produce(channel)

	// would print `1 2 3 0`
	// fmt.Println(<- channel, <- channel, <- channel, <- channel)

	for value := range channel {
		fmt.Println(value)
	}
}

func produce(channel chan int) { // can also be `chan<- int`
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)
}
