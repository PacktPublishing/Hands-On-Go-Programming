package main

import (
	"os"
	"io"
)

func main(){
	original, err := os.Open("original.txt")

	if err != nil{
		panic(err)
	}



	original_copy, err2 := os.Create("target/original.txt")

	if err2 != nil{
		panic(err2)
	}

	defer original_copy.Close()

	_, err3 := io.Copy(original_copy, original)

	if err3 != nil{
		panic(err3)
	}

	original.Close()
	os.Remove("original.txt")
}
