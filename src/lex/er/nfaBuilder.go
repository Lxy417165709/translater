package er

import (
	"env"
	"fmt"
	"lex/tb"
)

type NFABuilder struct {
	specialCharTable *tb.SpecialCharTable
}


func NewNFABuilder(conf *env.NFABuilderConf) (*NFABuilder, error) {
	var err error
	nb := &NFABuilder{}
	if nb.specialCharTable, err = tb.NewSpecialCharTable(conf.SpecialCharTableConf); err != nil {
		return nil, fmt.Errorf("初始化NFA构建器失败，%s", err.Error())
	}
	return nb, nil
}

// 用于软件测试实验
func NewDefaultNFABuilder() (*NFABuilder, error) {
	var err error
	nb := &NFABuilder{}
	if nb.specialCharTable, err = tb.NewDefaultSpecialCharTable(); err != nil {
		return nil, fmt.Errorf("初始化NFA构建器失败，%s", err.Error())
	}
	return nb, nil
}


func (nb *NFABuilder) BuildNFABySpecialChar (specialChar byte) *NFA{
	nfa := nb.BuildNFAByWord(string(specialChar))
	nfa.EliminateBlankStates()
	nfa.SetSpecialChar(specialChar)
	return nfa
}


func (nb *NFABuilder) BuildNFAByWord(word string) *NFA {
	return nb.getWordNFA(word)
}
func (nb *NFABuilder) BuildNFAByRegexp(regexp *tb.Regexp) *NFA {
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
		if nb.specialCharTable.IsAdditionalChar(nextChar) {
			nfas = append(nfas, nb.getAdditionCharNFA(readingChar, nextChar))
			readingPosition += 2
		} else {
			nfas = append(nfas, nb.getCharNFA(readingChar))
			readingPosition += 1
		}
	}
	return nb.combineNFAsUsingSeries(nfas)
}
func (nb *NFABuilder) getAdditionCharNFA(char, additionalChar byte) *NFA {
	nfa := nb.getCharNFA(char)

	switch additionalChar {
	case nb.specialCharTable.GetSpecialCharOfMarchMoreThanOnce():
		nfa.endState.link(blankChar,nfa.startState)
	case nb.specialCharTable.GetSpecialCharOfMarchMoreThanZeroTimes():
		newEnd := NewState(true)
		nfa.endState.link(blankChar,nfa.startState)
		nfa.startState.link(blankChar,newEnd)
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
func (nb *NFABuilder) getRegexpNFA(regexp *tb.Regexp) *NFA {

	nfas := make([]*NFA, 0)

	for _, word := range regexp.Words {
		nfas = append(nfas, nb.getWordNFA(word))
	}
	return nb.combineNFAsUsingParallel(nfas)
}


func  (nb *NFABuilder)combineNFAsUsingParallel(nfas []*NFA) *NFA {
	if len(nfas) == 1 {
		return nfas[0]
	}
	startStates := getStartStates(nfas)
	endStates := getEndStates(nfas)
	finalNFA := NewEmptyNFA()
	for _, startState := range startStates {
		finalNFA.startState.link(blankChar,startState)
	}
	for _, endState := range endStates {
		endState.isEnd = false
		endState.link(blankChar,finalNFA.endState)
	}
	return finalNFA
}
func (nb *NFABuilder)combineNFAsUsingSeries(nfas []*NFA) *NFA {
	for i := 0; i < len(nfas)-1; i++ {
		nfas[i].endState.isEnd = false
		nfas[i].endState.link(blankChar,nfas[i+1].startState)
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


