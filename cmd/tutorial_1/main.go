package main

import (
	"fmt"
	// "math/rand"
	"time"
	"sync"
)

var mt sync.Mutex
var wg sync.WaitGroup

var dbData = []string{"id1", "id2", "id3", "id4", "id5"};

var results []string

func main(){
	t0 := time.Now();

	for i:=0;i<5;i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total execution time is %v \n", time.Since(t0))
	fmt.Printf("The result is %v \n", results)
}

func dbCall(index int){
	var delay = 2000;
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println(dbData[index])
	mt.Lock()
	results = append(results, dbData[index])
	mt.Unlock()
	wg.Done()
}