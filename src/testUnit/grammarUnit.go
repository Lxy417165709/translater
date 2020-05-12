package testUnit

type GrammarUnit struct{
	identity byte
	regexp   string
}
func NewGrammarUnit(identity byte,regexp string) *GrammarUnit{
	return &GrammarUnit{identity,regexp}
}
