package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	// t0 := time.Now()
	// for i := range dbData {
	// 	dbCall(i)
	// }
	// fmt.Println("Total execution time:", time.Since(t0))

	t1 := time.Now()
	for i := range dbData {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("Total execution time:", time.Since(t1))
	fmt.Println("The results are", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the DB is", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}