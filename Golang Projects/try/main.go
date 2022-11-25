package main

import (
	"fmt"
	"time"
)

type neowayEngineLegalPersonMapper struct{}

/* ##########   STRUCT & METHOD / Query Struct  ########## */
type Query struct {
	ID          string
	Integration string
	Type        string
	Documnet    string
	QueriedAt   time.Time
	Data        any
}

type Mapper[T any] interface {
	QueriesToDTO([]Query) (T, error)
}

/* ##########   LIST / CnpjEngineResponse Struct  ########## */
type Log struct {
	Description string
	Field       []string
	Value       interface{}
}

type CnpjEngineResponse struct {
	Id        string
	Logs      []Log
	Score     int
	Status    string
	QueriedAt time.Time
}

type CnpjEnginesResponse []CnpjEngineResponse
type NeowayEngineLegalPersonMapper Mapper[CnpjEnginesResponse]

/* ##########   FUNCTIONS   ########## */
func (neowayEngineLegalPersonMapper) QueriesToDTO(queries []Query) (CnpjEnginesResponse, error) {
	queriesDTO := CnpjEnginesResponse{}

	return queriesDTO, nil
}

func main() {
	newStruct := neowayEngineLegalPersonMapper{}
	query1 := Query{ID: "a"}
	query2 := Query{ID: "b"}
	liste := []Query{query1, query2}

	fmt.Println(newStruct.QueriesToDTO(liste))
}
