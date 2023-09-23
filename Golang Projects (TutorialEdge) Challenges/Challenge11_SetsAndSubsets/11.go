package Challenge11_SetsAndSubsets

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func IsSubset(first, second []Flight) bool {
	set := make(map[Flight]bool)

	for _, value := range second {
		set[value] = true
	}

	for _, value := range first {
		a := set[value]
		if a {
			return true
		} else {
			return false
		}
	}

	return true
}

func Test() bool {
	firstFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
	}

	secondFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
		{Origin: "GLA", Destination: "AMS", Price: 500},
	}

	subset := IsSubset(firstFlights, secondFlights)
	return subset
}
