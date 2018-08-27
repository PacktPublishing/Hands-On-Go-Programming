package main

import (
	"time"
	"fmt"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func(){
		time.Sleep(1*time.Second)
		channel1 <- "Hello from channel 1"
	}()

	go func(){
		time.Sleep(1 * time.Second)
		channel2 <- "Hello from channel 2"
	}()

	var result string
	for {
		select {
		case result = <-channel1:
			fmt.Println(result)
		case result = <-channel2:
			fmt.Println(result)
		}
		case <-quit:
			return
	}
}
