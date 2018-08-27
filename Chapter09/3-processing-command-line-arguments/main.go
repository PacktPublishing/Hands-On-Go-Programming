package main

import (
	"os"
	"fmt"
)

func main(){
	realArgs := os.Args[1:]

	if len(realArgs) == 0{
		fmt.Println("Please pass an argument.")
		return
	}

	if realArgs[0] == "a"{
		writeHelloWorld()
	}else if realArgs[0] == "b"{
		writeHelloMars()
	}else{
		fmt.Println("Please pass a valid argument.")
	}
}

func writeHelloWorld(){
	fmt.Println("Hello, World")
}

func writeHelloMars(){
	fmt.Println("Hello, Mars")
}
