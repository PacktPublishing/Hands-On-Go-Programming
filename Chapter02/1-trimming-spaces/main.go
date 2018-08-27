package main

import (
	"fmt"
	"strings"
)

func main(){
	greetings := "\t Hello, World "
	fmt.Printf("%d %s\n", len(greetings), greetings)

	trimmed := strings.TrimSpace(greetings)
	fmt.Printf("%d %s\n", len(trimmed), trimmed)
}
