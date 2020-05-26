package token

import (
	"conf"
	"grammar/char"
	"grammar/machine"
)

type TokenParser struct {
	specialCharTable *char.SpecialCharTable
	finalNFA         *machine.NFA
}

func NewTokenParser() *TokenParser {
	specialChars := []byte(conf.GetConf().LexicalConf.SpecialCharsOfNFAs)
	return &TokenParser{
		specialCharTable: char.NewSpecialCharTable(),
		finalNFA:         machine.NewNFABuilder().BuildNFABySpecialChars(specialChars).EliminateBlankStates(),
	}
}
