package main

import (
	"fmt"
	Variables "golang_tutorial_2/1-Variables"
	Functions "golang_tutorial_2/10-Functions"
	Interfaces "golang_tutorial_2/11-Interfaces"
	GoRoutines "golang_tutorial_2/12-GoRoutines"
	Channels "golang_tutorial_2/13-Channels"
	Primitives "golang_tutorial_2/2-Primitives"
	Constants "golang_tutorial_2/3-Constants"
	Arrays_Slices "golang_tutorial_2/4-Arrays_Slices"
	Maps_Structs "golang_tutorial_2/5-Maps_Structs"
	If_Switch "golang_tutorial_2/6-If_Switch"
	Loop "golang_tutorial_2/7-Loop"
	Defer_Panic_Recover "golang_tutorial_2/8-Defer_Panic_Recover"
	Pointers "golang_tutorial_2/9-Pointers"
)

func main() {
	fmt.Println("\n1-Variables:")
	Variables.Variable1()
	Variables.Variable2()
	Variables.VariableConversion()

	fmt.Println("\n2-Primitives:")
	Primitives.Boolean()
	Primitives.Integer()
	Primitives.Operator()
	Primitives.String()

	fmt.Println("\n3-Constants:")
	Constants.Constant()
	Constants.Iota()
	Constants.IotaSize()

	fmt.Println("\n4-Arrays_Slices:")
	Arrays_Slices.Array()
	Arrays_Slices.Slice()
	Arrays_Slices.Make()

	fmt.Println("\n5-Maps_Structs:")
	Maps_Structs.StatePopulation()
	Maps_Structs.Doctors()
	Maps_Structs.BirdFly()

	fmt.Println("\n6-If_Switch:")
	If_Switch.NumberGuess1()
	If_Switch.NumberGuess2()
	If_Switch.Switch1()
	If_Switch.Switch2()

	fmt.Println("\n7-Loop:")
	Loop.Loop()
	Loop.Collection()

	fmt.Println("\n8-Defer_Panic_Recover:")
	Defer_Panic_Recover.Defer()
	// Defer_Panic_Recover.Panic() // Starts Server
	// Defer_Panic_Recover.Recover() // Recovers Error

	fmt.Println("\n9-Pointers:")
	Pointers.Pointer()

	fmt.Println("\n10-Functions:")
	Functions.Function()
	Functions.Method()

	fmt.Println("\n11-Interfaces:")
	Interfaces.Interface()
	Interfaces.Composing()
	// Interfaces.Conversion() // Complicated Conversion
	Interfaces.Sample()

	fmt.Println("\n12-GoRoutines:")
	GoRoutines.GoRoutine()
	GoRoutines.SingleThread()
	GoRoutines.MultipleThread()

	fmt.Println("\n13-Channels:")
	Channels.Channel()
	Channels.BufferedChannel()
	Channels.LoopChannel()
	Channels.LogChannel()
}
