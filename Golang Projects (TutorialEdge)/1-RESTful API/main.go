package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func BEGINURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "HOMEPAGE")
}

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var TestArticle Article

type Article_Array []Article

var TestArray Article_Array = []Article{}

func GETURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "GET_TEST")

	json.NewEncoder(writer).Encode(TestArray)
}

func POSTURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "POST_TEST")

	json.NewDecoder(reader.Body).Decode(&TestArticle)
	json.NewEncoder(writer).Encode(TestArticle)

	TestArray = append(TestArray, TestArticle)

	writer.Header().Set("Content-Type", "application/json")
}

func DELETEURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "DELETE_TEST")

	url := reader.RequestURI[1:]

	for i := 0; i < len(TestArray); i++ {
		if url == TestArray[i].Title {
			TestArray = append(TestArray[:i], TestArray[i+1:]...)
		}
	}
}

func URL() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", BEGINURL)
	muxRouter.HandleFunc("/A", GETURL).Methods("GET")
	muxRouter.HandleFunc("/A", POSTURL).Methods("POST")
	muxRouter.HandleFunc("/{title}", DELETEURL).Methods("DELETE")

	http.ListenAndServe(":8888", muxRouter)
}

func main() {
	URL()
}
