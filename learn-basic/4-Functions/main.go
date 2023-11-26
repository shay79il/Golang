package main

import (
	"fmt"
	"strings"
)

func main()  {
	
	// Regular function
	// ################
	
	// No return value
	printGreeting("Shay")
	fmt.Println()
	
	// Function which gets 2 args and return 1 
	fmt.Println("// Function which gets 2 args and return 1 ")
	fmt.Println("addNumbers(1 ,2) = ", addNumbers(1 ,2))
	fmt.Println()
	
	// Function which gets 2 args and return 2 
	sum, diff := operation(1 ,2)
	fmt.Println("operation(1 ,2) = ",sum, ",", diff)
	fmt.Println()
	
	// Function which gets 2 args and return 2 - but now look at the func signature
	sum, diff = operation2(1 ,2)
	fmt.Println("operation2(1 ,2) = ",sum, ",", diff)
	fmt.Println()
	
	
	// Variadic function
	// ##################
	// Gets arbitrary number of args
	fmt.Println("sumNumbers = ",sumNumbers(1, 2, 3, 4))
	fmt.Println()
	
	// Gets first mandatory string and another arbitrary number of args 
	printDetails("Shay", "Math", "English", "Physics")
	fmt.Println()
	
	// Use of blank identifier
	fmt.Println("// Use of blank identifier")
	fmt.Println("sum, _ = operation2(1 ,2)")
	sum, _ = operation2(1 ,2)
	fmt.Println("sum = ",sum)
	fmt.Println()
	
	// Anonymous Functions
	// #####################
	fmt.Println("// Anonymous Functions")
	
	// Option 1
	x_1 := func(s string) string {
		return strings.ToUpper(s)
	}
	
	fmt.Printf("x_1 type is %T\n", x_1)
	fmt.Println("x_1(\"Hi\") = ", x_1("Hi"))
	fmt.Println()
	
	// Option 2
	x_2 := func(s string) string {
		return strings.ToUpper(s)
		}("Hi there")
		
		fmt.Printf("x_2 type is %T\n", x_2)
		fmt.Println("x_2 = ", x_2)
		fmt.Println()
		
		// Option 3
		var x_3 = func(a int, b int) int {
			return a + b
		}
		
		fmt.Printf("x_3 type is %T\n", x_3)
		fmt.Println("x_3(5, 7) = ", x_3(5, 7))
		fmt.Println()
		
		// Option 4
		var (
			x_4 = func(s string) { 
			fmt.Println("Hey there,", s) 
		}
	)
	
	fmt.Printf("x_4 type is %T\n", x_4)
	x_4("Yo, yo")
	fmt.Println()
	
	// High order functions 
	// A function that gets or returns a function as an arg
	// ######################################################
	fmt.Println("// High order functions")
	
	
	// 1 - area, 2 - perimeter, 3 = diameter:
	var query int = 1
	var radius float64 = 5.5
	printResult(radius, getFunction(query))

	// Defer Statement
	
	defer printName("shay")
	printRollNo(777)
	printAddress("Dgania")



}






// No return value
func printGreeting(str string) {
	fmt.Println("Hello", str)
}

// Return 1 value
func addNumbers(a int, b int) int {
	sum := a + b 
	return sum
}

// Return 2 values
func operation(a int, b int) (int, int) {
	sum := a + b 
	diff := a - b
	return sum, diff
}

// Return 2 DECLARED values
func operation2(a int, b int) (sum int, diff int) {
	sum = a + b 
	diff = a - b
	return 
}

// Variadic function
// numbers here is a slice
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum

}

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey", student, ", here are your subjects - ")
	for i, sub := range subjects {
		if (i == len(subjects) - 1) {
			fmt.Printf("%s ", sub)
			break
		}
		fmt.Printf("%s, ", sub)
	}
	fmt.Println()
}




// High order functions
// #####################
func calcArea(r float64) float64 {
	return 3.14 * r * r	
}

func calcPerimeter(r float64) float64 {
	return 2* 3.14 * r	
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ",result)
	fmt.Println("Thanks you!")
}

func getFunction(query int) func(r float64) float64 {
	query_to_function := map[int] func(r float64) float64 {
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_function[query]
}


// Using Defer Statement
func printName(str string) {
	fmt.Println(str)
}

func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(adr string) {
	fmt.Println(adr)
}

