package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// t0 := time.Now()
	// for i := range dbData {
	// 	dbCall(i)
	// }
	// fmt.Println("Total execution time:", time.Since(t0))

	t1 := time.Now()
	for i:=0; i<1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Println("Total execution time:", time.Since(t1))
}

func count() {
	var res int
	for i:=0; i<100000000; i++ {
		res += i
	}
	wg.Done()
}