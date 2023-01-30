package variable

import "fmt"

type Playingcard struct {
	Suit string
	Rank string
}

func (pc *Playingcard) String() string {
	return fmt.Sprintf("%s -- %s", pc.Suit, pc.Rank)
}

func (pc *Playingcard) Name() string {
	return pc.String()
}
