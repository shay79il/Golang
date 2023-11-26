package main

import (
	"fmt"
)

func main()  {
	
	// Declaring Pointers and dereference Operator
	x := 1
	// Option 1
	var x_ptr *int = nil
	x_ptr  = &x;
	fmt.Printf("x_ptr is %T and it's %v \n", x_ptr, x_ptr)
	fmt.Printf("*x_ptr is %T and it's %v \n", *x_ptr, *x_ptr)
	fmt.Printf("x is %T and it's %v \n", x, x)
	*x_ptr = 999
	fmt.Printf("x is %T and it's %v \n", x, x)
	fmt.Printf("x is %T and it's %v \n", *x_ptr, *x_ptr)
	fmt.Println()
	
	// Option 2
	var x_ptr2 = &x
	fmt.Printf("x_ptr2 is %T and it's %v \n", x_ptr2, x_ptr2)
	fmt.Printf("x is %T and it's %v \n", *x_ptr2, *x_ptr2)
	fmt.Println()
	
	// Option 3
	x_ptr3 := &x
	fmt.Printf("x_ptr3 is %T and it's %v \n", x_ptr3, x_ptr3)
	fmt.Printf("x is %T and it's %v \n", *x_ptr3, *x_ptr3)
	fmt.Println()
	
	
	// Passing by Reference in Functions 
	var str string = "Hello"
	fmt.Println(str)
	modify_1(&str)
	fmt.Println(str)
	fmt.Println()
	
	// Slices are passed by Reference by default
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modify_2(slice)
	fmt.Println(slice)
	fmt.Println()
	
	// Maps are passed by Reference by default
	map_1 := make(map[string]int)
	map_1["A"] = 1
	map_1["B"] = 2
	fmt.Println(map_1)
	modify_3(map_1)
	fmt.Println(map_1)
	fmt.Println()
	

}

func modify_1(str *string) {
	*str = "World"
}

func modify_2(arr []int) {
	arr[0] = 100
}

func modify_3(m map[string]int) {
	m["K"] = 100
}

