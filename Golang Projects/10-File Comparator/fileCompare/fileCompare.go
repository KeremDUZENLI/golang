package fileCompare

import (
	"FileComparator/tools"
	"fmt"
)

func FileCompare(liste1 []string, liste2 []string) {
	for _, i := range liste1 {
		if tools.Count(liste2, i) == 0 {
			fmt.Printf("\ndifferent in path1: %v", i)
		}
	}

	fmt.Print("\n")

	for _, l := range liste2 {
		if tools.Count(liste1, l) == 0 {
			fmt.Printf("\ndifferent in path2: %v", l)
		}
	}

	fmt.Print("\n\n")
}
