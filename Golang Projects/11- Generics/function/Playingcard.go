package function

import (
	"generics/variable"
)

func PlayingcardFuncNew(x string, y string) variable.Playingcard {
	return variable.Playingcard{
		Suit: x,
		Rank: y,
	}
}

func PlayingcardFuncPoint(x string, y string) *variable.Playingcard {
	return &variable.Playingcard{
		Suit: x,
		Rank: y,
	}
}
