package machine

import (
	"grammar"
)


const endChar = byte('#')
type NFABuilder struct {
	specialCharTable *grammar.SpecialCharTable
}

func NewNFABuilder(specialCharTable *grammar.SpecialCharTable) *NFABuilder{
	return &NFABuilder{
		specialCharTable:specialCharTable,
	}
}

func (nb *NFABuilder)BuildNFABySpecialChars(specialChars []byte) *NFA{
	nfas := make([]*NFA,0)
	for _,specialChar := range specialChars {
		specialCharNFA := nb.getSpecialCharNFA(specialChar)
		// TODO: specialCharNFA.MarkSpecialChar()  为了最终状态机有specialChar信息，这里要进行标记
		nfas = append(nfas,specialCharNFA)
	}
	return combineNFAsUsingParallel(nfas)
}
func (nb *NFABuilder)BuildNFAByWord(word string) *NFA{
	return nb.getWordNFA(word)
}





func (nb *NFABuilder) getWordNFA(word string) *NFA{
	readingPosition := 0
	nfas := make([]*NFA,0)
	for readingPosition!=len(word){
		readingChar := word[readingPosition]
		nextChar := endChar
		nextPosition := readingPosition+1
		if nextPosition!=len(word){
			nextChar=word[nextPosition]
		}
		if isAdditionChar(nextChar){
			nfas = append(nfas,nb.getAdditionCharNFA(readingChar,nextChar))
			readingPosition+=2
		}else{
			nfas = append(nfas,nb.getCharNFA(readingChar))
			readingPosition+=1
		}
	}
	return combineNFAsUsingSeries(nfas)
}


func (nb *NFABuilder) getAdditionCharNFA(char,additionalChar byte) *NFA{
	nfa := nb.getCharNFA(char)
	addedNFA := nb.getCharNFA(char)
	switch additionalChar {
	case '@':
		addedNFA.endState.isEnd=false
		addedNFA.endState.linkByEpsChar(nfa.startState)
		nfa.endState.linkByEpsChar(addedNFA.startState)
	case '$':
		nfa.endState.isEnd=false
		addedNFA.endState.isEnd=false

		newEnd := NewState(true)
		addedNFA.endState.linkByEpsChar(nfa.startState)
		nfa.endState.linkByEpsChar(addedNFA.startState)
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
func (nb *NFABuilder) getOrdinaryCharNFA(ordinaryChar byte) *NFA{
	return NewNFA(ordinaryChar)
}
func (nb *NFABuilder) getSpecialCharNFA(specialChar byte) *NFA{
	regexp := nb.specialCharTable.GetRegexp(specialChar)
	return nb.getRegexpNFA(regexp)
}
func (nb *NFABuilder)getRegexpNFA(regexp *grammar.Regexp) *NFA{
	nfas := make([]*NFA,0)
	for _,word := range regexp.GetWords(){
		nfas = append(nfas,nb.getWordNFA(word))
	}
	return combineNFAsUsingParallel(nfas)
}




func combineNFAsUsingParallel(nfas []*NFA) *NFA{
	startStates := getStartStates(nfas)
	endStates := getEndStates(nfas)
	finalNFA := NewEmptyNFA()
	for _,startState := range startStates{
		finalNFA.startState.linkByEpsChar(startState)
	}
	for _,endState := range endStates{
		endState.linkByEpsChar(finalNFA.endState)
	}
	return finalNFA
}
func combineNFAsUsingSeries(nfas []*NFA) *NFA{
	for i := 0;i<len(nfas)-1;i++{
		nfas[i].endState.linkByEpsChar(nfas[i+1].startState)
	}
	return nfas[0]
}

func getStartStates(nfas []*NFA) []*state{
	startStates := make([]*state,0)
	for _,nfa := range nfas {
		startStates = append(startStates,nfa.startState)
	}
	return startStates
}

// 假定只有一个终止态，并处于endState
func getEndStates(nfas []*NFA) []*state{
	endStates := make([]*state,0)
	for _,nfa := range nfas {
		endStates = append(endStates,nfa.endState)
	}
	return endStates
}

// TODO: 这个也需要配置
var additionChars = []byte{
	'@','$',
}
func  isAdditionChar(char byte) bool {
	for _,additionChar := range additionChars{
		if additionChar==char{
			return true
		}
	}
	return false
}

