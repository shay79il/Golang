package main

import (
	"fmt"
	"runtime"
	"time"
)


func main() {
	
	// Anonymous Go-routine behaves same as regular Anonymous function
	go func(){
		fmt.Println("In anonymous func")
	}()

	fmt.Println("runtime.NumCPU() = ", runtime.NumCPU())
	time.Sleep(1* time.Second)
}
