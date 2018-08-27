package main

import (
	"os/exec"
	"fmt"
)

func main() {
	lsCommand := exec.Command("ls")
	output,_ := lsCommand.Output()
	lsCommand.Run()

	fmt.Println(lsCommand.Process.Pid)
	fmt.Println(string(output))
}
