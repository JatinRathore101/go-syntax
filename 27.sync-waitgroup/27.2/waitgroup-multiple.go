package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup with multiple goroutines
func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, l := range "ABCDE" {
		fmt.Println("Letter:", string(l))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Starting goroutines with WaitGroup...")

	wg.Add(2) // 2 goroutines to wait for
	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait() // waits until both call wg.Done()
	fmt.Println("All goroutines finished safely.")
}
