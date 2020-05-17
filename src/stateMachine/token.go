package stateMachine


type Token struct {
	specialChar byte
	kindCode int
	value interface{}
}

func NewPair(speicalChar byte,kindCode int,value interface{}) *Token{
	return &Token{speicalChar,kindCode,value}
}
func (p *Token) GetSpecialChar() byte{
	return p.specialChar
}
func (p *Token) GetKindCode() int{
	return p.kindCode
}
func (p *Token) GetValue() interface{}{
	return p.value
}
