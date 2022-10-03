package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	pl("\n***Random Number is: ", randNum)
	fmt.Print("Guess a number between + and 50:")

	for true {
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)

		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guessed It")
			break
		}
	}
}
