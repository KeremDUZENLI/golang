package Maps_Structs

import (
	"fmt"
)

// Struct
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func Doctors() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Jon Pertwee",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
}
