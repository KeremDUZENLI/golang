package Arrays_Slices

import (
	"fmt"
)

func Array() {
	var students [3]string
	fmt.Printf("Students: %v\n", students)

	students[0] = "Lisa"
	students[1] = "Ahmet"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[0])
	fmt.Printf("Number of students: %v\n", len(students))

	grades := [...]int{80, 90, 100}
	fmt.Printf("Grades: %v\n", grades)
}
