package tools

import "fmt"

func Input(varNew *string) {
	var path string
	print("Folder Path:   ")
	fmt.Scanln(&path)

	*varNew = path
}

func Count(liste []string, findVar string) int {
	dict := make(map[string]int)

	for _, str := range liste {
		dict[str] += 1
	}

	return dict[findVar]
}
