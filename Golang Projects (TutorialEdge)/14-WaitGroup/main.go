package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp.Status)
			wg.Done()
		}(url)
	}

	wg.Wait()
}

func main() {
	fmt.Println("Waitgroup")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func myFunc(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Finished")
// 	wg.Done()
// }

// func main() {
// 	fmt.Println("GO")
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go myFunc(&wg)
// 	wg.Wait()
// 	fmt.Println("MAIN")
// }
