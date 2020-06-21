package token

import (
	"conf"
	"grammar/char"
	"grammar/machine"
)

type Parser struct {
	specialCharTable *char.SpecialCharTable
	finalNFA         *machine.NFA
}

func NewParser() *Parser {
	return &Parser{
		specialCharTable: char.NewSpecialCharTable(),
		finalNFA:         machine.NewNFABuilder().BuildNFABySpecialChars([]byte(conf.GetConf().LexicalConf.SpecialCharsOfNFAs)).EliminateBlankStates(),
	}
}
