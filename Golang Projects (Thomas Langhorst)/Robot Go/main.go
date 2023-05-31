package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MouseSleep = 100
	robotgo.KeySleep = 100

	currentMousePosition()
	copyAndPaste()
	writeText()
}

func currentMousePosition() {
	for {
		robotgo.MilliSleep(200)
		fmt.Println(robotgo.GetMousePos())
	}
}

func copyAndPaste() {
	robotgo.Move(1767, 290)
	robotgo.Click("left", false)
	robotgo.Click("left", true)
	robotgo.KeyTap("c", "control")
	robotgo.Move(1883, 285)
	robotgo.Click("left", false)
	robotgo.KeyTap("v", "control")
}

func writeText() {
	robotgo.Move(1908, 286)
	robotgo.Click("left", false)
	robotgo.TypeStr("hello world")
}
