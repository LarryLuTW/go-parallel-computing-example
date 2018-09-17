package main

import (
	"testing"
)

func BenchmarkSum(b *testing.B) {
	arr := generateRandomArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(arr)
	}
}

func BenchmarkSumTwoParallel(b *testing.B) {
	arr := generateRandomArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumTwoParallel(arr)
	}
}

func BenchmarkSumMaxParallel(b *testing.B) {
	arr := generateRandomArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumMaxParallel(arr)
	}
}
