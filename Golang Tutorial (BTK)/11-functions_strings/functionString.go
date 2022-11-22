package functions_strings

import (
	"fmt"
	"strings"
)

// case sensitive
// ascii
func FunctionString() {
	name := "Kerem"

	fmt.Println(strings.Count(name, "e")) // Amount of characters

	fmt.Println(strings.Contains(name, "e")) // If it includes

	fmt.Println(strings.Index(name, "e")) // Index number of character

	fmt.Println(strings.ToLower(name)) // lower case
	fmt.Println(strings.ToUpper(name)) // upper case

	fmt.Println(strings.HasPrefix(name, "Ke")) // Start with
	fmt.Println(strings.HasSuffix(name, "em")) // End with
	fmt.Println(strings.Index(name, "re"))

	letters := []string{"k", "e", "r", "e", "m"}
	result := strings.Join(letters, "*") // Join all the letters
	fmt.Println(result)

	fmt.Println(strings.Replace(result, "*", "+", 2)) // Replace

	fmt.Println(strings.Split(name, "")) // Split

	fmt.Println(strings.Repeat(name, 5)) // Repeat
}
