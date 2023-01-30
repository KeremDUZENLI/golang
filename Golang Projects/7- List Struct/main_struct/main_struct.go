package main_struct

const (
	name_1 money = "BITCOIN"
	name_2 money = "ETHEREUM"
)

type money string

type Structe struct {
	liste []money
}

func AddStruct1() Structe {
	liste := []money{name_1, name_2}

	b := Structe{liste: liste}
	return b
}

func AddStruct2() []Structe {
	liste := []money{name_1, name_2}

	liste_of_struct := []Structe{}

	liste_of_struct = append(liste_of_struct, Structe{})
	liste_of_struct[0].liste = liste

	return liste_of_struct
}
