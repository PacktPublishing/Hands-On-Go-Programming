package main

import "fmt"

func main() {
	nameAges := map[string]int{
		"Tarik":   32,
		"Michael": 30,
		"Jon":     25,
		"Jessica" : 20,
	}

	if value, exists := nameAges["Jessica"]; exists{
		fmt.Println(value)
	}else {
		fmt.Println("Jessica cannot be found")
	}

}
