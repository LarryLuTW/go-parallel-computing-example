package main

import (
	"math/rand"
	"runtime"
)

func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SumTwoParallel(numbers []int) int {
	mid := len(numbers) / 2

	ch := make(chan int)
	go func() { ch <- Sum(numbers[:mid]) }()
	go func() { ch <- Sum(numbers[mid:]) }()

	total := <-ch + <-ch
	return total
}

func SumMaxParallel(numbers []int) int {
	nCPU := runtime.NumCPU()
	nNum := len(numbers)

	ch := make(chan int)
	for i := 0; i < nCPU; i++ {
		from := i * nNum / nCPU
		to := (i + 1) * nNum / nCPU
		go func() { ch <- Sum(numbers[from:to]) }()
	}

	total := 0
	for i := 0; i < nCPU; i++ {
		total += <-ch
	}
	return total
}

func generateRandomArray() []int {
	n := 1000000
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	return arr
}

func main() {
	arr := generateRandomArray()
	// _ = trace.Start(os.Stdout)
	// defer trace.Stop()
	SumMaxParallel(arr)
}
