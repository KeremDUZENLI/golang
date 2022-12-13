package util

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
)

func ParseBody(reader *http.Request, x interface{}) {
	body, _ := ioutil.ReadAll(reader.Body)
	json.Unmarshal([]byte(body), x)
}
