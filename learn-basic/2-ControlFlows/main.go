package main

import (
	"fmt"
)

	func main()  {
	
	// ###########################
	// ### Control flow
	// ###########################
	
	fmt.Println()
	// if statement
	fmt.Println("// if statement")
	var num1 int = 5
	var num2 int = 5
	if (num1 < num2) {
		fmt.Printf("num1 <num2\n")
		} else if (num1 == num2) {
			fmt.Printf("num1 = num2\n")
			}	else {
				fmt.Printf("num1 > num2\n")
			}
			
	fmt.Println()
	
	// Switch statement on a variable
	fmt.Println("// Switch statement on a variable")
	switch num1 {
	case 5:
		fmt.Println("num1 = ", num1)
		fallthrough
	default:
		fmt.Printf("Default\n")
		
	}	
	fmt.Println()
	
	
	// Switch statement with conditions
	fmt.Println("// Switch statement with conditions")
	switch {
	case (num1 < num2), (num1 == num2):
		fmt.Printf("num1 <= num2\n")
		fallthrough
	default:
		fmt.Printf("Default\n")
	}
	fmt.Println()
	
	// For loop statement
	fmt.Println("// For loop statement")
	for i := 1; i <= 3 ;i++ {
		fmt.Printf("Hello world!!!\n")
	}
	fmt.Println()
	
	// while loop statement
	fmt.Println("// while loop statement")
	var i = 1
	for i <= 3 {
		fmt.Printf("Hello world!!!\n")
		i++
	}
	fmt.Println()
	
	// Infinite loop
	fmt.Println("// Infinite loop")
	// for {
		// 	fmt.Printf("infinite loop\n")
		// }
		// fmt.Println()
		
		// Break & continue statement
		fmt.Println("// Break & continue statement")
		for i := 1; i <= 10 ;i++ {
			if (i == 7){

				fmt.Printf("Break at i = %d\n", i)
				break

			} else if (i % 2 == 0) {

				fmt.Printf("i is even so skip, i = %d\n", i)
				continue
			}
			fmt.Printf("i = %d\n", i)

		}
		fmt.Println()
		
		
	}
	
	
	
	