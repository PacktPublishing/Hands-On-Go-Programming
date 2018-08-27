package main

import (
	"fmt"
	"time"
)

func main() {

	nameChannel := make(chan string)

	go func() {
		names := []string{"tarik", "john", "michael", "jessica"}

		for _, name := range names {
			time.Sleep(1 * time.Second)
			fmt.Println(name)
		}
		nameChannel <- ""
	}()

	<-nameChannel
}
