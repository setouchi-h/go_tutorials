package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
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
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("The current results are", results)
	m.RUnlock()
}