package main

import (
	"fmt"
	"unsafe"
	"reflect"
	"strconv"
)

// Global variable
var globalVar string = "I am a global var"

func main()  {

	// Printf-format specifiers 
	// ----   -----------
	// Verb 	Description 
	// ----   -----------
	// %v 		default format 
	// %T 		type of the value 
	// %d 		integers 
	// %c 		character 
	// %q 		quoted characters/string 
	// %s 		plain string 
	// %t 		true or false 
	// %f 		floating numbers 
	// %.2f 	floating numbers upto2 decimal places

	fmt.Println(globalVar)

	// ------------------------
	var str1 string = "Hello World" 
	str1 += "!!!"
	fmt.Println(str1)
	// ------------------------
	
	// Declaration of 2 vars of the same data type
	var str2, str3 string = "Hello","World!!!!!!" 
	fmt.Println(str2, str3)
	
	// ------------------------
	// Declaration of 2 vars of different data types
	var (
		str4 string = "I am a string"
		intVal int = 5
	)
	fmt.Println(str4, intVal)
	
	// ------------------------
	var myUint uint
	sizeOfUint := unsafe.Sizeof(myUint)
	fmt.Printf("Size of uint: %d bytes\n", sizeOfUint)
	fmt.Print("Size of uint: ", sizeOfUint, " bytes\n")
	
	
	// ------------------------
	var i uint = 8
	i += 1
	fmt.Println(i)
	
	// ------------------------
	var str5 string = "Hello World!"
	fmt.Println(str5)

	
	// ------------------------
	str6 := "short string declaration"
	intVal2 := 19
	fmt.Println(str6, intVal2)
	
	// ------------------------
	// Zero values 
	// Uninitialized variables get initial value as follows
	
	var isOK bool // isOK = 0
	var intVal3 int // intVal3 = 0
	var floatVal float64 // floatVal = 0
	var str7 string // str7 = ""
	fmt.Printf("isOK = %t, intVal3 = %d, floatVal = %.2f, str7 = %q\n\n",isOK,intVal3,floatVal, str7)
	
	
	// ------------------------
	// User input - option for multiple inputs
	var name string = "Shay"
	var age int = 44
	// fmt.Print("Please enter your name and age: ") 
	// num, err := fmt.Scanf("%s %d", &name, &age)
	// fmt.Println("Hey there,", name, "your age is", age) 
	// fmt.Println("Count,", num) 
	// fmt.Println("Errors,", err) 
	
	// ------------------------
	// User input TYPE
	fmt.Printf("variable name = %v is of type %T\n",name, name)
	fmt.Printf("variable age = %v is of type %T\n",age, age)
	fmt.Printf("variable isOK = %v is of type %T\n",isOK, reflect.TypeOf(isOK))
	
	// ------------------------
	// Casting
	var floatNum float64 = float64(age)
	fmt.Printf("floatNum = %.2f \n",floatNum)
	
	var str8 string = strconv.Itoa(age)
	fmt.Printf("str8 = %q \n",str8)
	
	var str9 string = "200"
	intVal4, err := strconv.Atoi(str9)
	fmt.Printf("intVal4 = %v, %T\n",intVal4, intVal4)
	fmt.Printf("err = %v, %T\n",err, err)
	
	var str10 string = "200abc"
	intVal5, err := strconv.Atoi(str10)
	fmt.Printf("intVal4 = %v, %T\n",intVal5, intVal5)
	fmt.Printf("err = %v, %T\n",err, err)

	// ------------------------
	// Constants
	const constVar int = 64
	fmt.Printf("const constVar = %v is of type %v\n",constVar, reflect.TypeOf(constVar))
	
	// UNtyped constant
	const name1 = "Harry Potter"
	const isMuggle = false
	const age3 = 21
	fmt.Printf("variable name = %v is of type %T\n",name1, name1)
	fmt.Printf("variable isMuggle = %v is of type %T\n",isMuggle, isMuggle)
	fmt.Printf("variable age3 = %v is of type %T\n",age3, age3)
	
	// ###########################
	// ### Control flow
	// ###########################
	
	// if statement
	var num1 int = 5
	var num2 int = 5
	if num1 <= num2 {
		fmt.Printf("num1 <= num2\n")
	}	else {
		fmt.Printf("num1 > num2\n")
	}
		
	// Switch statement
	switch {
	case num1 < num2, num1 == num2:
		fmt.Printf("num1 <= num2\n")
		fallthrough
	default:
		fmt.Printf("Default\n")
		
	}
	
	// For loop statement
	for i := 1; i <= 3 ;i++ {
		fmt.Printf("Hello world!!!\n")
	}
	
	// while loop statement
	i = 1
	for i <= 3 {
		fmt.Printf("Hello world!!!\n")
		i++
	}


	// #############################
	// Arrays

	// Option1
	var grades1 [5]int = [5]int {10, 20, 30, 40, 50}
	
	// Option2
	grades2 := [5]int {10, 20, 30, 40, 50}
	
	// Option3 - ... indicates unknown number of array elements on init
	grades3 := [...]int {10, 20, 30, 40, 50}
	
	var fruits [3]string = [3]string {"Apple", "Orange", "Banana"}
	
	fmt.Println(grades1)
	fmt.Println(grades2)
	fmt.Println(grades3)
	fmt.Println(fruits)
	
	fmt.Println(len(fruits))
	
	for index, element := range grades1 {
		fmt.Println(index ,"=>", element)
	}
	
	
	arr := [3][2] int{{2, 4}, {4, 16}, {8, 64}}
	for index, element := range arr {
		fmt.Println(index ,"=>", element)
	}
	
	// #############################
	// Slices
	// #############################
	
	// Declaration options
	// ###################
	// Option 1
	slice := [] int{1, 2, 3}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
	// Option 2
	grades4 := [10]int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := grades4[1:8]
	fmt.Println(slice_1)
	fmt.Println(len(slice_1))
	fmt.Println(cap(slice_1))
	
	// Slice of a slice
	slice_2 := slice_1[0:3]
	fmt.Println(slice_2)
	
	// Option 3
	slice_4 := make([]int, 5, 10)
	fmt.Println(slice_4)
	fmt.Println(len(slice_4))
	fmt.Println(cap(slice_4))
	
	// Appending slices
	// ################

	// Option 1
	grades5 := [4]int {10, 20, 30, 40}
	slice_5 := grades5[1:3]
	sub_slice_1 := append(slice_5, 900)
	sub_slice_2 := append(slice_5, -9, 500)
	
	fmt.Println()
	fmt.Println("grades5 = ", grades5)
	fmt.Println("len(grades5) = ", len(grades5))
	fmt.Println("cap(grades5) = ", cap(grades5))
	fmt.Println()
	
	fmt.Println("slice_5 = ", slice_5)
	fmt.Println("len(slice_5) = ", len(slice_5))
	fmt.Println("cap(slice_5) = ", cap(slice_5))
	
	fmt.Println()
	fmt.Println("sub_slice_1 = ", sub_slice_1)
	fmt.Println("len(sub_slice_1) = ", len(sub_slice_1))
	fmt.Println("cap(sub_slice_1) = ", cap(sub_slice_1))
	
	fmt.Println()
	fmt.Println("sub_slice_2 = ", sub_slice_2)
	fmt.Println("len(sub_slice_2) = ", len(sub_slice_2))
	fmt.Println("cap(sub_slice_2) = ", cap(sub_slice_2))
	
	// Appending 2 slices with  "..." operator
	sub_slice_3 := append(sub_slice_1, sub_slice_2...)
	fmt.Println()
	fmt.Println("sub_slice_3 = ", sub_slice_3)
	fmt.Println("len(sub_slice_3) = ", len(sub_slice_3))
	fmt.Println("cap(sub_slice_3) = ", cap(sub_slice_3))
	
	// Option 2 - "removing" element grades6[2]
	grades6 := [5]int {10, 20, 30, 40, 50}
	i2 := 2
	fmt.Println()
	fmt.Println("\"removing\" element grades6[2]")
	fmt.Println("###############################")
	fmt.Println("grades6 = ", grades6)
	
	sub_slice_4 := grades6[:i2]
	sub_slice_5 := grades6[i2+1:]
	new_slice := append(sub_slice_4, sub_slice_5...)
	fmt.Println("new_slice = ", new_slice)
	fmt.Println("len(new_slice) = ", len(new_slice))
	fmt.Println("cap(new_slice) = ", cap(new_slice))


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
	fmt.Println("num of copied elements = ", num)


	
}
