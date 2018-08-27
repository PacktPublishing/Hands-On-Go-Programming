package main

import (
	"time"
	"fmt"
)

func main(){
	str := "2018-08-08T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"

	t,err := time.Parse(layout, str)

	if err != nil{
		fmt.Println(err)
	}	

	fmt.Println(t.String())
}
