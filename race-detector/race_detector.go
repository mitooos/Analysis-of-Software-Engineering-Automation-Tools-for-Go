package main

import (
	"fmt"
	"sync"
)

var raceCondition int

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go sum10(&wg)
	}
	wg.Wait()
	fmt.Println(raceCondition)
}

func sum10(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		value := raceCondition
		value++
		raceCondition = value
	}
	wg.Done()
}
