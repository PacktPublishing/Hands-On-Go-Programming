package main

import (
	"fmt"
	"errors"
)

func main(){
	_, err := doSomething()
	if err != nil {
		fmt.Println(err)
	}
}

func doSomething() (string,error) {
	return "", errors.New("Something happened.")
}



