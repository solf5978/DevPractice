package main

import (
	"fmt"
	"os/exec"
)

func main() {
	prc := exec.Command("dir", "-a")
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		//fmt.Println(out.String())
	}
}
