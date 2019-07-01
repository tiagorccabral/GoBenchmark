package main

import (
	"sync"
)

var waitGroup sync.WaitGroup

const sizeOfSlice = 1000000
const pi float64 = 3.141592653589793238462643
const eulerNum float64 = 2.718281828459045235360287
const bigFloat float64 = 10987654321.123456789

func fpBenchmark(i int) {

	multiples := make([]float64, sizeOfSlice)
	results := make([]float64, sizeOfSlice)

	// initializes slice with a random floating point number that is the result of a floating point operation
	for i := 0; i < sizeOfSlice; i++ {
		multiples[i] = pi * (bigFloat * pi) * (eulerNum * eulerNum)
	}

	// multiplies number by big floating point numbers
	for i := 0; i < sizeOfSlice; i++ {
		results[i] = multiples[i] * bigFloat * pi * eulerNum
	}

	// division of floating point numbers
	for i := 0; i < sizeOfSlice; i++ {
		results[i] = (results[i]) / ((pi * bigFloat) * (eulerNum / ((bigFloat * pi * pi) / eulerNum)))
	}

	waitGroup.Done()
}

// RunFpBenchmark runs the Floating Point algorithms
func RunFpBenchmark(amountOfThreads int) {

	// fmt.Println("Floating point tests starting...")

	// start := time.Now()

	waitGroup.Add(amountOfThreads)

	for i := 0; i < amountOfThreads; i++ {
		go fpBenchmark(i)
	}

	waitGroup.Wait()

	// timeOfExecution := time.Since(start)
	// fmt.Println("Floating point benchmarks done in: ", timeOfExecution)
}
