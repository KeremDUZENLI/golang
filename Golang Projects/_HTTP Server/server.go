package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", r.URL)
	var html, err = loadFile("index.html")

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404 Sorry")
	}

	fmt.Fprint(w, html)
	fmt.Fprintf(w, "Hello %s, welcome to http server", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
