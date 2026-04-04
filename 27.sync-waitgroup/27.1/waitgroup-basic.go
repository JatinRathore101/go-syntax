package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup: wait for goroutines to finish without arbitrary sleeps
func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // signal completion when this function returns
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)          // we are launching 1 goroutine
	go printNumbers(&wg) // pass pointer to WaitGroup

	wg.Wait() // block until all goroutines call wg.Done()
	fmt.Println("Goroutine finished safely.")
}
