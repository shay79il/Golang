package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(i int, wg *sync.WaitGroup) {
	fmt.Println(i*i)

	// Notify end of go routine
	wg.Done()
}

func main() {

	// Declare WaitGroup
	var wg sync.WaitGroup

	start := time.Now()

	// Decide how many go routines will run Concurrently
	wg.Add(100)


	for i := 1; i <= 100; i++ {
		go calculateSquare(i, &wg)
	}
	elapsed := time.Since(start)

	// Wait for ALL concurrent go routines
	wg.Wait()
	
	fmt.Println("Function took ", elapsed)

}