package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func sumTwoParallel(numbers []int) int {
	mid := len(numbers) / 2

	ch := make(chan int)
	go func() { ch <- sum(numbers[:mid]) }()
	go func() { ch <- sum(numbers[mid:]) }()

	total := <-ch + <-ch
	return total
}

func sumMaxParallel(numbers []int) int {
	nCPU := runtime.NumCPU()
	nNum := len(numbers)

	ch := make(chan int)
	for i := 0; i < nCPU; i++ {
		from := i * nNum / nCPU
		to := (i + 1) * nNum / nCPU
		go func() { ch <- sum(numbers[from:to]) }()
	}

	total := 0
	for i := 0; i < nCPU; i++ {
		total += <-ch
	}
	return total
}

func main() {
	n := 1000000
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	times := 1000

	startTime := time.Now()
	for i := 0; i < times; i++ {
		sum(arr)
	}
	fmt.Println(time.Since(startTime))

	startTime = time.Now()
	for i := 0; i < times; i++ {
		sumTwoParallel(arr)
	}
	fmt.Println(time.Since(startTime))

	startTime = time.Now()
	for i := 0; i < times; i++ {
		sumMaxParallel(arr)
	}
	fmt.Println(time.Since(startTime))
}
