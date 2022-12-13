package main

import (
	"bufio"
	"os"

	gt "github.com/bas24/googletranslatefree"
)

func ReadFile() {
	content_read, _ := os.Open("Translating.txt")
	content_write, _ := os.Create("Translated.txt")

	defer content_read.Close()
	defer content_write.Close()

	scanner := bufio.NewScanner(content_read)
	for scanner.Scan() {
		// content_read
		text := scanner.Text()

		// googletranslatefree
		trans, _ := gt.Translate(text, "tr", "en")

		// content_write
		content_write.WriteString(trans + "\n")
	}
}

func main() {
	ReadFile()
}
