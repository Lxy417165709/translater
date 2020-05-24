package machine

import (
	"bytes"
	"fmt"
)

func (tp *TokenParser) SetText(textNeedToHandle []byte) {
	tp.text = textNeedToHandle
}
func (tp *TokenParser) ParseTextToTokens()  {
	tp.getTokenInit()
	for tp.readingIsNotOver() {
		tp.updateFirstEndState()
		tp.expandStateQueue()
		if tp.isReachParseBoundary() {
			tp.handleParseBoundary()
		} else {
			tp.handleParseProcess()
		}
	}
}
func (tp *TokenParser) GetTokens() []*Token{
	return tp.tokens
}

func (tp *TokenParser) getTokenInit() {
	tp.text = append(tp.text, endChar)
	tp.bufferOfChars = bytes.Buffer{}
	tp.tokens = make([]*Token, 0)
	tp.stateQueue = make([]*state, 0)
	tp.stateQueue = append(tp.stateQueue, tp.finalNFA.startState)
	tp.readingPosition = 0
}
func (tp *TokenParser) readingIsNotOver() bool {
	return tp.text[tp.readingPosition] != endChar
}
func (tp *TokenParser) updateFirstEndState() {
	tp.preEndState = getFirstEndState(tp.stateQueue)
}
func (tp *TokenParser) expandStateQueue() {
	readingChar := tp.text[tp.readingPosition]
	tmpQueue := make([]*state, 0)
	for i := 0; i < len(tp.stateQueue); i++ {
		if tp.stateQueue[i].next[readingChar] != nil {
			tmpQueue = append(tmpQueue, tp.stateQueue[i].next[readingChar]...)
		}
	}
	tp.stateQueue = tmpQueue
}
func (tp *TokenParser) isReachParseBoundary() bool {
	return len(tp.stateQueue) == 0
}
func (tp *TokenParser) handleParseProcess() {
	tp.writeReadingCharToBuffer()
	tp.readNextOne()
}
func (tp *TokenParser) writeReadingCharToBuffer() {
	tp.bufferOfChars.WriteByte(tp.text[tp.readingPosition])
}
func (tp *TokenParser) readNextOne() {
	tp.readingPosition++
}
func (tp *TokenParser) handleParseBoundary() {
	if tp.preEndStateIsNil() {
		if tp.readingCharIsNotBlank() {
			tp.promptError()
			return
		}
		tp.readNextOne()
	} else {
		tp.generateToken()
	}
	tp.reParse()
}
func (tp *TokenParser) promptError() {
	panic(fmt.Sprintf("非法字符: %s, 索引位置为: %d", string(tp.text[tp.readingPosition]), tp.readingPosition))
}
func (tp *TokenParser) preEndStateIsNil() bool {
	return tp.preEndState == nil
}
func (tp *TokenParser) readingCharIsNotBlank() bool {
	return !isBlank(tp.text[tp.readingPosition])
}
func (tp *TokenParser) reParse() {
	tp.bufferOfChars = bytes.Buffer{}
	tp.stateQueue = nil
	tp.stateQueue = append(tp.stateQueue, tp.finalNFA.startState)
}
func (tp *TokenParser) generateToken() {
	specialChar := tp.preEndState.specialChar
	word := tp.bufferOfChars.String()
	_type := tp.specialCharTable.GetType(specialChar)
	kindCode := tp.specialCharTable.GetCode(specialChar, word)
	tp.tokens = append(tp.tokens, &Token{
		specialChar,
		kindCode,
		_type,
		word,
	})
}
func getFirstEndState(states []*state) *state {
	for _, state := range states {
		if state.isEnd {
			return state
		}
	}
	return nil
}
func isBlank(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}
