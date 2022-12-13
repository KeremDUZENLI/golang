package main

import (
	"fmt"
	"net/http"

	"orm/users"

	"github.com/gorilla/mux"
)

func BEGINURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}

func URL() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", BEGINURL)

	muxRouter.HandleFunc("/users", users.AllUser).Methods("GET")
	muxRouter.HandleFunc("/users/{name}/{email}", users.NewUser).Methods("POST")
	muxRouter.HandleFunc("/users/{name}", users.DeleteUser).Methods("DELETE")
	muxRouter.HandleFunc("/users/{name}/{email}", users.UpdateUser).Methods("PUT")

	http.ListenAndServe(":8081", muxRouter)
}

func main() {
	users.CreateDatabase()
	users.SendDatabase()

	URL()
}
