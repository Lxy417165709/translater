package lexical


type pair struct {
	kind string
	kindCode int
	value interface{}
}

func NewPair(kind string,kindCode int,value interface{}) *pair{
	return &pair{kind,kindCode,value}
}
func (p *pair) GetKind() string{
	return p.kind
}
func (p *pair) GetKindCode() int{
	return p.kindCode
}
func (p *pair) GetValue() interface{}{
	return p.value
}
