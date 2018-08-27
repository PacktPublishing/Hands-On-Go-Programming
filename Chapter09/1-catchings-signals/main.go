package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main(){
	signals := make (chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func (){
		sig := <- signals
		fmt.Println(sig)
		fmt.Println("Signal captured and processed...")
		done <- true
	}()

	fmt.Println("Waiting for signal")
	<-done
	fmt.Println("Exiting the application...")
}
