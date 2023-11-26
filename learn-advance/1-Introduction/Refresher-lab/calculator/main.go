package main

import "fmt"

func calculate(a int, b int) (results []float64) {
  	return append(results, float64(a+b), float64(a-b), float64(a*b), float64(a/b))
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}