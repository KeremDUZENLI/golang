package main

import (
	"fmt"
	"net/http"
)

func BEGINURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "HOMEPAGE")
}

func main() {
	http.HandleFunc("/", BEGINURL)
	http.ListenAndServe(":9999", nil)
}
