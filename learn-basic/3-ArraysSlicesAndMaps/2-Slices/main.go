package main

import (
	"fmt"
)

func main()  {
	
	// Slices declaration options
	// #########################
	
	fmt.Println("// Slices")
	fmt.Println("// #########")
	fmt.Println()
	
	fmt.Println("// Slices declaration options")
	
	// Option 1
	slice_1 := [] int{1, 2, 3}
	fmt.Println("slice_1 = ", slice_1)
	fmt.Println("len(slice_1) = ", len(slice_1))
	fmt.Println("cap(slice_1) = ", cap(slice_1))
	fmt.Println()
	
	
	// Option 2
	grades_1 := [10]int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_2 := grades_1[1:8]
	fmt.Println("slice_2 = ", slice_2)
	fmt.Println("len(slice_2) = ", len(slice_2))
	fmt.Println("cap(slice_2) = ", cap(slice_2))
	fmt.Println()
	
	// Option 3
	slice_3 := make([]int, 5, 10)
	fmt.Println("slice_3 = ", slice_3)
	fmt.Println("len(slice_3) = ", len(slice_3))
	fmt.Println("cap(slice_3) = ", cap(slice_3))
	fmt.Println()
	
	
	// Modify element in the slice will effect the underling array
	fmt.Println("// Modify element in the slice")
	slice_2[0] = 7777
	fmt.Println("slice_2[0] = 7777")
	fmt.Println("grades_1 = ", grades_1)
	fmt.Println("slice_2 = ", slice_2)
	fmt.Println()
	
	// Slice of a slice
	fmt.Println("// Slice of a slice")
	slice_4 := slice_2[0:3]
	fmt.Println("slice_4 := slice_2[0:3]")
	fmt.Println("slice_4 = ", slice_4)
	fmt.Println("len(slice_4) = ", len(slice_4))
	fmt.Println("cap(slice_4) = ", cap(slice_4))
	fmt.Println()
	
	
	// Appending slices
	// ################
	fmt.Println("// Appending slices")
	fmt.Println()

	// Option 1
	grades_2 := [4]int {10, 20, 30, 40}
	slice_5 := grades_2[1:3]
	fmt.Println("grades_2 = ", grades_2)
	fmt.Println("slice_5 := grades_2[1:3]")
	fmt.Println("slice_5 = ", slice_5)
	fmt.Println("len(slice_5) = ", len(slice_5))
	fmt.Println("cap(slice_5) = ", cap(slice_5))
	fmt.Println()

	slice_5 = append(slice_5, 900)
	fmt.Println("slice_5 = append(slice_5, 900)")
	fmt.Println("slice_5 = ", slice_5)
	fmt.Println("len(slice_5) = ", len(slice_5))
	fmt.Println("cap(slice_5) = ", cap(slice_5))
	fmt.Println()
	
	slice_5 = append(slice_5, -9, 500)
	fmt.Println("slice_5 = append(slice_5, -9, 500)")
	fmt.Println("slice_5 = ", slice_5)
	fmt.Println("len(slice_5) = ", len(slice_5))
	fmt.Println("cap(slice_5) = ", cap(slice_5))
	fmt.Println()
	

	sub_slice_1 := append(slice_5, 555)
	fmt.Println("sub_slice_1 := append(slice_5, 555)")
	fmt.Println("sub_slice_1 = ", sub_slice_1)
	fmt.Println("len(sub_slice_1) = ", len(sub_slice_1))
	fmt.Println("cap(sub_slice_1) = ", cap(sub_slice_1))
	fmt.Println()
	
	sub_slice_2 := append(slice_5, -7, 444)
	fmt.Println("sub_slice_2 := append(slice_5, -7, 444)")
	fmt.Println("sub_slice_2 = ", sub_slice_2)
	fmt.Println("len(sub_slice_2) = ", len(sub_slice_2))
	fmt.Println("cap(sub_slice_2) = ", cap(sub_slice_2))
	fmt.Println()
	
	// Appending 2 slices with  "..." operator
	fmt.Println("// Appending 2 slices with  \"...\" operator")
	fmt.Println("sub_slice_3 := append(sub_slice_1, sub_slice_2...)")
	sub_slice_3 := append(sub_slice_1, sub_slice_2...)
	fmt.Println("sub_slice_3 = ", sub_slice_3)
	fmt.Println("len(sub_slice_3) = ", len(sub_slice_3))
	fmt.Println("cap(sub_slice_3) = ", cap(sub_slice_3))
	fmt.Println()
	
	// Option 2 - "removing" element grades_3[2]
	grades_3 := [5]int {10, 20, 30, 40, 50}
	fmt.Println()
	fmt.Println("\"Removing\" element grades_3[2]")
	fmt.Println("###############################")
	
	fmt.Println("grades_3 = ", grades_3)
	i2 := 2
	new_slice := append(grades_3[:i2], grades_3[i2+1:]...)
	fmt.Println("new_slice = ", new_slice)
	fmt.Println("len(new_slice) = ", len(new_slice))
	fmt.Println("cap(new_slice) = ", cap(new_slice))
	fmt.Println()
	
	
	// Coping slices
	// ################
	src_slice := []int {10, 20, 30, 40, 50}
	dst_slice := make([]int, 3)
	
	fmt.Println()
	fmt.Println("Coping slices")
	fmt.Println("#############")
	fmt.Println("src_slice = ", src_slice)
	fmt.Println("dst_slice = ", dst_slice)
	fmt.Println()
	
	fmt.Println("After coping")
	num := copy(dst_slice, src_slice)
	fmt.Println("src_slice = ", src_slice)
	fmt.Println("dst_slice = ", dst_slice)
	fmt.Println("len(dst_slice) = ", len(dst_slice))
	fmt.Println("cap(dst_slice) = ", cap(dst_slice))
	fmt.Println()
	fmt.Println("num of copied elements = ", num)
	fmt.Println()
	
	// Iterating a slice
	fmt.Println("// Iterating an slice")
	fmt.Println("sub_slice_3 = ", sub_slice_3)
	fmt.Println()
	
	// Option 1
	fmt.Println("// Option 1")
	for i := 0 ; i < len(sub_slice_3) ; i++ {
		fmt.Printf("%v ", sub_slice_3[i])
	}
	
	fmt.Println()
	fmt.Println()

	// Option 2
	fmt.Println("// Option 2")
	for _, element := range sub_slice_3 {
		fmt.Printf("%v \t ", element)
	}
	
	fmt.Println()
	fmt.Println()
	
	// Option 3
	fmt.Println("// Option 3")
	for index, element := range sub_slice_3 {
		if (index % 2 == 1) {
			fmt.Println()
		}
		fmt.Printf("%v => %v \t ", index, element)
	}
	
	fmt.Println()
}



