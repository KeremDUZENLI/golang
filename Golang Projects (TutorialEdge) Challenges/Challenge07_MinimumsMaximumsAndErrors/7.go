package Challenge07_MinimumsMaximumsAndErrors

import "errors"

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	if len(flights) == 0 {
		return 0, 0, errors.New("empty flight list")
	} else {
		min := flights[0].Price
		max := flights[0].Price

		for _, i := range flights {
			if max < i.Price {
				max = i.Price
			}

			if min > i.Price {
				min = i.Price
			}
		}

		return min, max, nil
	}
}

func Test() (int, int, error) {
	flights := []Flight{
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
	}

	return GetMinMax(flights)
}
