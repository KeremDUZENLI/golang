package Challenge06_ImplementingAQueueInGo

import (
	"errors"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type Queue struct {
	Items []Flight
}

func (q *Queue) IsEmpty() bool {
	if len(q.Items) == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is empty")
	}

	flight := q.Items[0]
	q.Items = q.Items[1:]
	return flight, nil
}

func (q *Queue) Push(f Flight) {
	q.Items = append(q.Items, f)
}

func (q *Queue) Peek() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is empty")
	}

	return q.Items[0], nil
}

func Test() *Queue {
	newQueue := &Queue{
		Items: []Flight{
			{
				Origin:      "A",
				Destination: "B",
				Price:       100,
			},
			{
				Origin:      "C",
				Destination: "D",
				Price:       200,
			},
			{
				Origin:      "E",
				Destination: "F",
				Price:       300,
			},
		},
	}

	newFlight := &Flight{
		Origin:      "Y",
		Destination: "Z",
		Price:       400,
	}

	newQueue.Pop()
	newQueue.Push(*newFlight)

	return newQueue
}
