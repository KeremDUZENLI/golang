package Challenge05_ImplementingAStackInGo

import "errors"

type Stack struct {
	Flights []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("list is empty")
	}

	last := len(s.Flights) - 1
	popped := s.Flights[last]
	s.Flights = s.Flights[:last]

	return popped, nil
}

func (s *Stack) Push(f Flight) {
	s.Flights = append(s.Flights, f)
}

func (s *Stack) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("list is empty")
	} else {
		last := len(s.Flights) - 1
		return s.Flights[last], nil
	}
}

func (s *Stack) IsEmpty() bool {
	if len(s.Flights) == 0 {
		return true
	} else {
		return false
	}
}

func Test() *Stack {
	newStack := &Stack{
		Flights: []Flight{
			{
				Origin:      "F1",
				Destination: "F2",
				Price:       100,
			},
			{
				Origin:      "F3",
				Destination: "F4",
				Price:       200,
			},
			{
				Origin:      "F5",
				Destination: "F6",
				Price:       300,
			},
		},
	}

	newFlight := &Flight{
		Origin:      "Y",
		Destination: "Z",
		Price:       400,
	}

	newStack.Pop()
	newStack.Push(*newFlight)

	return newStack
}
