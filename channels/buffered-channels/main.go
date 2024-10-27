package main

import (
	"fmt"
	"time"
)

func listenToChannel(ch chan int) {

	for {
		i := <-ch
		fmt.Println("Got", i)

		time.Sleep(1 * time.Second)
	}

}

func main() {

	ch := make(chan int, 10)

	go listenToChannel(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending", i)
		ch <- i

	}

	close(ch)

}
