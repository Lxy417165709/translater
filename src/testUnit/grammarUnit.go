package testUnit

type GrammarUnit struct{
	specialChar byte
	regexp   string
}
func NewGrammarUnit(specialChar byte,regexp string) *GrammarUnit{
	return &GrammarUnit{specialChar,regexp}
}
