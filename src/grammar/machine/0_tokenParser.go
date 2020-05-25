package machine

import (
	"bytes"
	"conf"
	"grammar/char"
)

type TokenParser struct {
	specialCharTable *char.SpecialCharTable
	specialChars     []byte
	finalNFA         *NFA

	text            []byte
	preEndState     *state
	stateQueue      []*state
	bufferOfChars   bytes.Buffer
	readingPosition int
	finalTokens          []*Token
}


// TODO: specialChars可以加入配置
func NewTokenParser() *TokenParser {
	specialChars := []byte(conf.GetConf().LexicalConf.SpecialCharsOfNFAs)
	return &TokenParser{
		specialChars:     specialChars,
		specialCharTable: char.NewSpecialCharTable(),
		finalNFA:         NewNFABuilder().BuildNFABySpecialChars(specialChars).EliminateBlankStates(),
	}
}

