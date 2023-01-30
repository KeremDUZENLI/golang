package main

import "fmt"

func main() {
	scoresIELTS := &Generic{liste: []interface{}{4.0, 4.5, 5.0, 5.5, 6.0, 6.5, 7.0, 7.5, 8.0, 8.5, 9.0}}
	scoresIELTS.Loop()

	scoresTOEFL := &Generic{liste: []interface{}{[]interface{}{0, 31}, []interface{}{32, 34}, []interface{}{35, 45},
		[]interface{}{46, 59}, []interface{}{60, 78}, []interface{}{79, 93}, []interface{}{94, 101},
		[]interface{}{102, 109}, []interface{}{110, 114}, []interface{}{115, 117}, []interface{}{118, 120}}}
	scoresTOEFL.Loop()
}

type Generic struct {
	liste []interface{}
}

func (t *Generic) Loop() {
	for _, item := range t.liste {
		fmt.Println(item)
	}
}
