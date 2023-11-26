package main

import (
	"fmt"
	"sync"
	"time"
)


// ##################################
// Example 
// ##################################

// Sends data to the channel
func sell(ch chan string, wg *sync.WaitGroup) {
	defer close(ch)
	ch <- "Chair"
	ch <- "Desk"
	ch <- "Map"
	fmt.Println("Sent all data to the channel")
	wg.Done()

}

// Receives data from the channel
func buy(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data from the channel")
	for message := range ch{
		
		fmt.Println("Received:", message)
		time.Sleep(time.Second) // Simulating some processing time
	}
	// for {
	// 	// Receiving messages from the channel
	// 	message, isOpen := <-ch
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println("Received:", message)
	// 	time.Sleep(time.Second) // Simulating some processing time
	// }
	wg.Done()
}


func main() {

	// Declaring sync mechanism
	var wg sync.WaitGroup
	
	// Calculating the time to process
	start := time.Now()

	
	// Declaring channels
	// we need to specify the channel datatype 
	// Available channel operations:
	// - Sending a value 		(E.g. ch <- v)
	// - Receiving a value 	(E.g. ch := <- v)
	// - Closing a channel	(E.g. close(ch) )
	// - Querying Buffer of a channel (E.g. cap(ch) )
	// - Querying length of a channel (E.g. len(ch) )
	
	// Option 1
	var ch_1 chan string
	fmt.Printf("ch_1 = %v and of type %T\n", ch_1, ch_1)
	
	// Option 2 
	ch_2 := make(chan string, 3)
	

	wg.Add(2)
	go sell(ch_2, &wg)
	go buy(ch_2, &wg)
	
	
	// Wait for the go routines to end
	wg.Wait()

	fmt.Println("Function took ", time.Since(start))
}