package lexicalTest

import (
	"strings"
)

const RegexSplitString = "|"

// TODO: 重构
type NFABuilder struct {
	buildRegex  string
	readingRegex string
	readingPosition int
	buildingNFA *NFA
	endChar     byte
}

func NewNFABuilder(buildRegex string) *NFABuilder {
	endChar := byte('#')
	buildRegex += string(endChar)
	return &NFABuilder{
		buildRegex,
		buildRegex,
		0,
		NewNFA(),
		endChar,
	}
}

func (nb *NFABuilder) BuildNFA() *NFA {
	regexs := strings.Split(nb.buildRegex, RegexSplitString)
	if len(regexs) == 1 {
		nb.readingRegex = regexs[0]
		for !nb.readingIsOver(){
			nb.parseChar()
		}
		return nb.buildingNFA
	}
	finalNFA := NewNFABuilder(regexs[0]).BuildNFA()
	for i := 1; i < len(regexs); i++ {
		regex := regexs[i]
		finalNFA.AddParallelGraph(NewNFABuilder(regex).BuildNFA())
	}
	return finalNFA
}

func (nb *NFABuilder) parseChar() {
	baseChar := nb.readingRegex[nb.readingPosition]
	nextChar := nb.readingRegex[nb.readingPosition+1]
	switch {
	case nextChar == '+':
		nb.buildingNFA.RepeatPlus(baseChar)
		nb.readingPosition+=2
	case nextChar == '*':
		nb.buildingNFA.RepeatZero(baseChar)
		nb.readingPosition+=2
	default:
		nb.buildingNFA.Once(baseChar)
		nb.readingPosition++
	}
}


func (nb *NFABuilder) readingIsOver() bool{
	return nb.readingRegex[nb.readingPosition] == nb.endChar
}

