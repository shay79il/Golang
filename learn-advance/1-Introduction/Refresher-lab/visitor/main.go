package main

import "fmt"

// Declare variable activeUserCount
var activeUserCount int

func entry() {
	activeUserCount++
}

func exit() {
	activeUserCount--
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}
