package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	files, err := ioutil.ReadDir("hello")

	if err != nil{
		panic(nil)

	}

	for _,f := range files{
		fmt.Println(f.Name())
	}
}
