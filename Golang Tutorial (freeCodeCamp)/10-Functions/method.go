package Functions

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

func Method() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

// Local Functions
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
