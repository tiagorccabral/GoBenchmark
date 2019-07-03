package main

func stringConcatBenchmark(lengthOfString int) {

	var str string

	for n := 0; n < lengthOfString; n++ {
		str += "x"
	}

	waitGroup.Done()
}

// RunStringConcat runs a sequence of operations with + operator to concat a string
// It receives the amount of threads that will execute the sort at the same time
// and the length of the string to be concateneted
func RunStringConcat(amountOfThreads int, lengthOfString int) {

	waitGroup.Add(amountOfThreads)

	for i := 0; i < amountOfThreads; i++ {
		go stringConcatBenchmark(lengthOfString)
	}

	waitGroup.Wait()

}
