// Created by: Tiago Cabral
// This software was develop in the discipline of Experimental Computing to verify the performance of a CPU
// It consists of a heavy load of algorithms being executed in many threads at once.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	CallClear()

	PrintSystemInfo()
	
	fmt.Println("\n\nPress enter to begin tests >")
	buf := bufio.NewReader(os.Stdin)
	buf.ReadBytes('\n')
}