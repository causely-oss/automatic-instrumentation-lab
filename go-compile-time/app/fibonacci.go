package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	defer trace_fibonacci()()
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(10))
}

func trace_fibonacci() func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		fmt.Printf("Function fibonacci took: %v\n", duration)
	}
}
