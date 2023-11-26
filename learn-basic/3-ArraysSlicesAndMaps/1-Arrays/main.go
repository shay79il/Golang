package main

import (
	"fmt"
)

func main()  {
	
	// Arrays declaration options
	fmt.Println("// Arrays declaration options")
	
	// Option1 - default init is all values are "0"
	var grades_1 [5]int
	fmt.Println("grades_1 = ", grades_1)
	fmt.Println()
	
	
	// Option2 - explicit init
	var grades_2 [5]int = [5]int {10, 20, 30, 40, 50}
	fmt.Println("grades_2 = ", grades_2)
	fmt.Println()
	
	// Option3 - Short declaration
	grades_3 := [5]int {10, 20, 30, 40, 50}
	fmt.Println("grades_3 = ", grades_3)
	fmt.Println()
	
	// Option4 - ... indicates unknown number of array elements on init
	grades_4 := [...]int {10, 20, 30, 40, 50}
	fmt.Println("grades_4 = ", grades_4)
	fmt.Println()
	
	
	var fruits [3]string = [3]string {"Apple", "Orange", "Banana"}
	fmt.Println("Fruits array")
	fmt.Println("len(fruits)=  ", len(fruits))
	fmt.Println()
	
	// Iterating an array
	fmt.Println("// Iterating an array")
	
	// Option 1
	fmt.Println("// Option 1")
	for i := 0 ; i < len(fruits) ; i++ {
		fmt.Println(fruits[i])
	}
	
	fmt.Println()
	
	// Option 2
	fmt.Println("// Option 2")
	for index, element := range fruits {
		fmt.Println(index ,"=>", element)
	}
	
	fmt.Println()
	
	// 2 Dimensional array
	fmt.Println("// 2 Dimensional array")
	arr1 := [3][2] int{{2, 4}, {4, 16}, {8, 64}}
	for index, element := range arr1 {
		fmt.Println(index ,"=>", element)
	}
	
	fmt.Println()
	
	
	
}



