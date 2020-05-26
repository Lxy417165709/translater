package machine

import (
	"bytes"
	"fmt"
)

func (nfa *NFA) GetWordPairs(text []byte) []*WordPair{
	nfa.text = text
	nfa.parseTextToFinalWordPairs()
	return nfa.finalWordPairs
}
func (nfa *NFA) parseTextToFinalWordPairs()  {
	nfa.getFinalWordPairsInit()
	for nfa.readingIsNotOver() {
		nfa.updateFirstEndState()
		nfa.expandStateQueue()
		if nfa.isReachParseBoundary() {
			nfa.handleParseBoundary()
		} else {
			nfa.handleParseProcess()
		}
	}
}


func (nfa *NFA) getFinalWordPairsInit() {
	nfa.bufferOfChars = bytes.Buffer{}
	nfa.finalWordPairs = make([]*WordPair, 0)
	nfa.stateQueue = make([]*State, 0)
	nfa.stateQueue = append(nfa.stateQueue, nfa.GetStartState())
	nfa.readingPosition = 0
}
func (nfa *NFA) readingIsNotOver() bool {
	return nfa.readingPosition != len(nfa.text)
}
func (nfa *NFA) updateFirstEndState() {
	nfa.preEndState = getFirstEndState(nfa.stateQueue)
}
func (nfa *NFA) expandStateQueue() {
	readingChar := nfa.text[nfa.readingPosition]
	tmpQueue := make([]*State, 0)
	for i := 0; i < len(nfa.stateQueue); i++ {
		if nfa.stateQueue[i].GetNext()[readingChar] != nil {
			tmpQueue = append(tmpQueue, nfa.stateQueue[i].GetNext()[readingChar]...)
		}
	}
	nfa.stateQueue = tmpQueue
}
func (nfa *NFA) isReachParseBoundary() bool {
	return len(nfa.stateQueue) == 0
}
func (nfa *NFA) handleParseProcess() {
	nfa.writeReadingCharToBuffer()
	nfa.readNextOne()
}
func (nfa *NFA) writeReadingCharToBuffer() {
	nfa.bufferOfChars.WriteByte(nfa.text[nfa.readingPosition])
}
func (nfa *NFA) readNextOne() {
	nfa.readingPosition++
}
func (nfa *NFA) handleParseBoundary() {
	if nfa.preEndStateNotExist() {
		if nfa.readingCharIsNotBlank() {
			nfa.promptError()
			return
		}
		nfa.readNextOne()
	} else {
		nfa.generateWordPair()
	}
	nfa.reParse()
}
func (nfa *NFA) promptError() {
	panic(fmt.Sprintf("非法字符: %s, 索引位置为: %d", string(nfa.text[nfa.readingPosition]), nfa.readingPosition))
}
func (nfa *NFA) preEndStateNotExist() bool {
	return nfa.preEndState == nil
}
func (nfa *NFA) readingCharIsNotBlank() bool {
	return !isBlank(nfa.text[nfa.readingPosition])
}
func (nfa *NFA) reParse() {
	nfa.bufferOfChars = bytes.Buffer{}
	nfa.stateQueue = nil
	nfa.stateQueue = append(nfa.stateQueue, nfa.GetStartState())
}
func (nfa *NFA) generateWordPair() {
	specialChar := nfa.preEndState.GetSpecialChar()
	word := nfa.bufferOfChars.String()
	nfa.finalWordPairs = append(nfa.finalWordPairs, &WordPair{
		specialChar,
		word,
	})
}
func getFirstEndState(states []*State) *State {
	for _, state := range states {
		if state.GetIsEnd() {
			return state
		}
	}
	return nil
}
func isBlank(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}




