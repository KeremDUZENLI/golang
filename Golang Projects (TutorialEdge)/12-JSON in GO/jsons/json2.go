package jsons

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}

type Info struct {
	Description string `json:"description"`
}

func Json2() {
	fmt.Println("SensorReading")

	jsonString := `{"name": "battery sensor", "capacity": 40, "time": 
    "2022-11-23T16:53:30Z", "info": {
        "description": "XXX"
    }}`

	var read SensorReading
	// var read map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &read)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", read)
}
