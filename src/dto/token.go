package dto


type Token struct {
	Value interface{}
	Symbol string
}

func NewEmptyToken(symbol string) *Token{
	return &Token{Symbol:symbol}
}
