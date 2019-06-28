package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CallClear will clear screen on Windows, mac or unix systems
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! Can't clear terminal screen")
	}
}

// PrintSystemInfo prints general info about the system that is running the current program
func PrintSystemInfo() {
	fmt.Println("System Information:")
	fmt.Println("GOOS: ", runtime.GOOS)
	fmt.Println("GOARCHS: ", runtime.GOARCH)
	// NumCPU returns the number of logical CPUs usable by the current process
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("\nMemory usage")
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
