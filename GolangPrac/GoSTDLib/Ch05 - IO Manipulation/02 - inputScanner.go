package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//	Scanner is able to scan the input lines by lines
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}
}
