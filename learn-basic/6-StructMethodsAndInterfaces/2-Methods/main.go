package main

import "fmt"


func main() {
	
	student_3 := Student {
		"Shay", 
		1,
		[]int{1,2,3},
		map[string]int{"English": 95, "Math": 99},
	}

	// Use of method to get name 
	// call Student by value - (s Student)
	fmt.Printf("%v\n",student_3.getName())
	fmt.Printf("%+v\n",student_3)
	
	
	// Use of method to change name 
	// call Student by reference - (s *Student)
	student_3.changeName()
	fmt.Printf("%+v\n", student_3)
	fmt.Println()
	
	
}

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

func (s Student) getName() string {
	return s.name
}

func (s *Student) changeName() {
	s.name = "Changed"
}
