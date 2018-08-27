package main

import (
	"strconv"
	"fmt"
)

func main(){
	number := 100
	numberStr := strconv.Itoa(number)
	fmt.Println(numberStr)

	numberFloat := 23445221.1223356
	numberFloatStr := strconv.FormatFloat(numberFloat, 'f',-1,64 )
	fmt.Println(numberFloatStr)

}