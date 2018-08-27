package main

import (
	"os"
	"fmt"
)

func main(){
	if _, err := os.Stat("log.txt"); os.IsNotExist(err){
		fmt.Println("Log.txt file does not exist")
	}else{
		fmt.Println("Log.txt file exists")
	}
}
