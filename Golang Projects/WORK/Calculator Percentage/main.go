package main

import "fmt"

type IELTS struct {
	ScoresIELTS []float64
}

type TOEFL struct {
	ScoresTOEFL [][]int
}

type GenericT[T interfac] struct {
	ScoresT T
}

type interfac interface {
}

func (t *GenericT[T]) LOOP(liste T) T {
	t.ScoresT = liste
	return t.ScoresT
}

func GenericFuncT1() *GenericT[*IELTS] {
	return &GenericT[*IELTS]{}
}

func GenericFuncT2() *GenericT[*TOEFL] {
	return &GenericT[*TOEFL]{}
}

func main() {
	newIELTS := &IELTS{
		ScoresIELTS: []float64{4.0, 4.5, 5.0, 5.5, 6.0, 6.5, 7.0, 7.5, 8.0, 8.5, 9.0},
	}

	fmt.Println(GenericFuncT1().LOOP(newIELTS))
}
