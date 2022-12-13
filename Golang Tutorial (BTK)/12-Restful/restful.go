package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"usrid"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func RestfulGet() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	// fmt.Println(bodyBytes)
	fmt.Println(bodyString)

	var newStruct1 Todo
	json.Unmarshal(bodyBytes, &newStruct1)
	fmt.Println(newStruct1)
}

func RestfulPost() {
	newStructPost := Todo{UserID: 2, ID: 2, Title: "Shopping", Completed: false}
	json_newStructPost, err := json.Marshal(newStructPost)
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(json_newStructPost))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	// fmt.Println(bodyBytes)
	fmt.Println(bodyString)

	var newStruct2 Todo
	json.Unmarshal(bodyBytes, &newStruct2)
	fmt.Println(newStruct2)
}
