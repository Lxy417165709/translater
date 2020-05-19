package grammar


type Token struct {
	specialChar byte
	kindCode int
	value interface{}
}

func (t *Token)Copy() *Token{
	token := &Token{
		t.specialChar,
		t.kindCode,
		t.value,
	}
	return token
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
func (t *Token) SetValue(value interface{}) {
	t.value = value
}
