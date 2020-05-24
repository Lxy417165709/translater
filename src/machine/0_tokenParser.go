package machine

import (
	"bytes"
	"conf"
	"grammar"
)

type TokenParser struct {
	specialCharTable *grammar.SpecialCharTable
	specialChars     []byte
	finalNFA         *NFA

	text            []byte
	preEndState     *state
	stateQueue      []*state
	bufferOfChars   bytes.Buffer
	readingPosition int
	tokens          []*Token
}


// TODO: specialChars可以加入配置
func NewTokenParser(cf *conf.Conf) *TokenParser {
	specialChars := []byte(cf.LexicalConf.SpecialCharsOfNFAs)
	return &TokenParser{
		specialChars:     specialChars,
		specialCharTable: grammar.NewSpecialCharTable(cf),
		finalNFA:         NewNFABuilder(cf).BuildNFABySpecialChars(specialChars).EliminateBlankStates(),
	}
}

