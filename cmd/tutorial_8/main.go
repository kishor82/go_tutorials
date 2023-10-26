package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var rm = sync.RWMutex{}
var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v \n", time.Since(t0))
	fmt.Printf("\n The results are %v", results)
	t1 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v \n", time.Since(t1))
}

func dbCall(i int) {
	// Simulate DB call delay

	var delay float32 = 2000
	time.Sleep(time.Duration((delay) * float32(time.Millisecond)))
	fmt.Println("The result from the database is:", dbData[i])
	// m.Lock()
	// results = append(results, dbData[i])
	// m.Unlock()
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	rm.Lock()
	results = append(results, result)
	rm.Unlock()
}

func log() {
	rm.RLock()
	fmt.Printf("\n The current results are: %v", results)
	rm.RUnlock()
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}

	wg.Done()
}
