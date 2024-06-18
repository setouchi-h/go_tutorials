package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1000000
	testSlice := []int{}
	testSlice2 := make([]int, 0, n)

	fmt.Println("Total time without preallocation:", timeLoop(testSlice, n))
	fmt.Println("Total time with preallocation:", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}