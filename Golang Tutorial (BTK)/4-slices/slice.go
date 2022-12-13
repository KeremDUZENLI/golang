package slices

import "fmt"

func Slice1() {
	names := []string{"1", "2", "3"}
	fmt.Printf("Names 1 	 	: %v\n", names)

	namesCopy := make([]string, len(names))
	fmt.Printf("Names Copy 1 : %v\n", namesCopy)

	copy(namesCopy, names)
	fmt.Printf("Copy 		 	: %v\n", namesCopy)

	names = append(names, "4")
	fmt.Printf("Append		 	: %v\n", names)

	fmt.Printf("Slices		 	: %v\n", names[1:3])
}
