package variable

type Tradingcard struct {
	Collectablename string
}

func (tc *Tradingcard) String() string {
	return tc.Collectablename
}

func (tc *Tradingcard) Name() string {
	return tc.String()
}
