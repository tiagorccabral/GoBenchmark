package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// The benchmark function must run the target code b.N times.
// During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably.
func benchmarkFP(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		RunFpBenchmark(i)
	}
}

func BenchmarkFpBenchmark1(b *testing.B) { benchmarkFP(1, b) }
func BenchmarkFpBenchmark2(b *testing.B) { benchmarkFP(10, b) }
func BenchmarkFpBenchmark3(b *testing.B) { benchmarkFP(100, b) }
func BenchmarkFpBenchmark4(b *testing.B) { benchmarkFP(1000, b) }

func TestMain(m *testing.M) {
	ClearScreen()

	PrintSystemInfo()
	PrintMemUsage()

	start := time.Now()

	code := m.Run()

	PrintMemUsage()

	timeOfExecution := time.Since(start)

	fmt.Println("Benchmark done in: ", timeOfExecution)

	os.Exit(code)
}
