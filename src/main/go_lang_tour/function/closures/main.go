package main

import (
	"fmt"
)

func main() {

	base := 1
	channel := make(chan int)

	go func(channel chan int) {
		base = base + 1
		channel <- base
	}(channel)

	fmt.Println(<-channel)
}
