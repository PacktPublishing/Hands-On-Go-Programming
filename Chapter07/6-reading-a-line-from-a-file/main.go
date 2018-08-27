package main

import (
	"os"
	"bufio"
	"fmt"
)

func main(){
	fmt.Println(ReadLine(3))
}

func ReadLine(lineNumber int) string{
	file, _ := os.Open("names.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan(){
		if lineCount == lineNumber{
			return fileScanner.Text()
		}
		lineCount++
	}
	defer file.Close()
	return ""
}
