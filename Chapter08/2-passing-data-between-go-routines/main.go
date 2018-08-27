package main

import "fmt"

func main(){
	nameChannel := make(chan string, 5)
	done  := make(chan string)
	go func(){
		names := []string {"tarik", "michael", "gopi", "jessica"}
		for _, name := range names {
			// doing some operation
			fmt.Println("Processing the first stage of: " + name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	go func(){
		for name := range nameChannel{
			fmt.Println("Processing the second stage of: " + name)
		}
		done <- ""
	}()

	<-done
}
