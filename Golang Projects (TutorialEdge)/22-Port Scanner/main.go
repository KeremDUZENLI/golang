package main

import (
	"fmt"
	"portscanner/port"
)

func main() {
	fmt.Println("Port Scanner")

	// open := port.ScanPort("tcp", "localhost", 8080)
	// fmt.Printf("Port Open: %t\n", open)

	results := port.InitialScan("localhost")
	// fmt.Println(results)

	for i := 0; i < len(results); i++ {
		x := results[i]
		if x.State == "Open" {
			fmt.Println("\n", x.Port)
		}
	}
}
