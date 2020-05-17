package lexical


type token struct {
	kind string
	kindCode int
	value interface{}
}

func NewPair(kind string,kindCode int,value interface{}) *token{
	return &token{kind,kindCode,value}
}
func (p *token) GetKind() string{
	return p.kind
}
func (p *token) GetKindCode() int{
	return p.kindCode
}
func (p *token) GetValue() interface{}{
	return p.value
}
