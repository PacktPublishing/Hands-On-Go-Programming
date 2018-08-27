package main

import "fmt"

func main(){
	numbers := []int{1, 5, 3, 6, 2, 10, 8}

	for _,value := range numbers{
		// fmt.Printf("Index: %v and Value: %v\n", index, value)
		fmt.Println(value)
	}
}
