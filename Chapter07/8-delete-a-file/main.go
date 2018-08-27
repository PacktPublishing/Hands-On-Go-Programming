package main

import "os"

func main() {
	err := os.Remove("new.txt")

	if err != nil{
		panic(err)
	}
}
