package route

import (
	controller "http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func URL() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", controller.GetURL).Methods("GET")
	muxRouter.HandleFunc("/{ID}", controller.GetIDURL).Methods("GET")
	muxRouter.HandleFunc("/", controller.PostURL).Methods("POST")
	muxRouter.HandleFunc("/{ID}", controller.DeleteURL).Methods("DELETE")

	http.ListenAndServe(":9001", muxRouter)
}
