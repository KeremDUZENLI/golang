package main

import (
	"fmt"
)

/*
func (h recoveryHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.log(err)
		}
	}()

	h.Handler.ServerHTTP(w, req)
}
*/

func main() {
	fmt.Println("Panic!")

	defer func() {
		fmt.Println("Deferred")
	}()
	panic("oh no")
	fmt.Println("end of main")
}
