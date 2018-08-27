package main

import (
	"time"
	"fmt"
)

func main(){
	first := time.Date(2017, 1,1,0,0,0,0,time.UTC)
	second := time.Date(2018, 1,1,0,0,0,0,time.UTC)

	difference := second.Sub(first)
	fmt.Printf("Difference %v", difference)
}
