/*
We are building a software for a store that sells three products: apples, bananas, and oranges. 
We need to write a function that takes the name of a product and its price as arguments 
and returns the price of the product with a discount applied. 
The discount should be 10% for apples and 20% for bananas. 
Oranges do not have a discount.
*/

package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	var discount float64
	switch product{
	case "apples":
		discount = 0.1
	case "bananas":
		discount = 0.2
	default:
		discount = 0
	}
	return price * (1 - discount)
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
