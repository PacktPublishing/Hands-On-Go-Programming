package main

import (
	"net/http"
	"os"
	"io"
	"fmt"
)

func main(){
	imageUrl := "https://golang.org/doc/gopher/doc.png"
	response, err := http.Get(imageUrl)
	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	file, err2 := os.Create("gopher.png")

	if err2 != nil{
		panic(err2)
	}

	_, err3 := io.Copy(file, response.Body)
	if err3 != nil{
		panic(err3)
	}

	file.Close()
	fmt.Println("Image downloading is successful.")
}
