package Challenge21_JSONAndStockDividends

import (
	"encoding/json"
)

type Stock struct {
	Ticker   string  `json:"ticker"`
	Dividend float64 `json:"dividend"`
}

func HighestDividend(str string) Stock {
	var newStruct Stock
	var newListe []Stock

	json.Unmarshal([]byte(str), &newListe)

	for _, stock := range newListe {
		if stock.Dividend > newStruct.Dividend {
			newStruct = stock
		}
	}

	return newStruct
}

func Test() Stock {
	stocks_json := `[
    {"ticker":"APPL","dividend":0.5},
    {"ticker":"GOOG","dividend":0.2},
    {"ticker":"MSFT","dividend":0.3}
  ]`

	return HighestDividend(stocks_json)
}
