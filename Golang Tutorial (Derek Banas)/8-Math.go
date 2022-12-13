package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	pl("5 +4 = ", 5+4)
	pl("5-4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4)

	var mInt int
	mInt++
	pl(mInt)

	// 1/1/70
	seedSec := time.Now().Unix()
	rand.Seed(seedSec)
	randNum := rand.Intn(50) + 1
	pl("Random : ", randNum)

	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)
}
