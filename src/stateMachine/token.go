package stateMachine

import "grammar"

type Token struct {
	specialChar byte
	kindCode int
	_type string
	value interface{}
}


func NewToken(specialChar byte,word string) *Token{
	return &Token{
		specialChar:specialChar,
		value: word,
		kindCode:grammar.GetRegexpsManager().GetCode(word,specialChar),
		_type:grammar.GetRegexpsManager().GetType(specialChar),
	}
}
func NewEmptyToken() *Token{
	return &Token{
		specialChar:0,
		value: "",
		kindCode:0,
		_type:"END",
	}
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
