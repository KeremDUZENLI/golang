package maps

import "fmt"

func Map() {
	dictionary := make(map[string]string)
	dictionary["table"] = "mesa"
	dictionary["book"] = "livro"
	dictionary["pencil"] = "lapis"

	fmt.Println(dictionary["book"])
	fmt.Printf("Number of Items : %v\n", len(dictionary))
	fmt.Printf("Dictionary : %v\n\n", dictionary)

	delete(dictionary, "book")
	fmt.Printf("Number of Items : %v\n", len(dictionary))
	fmt.Printf("Dictionary : %v\n\n", dictionary)

	value, exist := dictionary["pencil"]
	fmt.Println(value)
	fmt.Println(exist)
}
