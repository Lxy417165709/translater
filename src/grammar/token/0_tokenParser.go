package token

import (
	"bytes"
	"conf"
	"grammar/char"
	"grammar/machine"
)

type TokenParser struct {
	specialCharTable *char.SpecialCharTable
	finalNFA         *machine.NFA

	text            []byte
	preEndState     *machine.State
	stateQueue      []*machine.State
	bufferOfChars   bytes.Buffer
	readingPosition int
	finalTokens          []*Token
}


// TODO: specialChars可以加入配置
func NewTokenParser() *TokenParser {
	specialChars := []byte(conf.GetConf().LexicalConf.SpecialCharsOfNFAs)
	return &TokenParser{
		specialCharTable: char.NewSpecialCharTable(),
		finalNFA:         machine.NewNFABuilder().BuildNFABySpecialChars(specialChars).EliminateBlankStates(),
	}
}

