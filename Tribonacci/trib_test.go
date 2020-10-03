package main

import "testing"

// from fib_test.go
const Num = 40
func BenchmarkTribonacciUsingRecursion(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		tribonacciUsingRecursion(Num)
	}
}

func BenchmarkTribonacciUsingOptimizedRecursion(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		tribonacciUsingOptimizedRecursion(Num)
	}
}

func BenchmarkTribonacciUsingArray10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		tribonacciUsingArray(Num)
	}
}