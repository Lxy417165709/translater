package lex


type TerminatorPair struct {
	terminator string
	value interface{}
}
func NewTerminatorPair(terminator string,value interface{}) *TerminatorPair {
	return &TerminatorPair{
		terminator:terminator,
		value:value,
	}
}
func (tp *TerminatorPair) GetSymbol() string{
	return tp.terminator
}
func (tp *TerminatorPair) GetValue() interface{}{
	return tp.value
}

