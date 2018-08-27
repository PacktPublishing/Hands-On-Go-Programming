package main

import (
	"strings"
	"fmt"
)

func main(){
	helloWorld := "Hello, World. How are you World, I am good, thanks World."
	helloMars := strings.Replace(helloWorld, "World", "Mars", -1)
	fmt.Println(helloMars)
}