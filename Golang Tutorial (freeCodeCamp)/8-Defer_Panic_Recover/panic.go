package Defer_Panic_Recover

import (
	"net/http"
)

func Panic() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}
