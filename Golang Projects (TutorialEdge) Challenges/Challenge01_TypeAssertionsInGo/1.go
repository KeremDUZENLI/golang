package Challenge01_TypeAssertionsInGo

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name string, age int) Developer {
	newDeveloper := Developer{
		Name: name,
		Age:  age,
	}

	return newDeveloper
}
