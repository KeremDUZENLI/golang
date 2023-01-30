package main

import (
	"fmt"
)

type Generic[C any] struct {
	Scores C
}

func (c *Generic[C]) Loop(scoreslist C) C {
	c.Scores = scoreslist
	return c.Scores
}

func GenericFunc() *Generic[any] {
	return &Generic[any]{}
}

func main2() {
	scoresIELTS := []float64{4.0, 4.5, 5.0, 5.5, 6.0, 6.5, 7.0, 7.5, 8.0, 8.5, 9.0}
	scoresTOEFL := [][]int{{0, 31}, {32, 34}, {35, 45}, {46, 59}, {60, 78}, {79, 93}, {94, 101}, {102, 109}, {110, 114}, {115, 117}, {118, 120}}

	fmt.Println(GenericFunc().Loop(scoresIELTS))
	fmt.Println(GenericFunc().Loop(scoresTOEFL))
}

/*
func LOOP(anyList []any) {
	for x := 0; x <= 10; x++ {
		fmt.Println(anyList[x])
	}
}

func IELTS() {
	scoresIELTS := []float64{4.0, 4.5, 5.0, 5.5, 6.0, 6.5, 7.0, 7.5, 8.0, 8.5, 9.0}

	for i := 0; i <= 10; i++ {
		// time.Sleep(1 * time.Second)
		fmt.Printf("%.1f\n", scoresIELTS[i])
	}
}

func TOEFL() {
	scoresTOEFL := [][]int{{0, 31}, {32, 34}, {35, 45}, {46, 59}, {60, 78}, {79, 93}, {94, 101}, {102, 109}, {110, 114}, {115, 117}, {118, 120}}

	for l := 0; l <= 10; l++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(scoresTOEFL[l])
	}
}
*/
