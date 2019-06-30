// Created by: Tiago Cabral
// This software was develop for the discipline of Experimental Computing to verify the performance of a CPU
// It consists of a heavy load of algorithms being executed in many threads at once.

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	ClearScreen()

	PrintSystemInfo()
	PrintMemUsage()

	fmt.Println("\n\nPress enter to begin tests >")
	buf := bufio.NewReader(os.Stdin)
	buf.ReadBytes('\n')

	start := time.Now()

	RunFpBenchmark(100)

	PrintMemUsage()

	timeOfExecution := time.Since(start)

	fmt.Println("Benchmark done in: ", timeOfExecution)
}
