package Maps_Structs

import (
	"fmt"
)

// Map
func StatePopulation() {
	statePopulations := make(map[string]int, 10)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	delete(statePopulations, "Illinois")

	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])

	pop, ok := statePopulations["Oho"]
	fmt.Println(pop, ok)
}
