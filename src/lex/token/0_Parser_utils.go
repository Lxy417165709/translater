package token

import (
	"grammar/machine"
)

func (tp *Parser) wordPairToToken(wordPair *machine.WordPair) *Token{
	token := &Token{
		wordPair.GetSpecialChar(),
		tp.specialCharTable.GetCode(wordPair.GetSpecialChar(),wordPair.GetWord()),
		tp.specialCharTable.GetType(wordPair.GetSpecialChar()),
		wordPair.GetWord(),
	}
	return token
}



