package regexpsManager


type Token struct {
	specialChar byte
	kindCode int
	value interface{}
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