package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("DOCKER")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HELLO")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
