package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)

	pl(sV2)
	pl("length: ", len(sV2))
	pl("Contains Another: ", strings.Contains(sV2, "Another"))
	pl("o index: ", strings.Index(sV2, "o"))
	pl("Replace: ", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n" // \t \"
	sV3 = strings.TrimSpace(sV3)
	pl("Split: ", strings.Split("a-b-c-d", "-"))
	pl("Lower: ", strings.ToLower(sV2))
	pl("Upper: ", strings.ToUpper(sV2))
	pl("Prefix: ", strings.HasPrefix("tacocat", "taco"))
}
