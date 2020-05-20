package stateMachine

import (
	"fmt"
)

const endSymbol = '#'


func (nfa *NFA) SetPattern(pattern []byte) {
	nfa.parsedPattern = pattern
}
func (nfa *NFA) ParsePattern() []*Token{
	nfa.parseInit()
	for nfa.readingIsNotOver() {
		nfa.updateFirstEndState()
		nfa.expandStateQueue()
		if nfa.isReachParseBoundary() {
			nfa.handleParseBoundary()
		}else{
			nfa.handleParseProcess()
		}
	}
	return nfa.tokens
}
func (nfa *NFA) parseInit() {
	nfa.parsedPattern = append(nfa.parsedPattern, endSymbol)
	nfa.buffer,nfa.tokens = "",make([]*Token,0)
	nfa.stateQueue = append(nfa.stateQueue,nfa.startState)
	nfa.readingPosition = 0
}
func (nfa *NFA) readingIsNotOver() bool{
	return nfa.parsedPattern[nfa.readingPosition] != endSymbol
}
func (nfa *NFA) updateFirstEndState() {
	nfa.firstEndState = getFirstEndState(nfa.stateQueue)
}
func (nfa *NFA) expandStateQueue() {
	nfa.stateQueue = getNextStates(nfa.stateQueue, nfa.parsedPattern[nfa.readingPosition])
}

func (nfa *NFA) isReachParseBoundary() bool {
	return len(nfa.stateQueue) == 0
}
func (nfa *NFA) handleParseProcess() {
	nfa.writeReadingCharToBuffer()
	nfa.readNextOne()
}
func (nfa *NFA) writeReadingCharToBuffer() {
	nfa.buffer += string(nfa.parsedPattern[nfa.readingPosition])
}
func (nfa *NFA) readNextOne() {
	nfa.readingPosition++
}

func (nfa *NFA) handleParseBoundary() {
	if nfa.firstEndStateIsNil(){
		if nfa.readingCharIsNotBlank(){
			nfa.promptError()
			return
		}
		nfa.readNextOne()
	}else{
		nfa.generateToken()
	}
	nfa.reParse()
}

func (nfa *NFA) promptError() {
	panic(fmt.Sprintf("非法字符: %s, 索引位置为: %d",string(nfa.parsedPattern[nfa.readingPosition]),nfa.readingPosition))
}

func (nfa *NFA) firstEndStateIsNil() bool{
	return nfa.firstEndState==nil
}
func (nfa *NFA) readingCharIsNotBlank() bool{
	return !isBlank(nfa.parsedPattern[nfa.readingPosition])
}

func (nfa *NFA) reParse() {
	nfa.buffer = ""
	nfa.stateQueue = nil
	nfa.stateQueue = append(nfa.stateQueue, nfa.startState)
}


func (nfa *NFA) readingCharIsInvalid() bool{
	return nfa.firstEndState==nil && !isBlank(nfa.parsedPattern[nfa.readingPosition])
}

func (nfa *NFA) generateToken() {
	specialChar := nfa.firstEndState.belongToSpecialChar
	word := nfa.buffer
	nfa.tokens = append(nfa.tokens, NewToken(specialChar,word))
}

