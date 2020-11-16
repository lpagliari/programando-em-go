package main

import (
	"fmt"
	"time"
)

func main() {
  go print(2)
  print(3)
}

func print(n int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", n)
		time.Sleep(200 * time.Millisecond)
	}
}
