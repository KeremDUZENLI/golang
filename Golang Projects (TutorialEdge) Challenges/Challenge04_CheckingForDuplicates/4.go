package Challenge04_CheckingForDuplicates

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []Developer {
	duplicated := map[string]bool{}
	result := []Developer{}

	for _, value := range developers {
		_, ok := duplicated[value.Name]
		if !ok {
			result = append(result, value)
			duplicated[value.Name] = true
		}
	}

	return result
}

func UniquesList() []Developer {
	x := []Developer{
		{
			Name: "A",
			Age:  10,
		},
		{
			Name: "B",
			Age:  20,
		},
		{
			Name: "C",
			Age:  30,
		},
		{
			Name: "D",
			Age:  40,
		},
	}

	return FilterUnique(x)
}
