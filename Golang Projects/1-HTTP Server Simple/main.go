package main

import (
	"net/http"
	"server/database"
	"server/server"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func URL() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", server.BEGINURL)
	muxRouter.HandleFunc("/1", server.GETURL).Methods("GET")
	muxRouter.HandleFunc("/1", server.POSTURL).Methods("POST")
	muxRouter.HandleFunc("/{name}", server.DELETEURL).Methods("DELETE")

	http.ListenAndServe(":9999", muxRouter)
}

func main() {
	database.CreateDatabase()
	database.SendDatabase()

	URL()
}
