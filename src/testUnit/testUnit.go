package testUnit

import (
	"lexicalTest"
)

type TestUnit struct {
	regex   string
	pattern string
	isMatch bool
}

func NewTestUnit(regex string, pattern string, isMatch bool) *TestUnit {
	return &TestUnit{regex, pattern, isMatch}
}


func (unit *TestUnit) test() bool {
	nfaBuilder := lexicalTest.NewNFABuilder(unit.regex)
	nfa := nfaBuilder.BuildNFA()

	return nfa.IsMatch(unit.pattern) == unit.isMatch
}






