package Challenge02_SatisfyingInterfacesInGo

var programmers []Employee

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name   string
	AgeOld int
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return e.AgeOld
}

func Programmer() []Employee {
	candidate := Engineer{
		Name:   "Elliot",
		AgeOld: 30,
	}

	programmers = append(programmers, candidate)

	return programmers
}
