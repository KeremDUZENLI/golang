package main

import "example/t"

func main() {
	sumNormal()
	sumT()
	typeCheck()
	max()
	maxWeek()
}

func sumNormal() {
	println("\n***********************")
	println(t.SumInt(1, 2))
	println(t.SumFloat64(1.0, 2.0))
	println(t.SumString("1", "2"))
}

func sumT() {
	println("\n***********************")
	println(t.SumT(1, 2))
	println(t.SumT(1.0, 2.0))
	println(t.SumT("1", "2"))
}

func typeCheck() {
	println("\n***********************")
	t.TypeT(1)
	t.TypeT(1.0)
	t.TypeT("1")
}

func max() {
	println("\n***********************")
	println(t.Max([]int{1, 2, 3}))
}

func maxWeek() {
	println("\n***********************")
	println(t.MaxWeek([]t.Weekday{t.Monday, t.Thursday, t.Sunday}))
}
