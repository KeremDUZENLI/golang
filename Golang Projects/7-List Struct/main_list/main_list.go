package main_list

const (
	name_1 money = "BITCOIN"
	name_2 money = "ETHEREUM"
)

type money string

func AddList() []money {
	liste := []money{name_1, name_2}
	return liste
}
