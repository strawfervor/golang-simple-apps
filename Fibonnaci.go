// This piece of code will calculate nth element of Fibonacci sequence, it takes very long to calculate numbers bigger than 40
package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println("nth element of Fibonacci sequence to calculate\n(it takes very long to calculate numbers bigger than 40): ")
	var nth int
	fmt.Scanf("%d", &nth)
	var result int
	result = fibonacci(nth)
	fmt.Printf("%d element of Fibonacci sequence is: %d", nth, result)
}
