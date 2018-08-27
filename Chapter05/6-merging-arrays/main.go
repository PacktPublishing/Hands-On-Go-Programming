package main

import "fmt"

func main(){
	items1 := []int{5,3, 3,5}
	items2 := []int{1,2}

	result := append(items1, items2...)
	fmt.Println(result)
}
