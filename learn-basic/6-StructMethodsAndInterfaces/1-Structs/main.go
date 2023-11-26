package main

import "fmt"

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

func main() {
	
	// Declaring and Initializing a Struct
	
	// Option 1
	var student_1 Student
	student_1.name = "Shay"
	student_1.grades = map[string]int{"English": 95, "Math": 99}
	fmt.Printf("%+v\n", student_1)
	modify(&student_1)
	fmt.Printf("%+v\n", student_1)
	fmt.Println()
	

	// Option 2 
	student_2 := Student {
		name: "Shay",
		rollNo: 1,
		marks: []int{1,2,3},
		grades: map[string]int{"English": 95, "Math": 99},
	}
	fmt.Printf("%+v\n",student_2)
	fmt.Println()
	
	// Option 3 
	student_3 := Student {
		"Shay", 
		1,
		[]int{1,2,3},
		map[string]int{"English": 95, "Math": 99},
	}
	fmt.Printf("%+v\n",student_3)
	fmt.Println()
	
	// Option 4
	// Pointer to the Student object
	student_4_ptr := new(Student)
	student_4_ptr.name = "Shay"
	student_4_ptr.grades = map[string]int{"English": 95, "Math": 99}
	fmt.Printf("%+v\n", student_4_ptr)
	modify(student_4_ptr)
	fmt.Printf("%+v\n", student_4_ptr)
	fmt.Println()

}


func modify(st_ptr *Student) {
	st_ptr.name = "Changed"
}