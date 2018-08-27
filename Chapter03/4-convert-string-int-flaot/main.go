package main

import (
	"strconv"
	"fmt"
)

func main(){


	numberFloat := "2.2"
	valueFloat, errFloat := strconv.ParseFloat(numberFloat, 64)
	if errFloat != nil {
		fmt.Print("Error happened.")
	} else {
		if valueFloat == 2.2 {
			fmt.Println("Success")
		}
	}
}
