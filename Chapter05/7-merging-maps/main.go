package main

import "fmt"

func main(){
	map1 := map[string]int {
		"Michael":10,
		"Jessica":20,
		"Tarik":33,
		"Jon": 22,
		}

	fmt.Println(map1)

	map2 := map[string]int {
		"Lord":11,
		"Of":22,
		"The":36,
		"Rings": 23,
	}

	for key, value := range map2{
		map1[key] = value
	}

	fmt.Println(map1)
}