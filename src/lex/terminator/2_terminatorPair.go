package terminator

type Pair struct {
	terminator string
	value interface{}
}
func NewPair(terminator string,value interface{}) *Pair {
	return &Pair{
		terminator:terminator,
		value:value,
	}
}
func (tp *Pair) GetSymbol() string{
	return tp.terminator
}
func (tp *Pair) GetValue() interface{}{
	return tp.value
}

