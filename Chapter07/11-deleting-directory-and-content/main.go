package main

import (
	"os"
	"fmt"
)

func main(){
	err := os.RemoveAll("hello")
	if err != nil{
		fmt.Println(err)
	}
}
