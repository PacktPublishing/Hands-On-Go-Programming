package main

import "os"

func main() {
	os.Rename("old.txt", "new.txt")
}

