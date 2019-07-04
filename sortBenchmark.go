package main

import (
	"sort"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, n)
	}
	return s
}

func sortBenchmark(sizeOfSlice int) {
	s := generateSlice(sizeOfSlice)
	sort.Ints(s)
	waitGroup.Done()
}

// RunSortBenchmark runs the sort algorithm from Go package "sort"
// It receives the amount of threads that will execute the sort at the same time and the size of the slice to be sorted
func RunSortBenchmark(amountOfThreads int, sizeOfSlice int) {

	waitGroup.Add(amountOfThreads)

	for i := 0; i < amountOfThreads; i++ {
		go sortBenchmark(sizeOfSlice)
	}

	waitGroup.Wait()

}
