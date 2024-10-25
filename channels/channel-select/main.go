package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(time.Second * 6)
		ch <- "from server1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(time.Second * 3)
		ch <- "from server2"
	}
}

func main() {

	fmt.Println("Select with channels")
	fmt.Println("---------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		// if more than one channel is ready, select will choose one randomly
		select {
		case s1 := <-channel1:
			fmt.Println("Case 1: ", s1)
		case s2 := <-channel1:
			fmt.Println("Case 2: ", s2)
		case s3 := <-channel2:
			fmt.Println("Case 3: ", s3)
		case s4 := <-channel2:
			fmt.Println("Case 4: ", s4)
		}
	}

}
