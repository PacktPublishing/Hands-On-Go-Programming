package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	hello := "Hello, World Again"
	err := ioutil.WriteFile("hello_world", []byte(hello), 0644)
	if err != nil{
		fmt.Println(err)
	}
}
