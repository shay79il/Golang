package main

import (
	"fmt"
	"unsafe"
	"reflect"
	"strconv"
	// "strings"
)

// Global variable
var globalVar string = "I am a global var"
const PI float64 = 3.14
	
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
	// %.2f 	floating numbers up to 2 decimal places
	
	
	// Short variable declaration 
	fmt.Println("// Short variable declaration")
	name_1 := "Shay"
	fmt.Println(name_1)
	fmt.Println()
	
	// Long variable declaration 
	fmt.Println("// Long variable declaration")
	
	// Option 1
	var name_2 string = "Guy"
	fmt.Println(name_2)
	fmt.Println()
	
	// Option 2
	var name_3 string
	name_3 = "Dan"
	fmt.Println(name_3)
	fmt.Println()
	
	// Declaration of 2 vars of the same data type
	fmt.Println("// Declaration of 2 vars of the same data type")
	var str2, str3 string = "Hello", "World!" 
	fmt.Println("str2 = ", str2)
	fmt.Println("str3 = ", str3)
	fmt.Println(str2, str3)
	fmt.Println()
	
	// Declaration of 2 vars of different data types
	fmt.Println("// Declaration of 2 vars of different data types")
	var (
		str4 string = "I am a string"
		intVal int = 5
	)
	fmt.Println("str4 = ", str4)
	fmt.Println("intVal = ", intVal)
	fmt.Println(str4, intVal)
	fmt.Println()
	
	// ------------------------
	var myUint uint
	sizeOfUint := unsafe.Sizeof(myUint)
	fmt.Printf("Size of uint: %d bytes\n", sizeOfUint)
	fmt.Print("Size of uint: ", sizeOfUint, " bytes\n")
	fmt.Println()
	
	// ------------------------
	// Zero values 
	// Uninitialized variables get initial default value as follows
	
	fmt.Println("// Zero values")
	var isOK bool // isOK = 0
	var intVal3 int // intVal3 = 0
	var floatVal float64 // floatVal = 0
	var str7 string // str7 = ""
	fmt.Printf("isOK = %t, intVal3 = %d, floatVal = %.2f, str7 = %q\n",isOK,intVal3,floatVal, str7)
	fmt.Println()
	
	
	// ------------------------
	// User input - option for multiple inputs
	var myName string = "Shay"
	var myAge int = 44
	// fmt.Print("Please enter your myName and myAge: ") 
	// num, err := fmt.Scanf("%s %d", &myName, &myAge)
	// fmt.Println("Hey there,", myName, "your myAge is", myAge) 
	// fmt.Println("Count,", num) 
	// fmt.Println("Errors,", err) 
	
	// ------------------------
	// User input TYPE
	fmt.Println("// User input TYPE")
	fmt.Printf("variable myName = %v is of type %T\n",myName, myName)
	fmt.Printf("variable myAge = %v is of type %T\n",myAge, myAge)
	fmt.Printf("variable isOK = %v is of type %v\n",isOK, reflect.TypeOf(isOK))
	fmt.Println()
	
	// ------------------------
	// Casting
	fmt.Println("// Casting")
	fmt.Printf("variable myAge = %v is of type %T\n",myAge, myAge)
	var floatAge float64 = float64(myAge)
	fmt.Printf("variable floatAge = %.2f is of type %T\n",floatAge, floatAge)
	
	var stringAge string = strconv.Itoa(myAge)
	fmt.Printf("stringAge = %q \n",stringAge)
	
	var stringNumber string = "200"
	intVal4, err := strconv.Atoi(stringNumber)
	fmt.Printf("intVal4 = %v, %T\n",intVal4, intVal4)
	fmt.Printf("err = %v, %T\n",err, err)
	
	var strErrConvert string = "200abc"
	errConvert, err := strconv.Atoi(strErrConvert)
	fmt.Printf("intVal4 = %v, %T\n",errConvert, errConvert)
	fmt.Printf("err = %v, %T\n",err, err)
	fmt.Println()
	
	// ------------------------
	// Constants
	fmt.Println("// Constants")
	const constVar int = 64
	fmt.Printf("const constVar = %v is of type %v\n",constVar, reflect.TypeOf(constVar))
	
	// UNtyped constant
	const name = "Harry Potter"
	const isMuggle = false
	const age = 21
	fmt.Printf("variable myName = %v is of type %T\n",name, name)
	fmt.Printf("variable isMuggle = %v is of type %T\n",isMuggle, isMuggle)
	fmt.Printf("variable age = %v is of type %T\n",age, age)

	// Globals
	fmt.Println("globalVar = ", globalVar)
	fmt.Printf("PI = %.2f", PI)
}