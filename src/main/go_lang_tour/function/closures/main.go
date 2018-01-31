package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int, 1)

	channel <- 1

	go up(channel)

	time.Sleep(1 * time.Second)

	fmt.Println(<-channel)

}

func up(channel chan int) {

	mid := <-channel + 1
	channel <- mid

}
