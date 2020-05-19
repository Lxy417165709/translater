package grammar


type Token struct {
	specialChar byte
	kindCode int
	_type string
	value interface{}
}

func (t *Token)Copy() *Token{
	token := &Token{
		t.specialChar,
		t.kindCode,
		t._type,
		t.value,
	}
	return token
}
func (t *Token) GetType() string{
	return t._type
}
func (t *Token) GetSpecialChar() byte{
	return t.specialChar
}
func (t *Token) GetKindCode() int{
	return t.kindCode
}
func (t *Token) GetValue() interface{}{
	return t.value
}
func (t *Token) SetValue(value interface{}) {
	t.value = value
}
