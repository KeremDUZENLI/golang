// package main

// import "net/http"

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello world!"))
// }

// func main() {
// 	http.HandleFunc("/1", handler)
// 	http.ListenAndServe(":8000", nil)
// }

package main

import (
	"net/http"

	"server/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
