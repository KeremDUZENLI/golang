package main

import (
	"net/http"
	"server/server"

	"github.com/gorilla/mux"
)

func URL() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", server.BEGINURL)
	muxRouter.HandleFunc("/users/", server.GETURL).Methods("GET")
	muxRouter.HandleFunc("/users/", server.POSTURL).Methods("POST")
	muxRouter.HandleFunc("/users/{name}", server.DELETEURL).Methods("DELETE")

	http.ListenAndServe(":9999", muxRouter)
}

func main() {
	URL()
}
