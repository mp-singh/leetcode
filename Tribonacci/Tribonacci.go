package main

import (
	"log"
)

func tribonacciUsingRecursion(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	return tribonacciUsingRecursion(n-3) + tribonacciUsingRecursion(n-2) + tribonacciUsingRecursion(n-1)
}

func tribonacciUsingOptimizedRecursion(n int) int {
	state := map[int]int{}
	return optimizedTrib(n, state)

}

func optimizedTrib(n int, state map[int]int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}

	if val, ok := state[n]; ok {
		return val
	}

	tribResult := optimizedTrib(n-3, state) + optimizedTrib(n-2, state) + optimizedTrib(n-1, state)
	state[n] = tribResult
	return tribResult
}

func tribonacciUsingArray(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}
	//fmt.Println(dp)
	return dp[n]
}

func main() {
	log.Printf("tribonacciUsingRecursion: %d", tribonacciUsingRecursion(10))
	log.Printf("tribonacciUsingOptimizedRecursion: %d", tribonacciUsingOptimizedRecursion(10))
	log.Printf("tribonacciUsingArray: %d", tribonacciUsingArray(10))
}
