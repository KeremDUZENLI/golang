package Challenge03_SortingFlightsByPrice

import (
	"fmt"
	"sort"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type Flights []Flight

func (f Flights) Len() int {
	return len(f)
}

func (f Flights) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Flights) Less(i, j int) bool {
	return f[i].Price > f[j].Price
}

func SortByPrice(flights []Flight) []Flight {
	sort.Sort(Flights(flights))
	return flights
}

func PrintFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d\n", flight.Origin, flight.Destination, flight.Price)
	}
}

func SortedList() []Flight {
	flights := Flights{
		{Origin: "New York", Destination: "Paris", Price: 800},
		{Origin: "London", Destination: "New York", Price: 700},
		{Origin: "Paris", Destination: "London", Price: 600},
	}

	sortedList := SortByPrice(flights)
	return sortedList
	// PrintFlights(sortedList)
}
