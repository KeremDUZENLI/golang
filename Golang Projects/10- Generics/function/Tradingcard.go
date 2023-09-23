package function

import (
	"generics/variable"
)

func TradingcardFuncNew(input string) variable.Tradingcard {
	return variable.Tradingcard{
		Collectablename: input,
	}
}

func TradingcardFuncPoint(input string) *variable.Tradingcard {
	return &variable.Tradingcard{
		Collectablename: input,
	}
}
