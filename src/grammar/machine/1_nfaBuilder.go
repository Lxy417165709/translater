package machine

import (
	"conf"
	"grammar/char"
)

// TODO: 这个也需要配置


type NFABuilder struct {
	specialCharTable *char.SpecialCharTable
	additionChars []byte
}

func NewNFABuilder() *NFABuilder {
	return &NFABuilder{
		specialCharTable: char.NewSpecialCharTable(),
		additionChars:[]byte{
			conf.GetConf().GrammarConf.MatchMoreThanOnceSymbol[0],
			conf.GetConf().GrammarConf.MatchMoreThanZeroTimesSymbol[0],
		},
	}
}
func (nb *NFABuilder) BuildNFABySpecialChars(specialChars []byte) *NFA {
	nfas := make([]*NFA, 0)
	for _, specialChar := range specialChars {
		specialCharNFA := nb.getSpecialCharNFA(specialChar)
		specialCharNFA.MarkSpecialChar(specialChar)
		nfas = append(nfas, specialCharNFA)
	}

	return combineNFAsUsingParallel(nfas)
}
func (nb *NFABuilder) BuildNFAByWord(word string) *NFA {
	return nb.getWordNFA(word)
}
func (nb *NFABuilder) BuildNFAByRegexp(regexp *char.Regexp) *NFA {
	return nb.getRegexpNFA(regexp)
}

func (nb *NFABuilder) getWordNFA(word string) *NFA {
	readingPosition := 0
	nfas := make([]*NFA, 0)
	for readingPosition != len(word) {
		readingChar := word[readingPosition]
		nextChar := byte(0)
		nextPosition := readingPosition + 1
		if nextPosition != len(word) {
			nextChar = word[nextPosition]
		}
		if nb.isAdditionalChar(nextChar) {
			nfas = append(nfas, nb.getAdditionCharNFA(readingChar, nextChar))
			readingPosition += 2
		} else {
			nfas = append(nfas, nb.getCharNFA(readingChar))
			readingPosition += 1
		}
	}
	return combineNFAsUsingSeries(nfas)
}
func (nb *NFABuilder) getAdditionCharNFA(char, additionalChar byte) *NFA {
	nfa := nb.getCharNFA(char)

	// TODO: 这有错误
	switch additionalChar {
	case conf.GetConf().GrammarConf.MatchMoreThanOnceSymbol[0]:
		nfa.endState.linkByEpsChar(nfa.startState)
	case conf.GetConf().GrammarConf.MatchMoreThanZeroTimesSymbol[0]:
		newEnd := NewState(true)
		nfa.endState.linkByEpsChar(nfa.startState)
		nfa.startState.linkByEpsChar(newEnd)
		nfa.endState = newEnd
	}
	return nfa
}
func (nb *NFABuilder) getCharNFA(char byte) *NFA {
	if nb.specialCharTable.CharIsSpecial(char) {
		return nb.getSpecialCharNFA(char)
	}
	return nb.getOrdinaryCharNFA(char)
}
func (nb *NFABuilder) getOrdinaryCharNFA(ordinaryChar byte) *NFA {
	return NewNFA(ordinaryChar)
}
func (nb *NFABuilder) getSpecialCharNFA(specialChar byte) *NFA {
	regexp := nb.specialCharTable.GetRegexp(specialChar)
	return nb.getRegexpNFA(regexp)
}
func (nb *NFABuilder) getRegexpNFA(regexp *char.Regexp) *NFA {

	nfas := make([]*NFA, 0)

	for _, word := range regexp.GetWords() {
		nfas = append(nfas, nb.getWordNFA(word))
	}
	return combineNFAsUsingParallel(nfas)
}

func (nb *NFABuilder)isAdditionalChar(char byte) bool {
	for _, additionChar := range nb.additionChars {
		if additionChar == char {
			return true
		}
	}
	return false
}
func combineNFAsUsingParallel(nfas []*NFA) *NFA {
	if len(nfas) == 1 {
		return nfas[0]
	}
	startStates := getStartStates(nfas)
	endStates := getEndStates(nfas)
	finalNFA := NewEmptyNFA()
	for _, startState := range startStates {
		finalNFA.startState.linkByEpsChar(startState)
	}
	for _, endState := range endStates {
		endState.isEnd = false
		endState.linkByEpsChar(finalNFA.endState)
	}
	return finalNFA
}
func combineNFAsUsingSeries(nfas []*NFA) *NFA {
	for i := 0; i < len(nfas)-1; i++ {
		nfas[i].endState.isEnd = false
		nfas[i].endState.linkByEpsChar(nfas[i+1].startState)
	}
	nfas[0].endState = nfas[len(nfas)-1].endState
	return nfas[0]
}
func getStartStates(nfas []*NFA) []*State {
	startStates := make([]*State, 0)
	for _, nfa := range nfas {
		startStates = append(startStates, nfa.startState)
	}
	return startStates
}
func getEndStates(nfas []*NFA) []*State {
	endStates := make([]*State, 0)
	for _, nfa := range nfas {
		endStates = append(endStates, nfa.endState)
	}
	return endStates
}
