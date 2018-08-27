package main

import (
	"strings"
	"fmt"
)

func main(){
	helloWorld := "hello world, how are you today!"
	helloWorldtitle := strings.Title(helloWorld)
	fmt.Println(helloWorldtitle)

	helloWorldUpper := strings.ToUpper(helloWorld)
	fmt.Println(helloWorldUpper)
}
