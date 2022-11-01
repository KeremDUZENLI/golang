package main

import (
	"net/http"
)

func WriteToServer(writeinserver http.ResponseWriter, receivefromserver *http.Request) {
	writeinserver.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/ok", WriteToServer)
	http.ListenAndServe(":9999", nil)
}
