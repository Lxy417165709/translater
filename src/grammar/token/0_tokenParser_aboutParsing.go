package token

import (
	"bytes"
	"fmt"
	"grammar/machine"
)



func (tp *TokenParser) GetTokens(text []byte) []*Token{
	tp.text = text
	tp.parseTextToFinalTokens()
	return tp.finalTokens
}
func (tp *TokenParser) parseTextToFinalTokens()  {
	tp.getFinalTokensInit()
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


func (tp *TokenParser) getFinalTokensInit() {
	tp.bufferOfChars = bytes.Buffer{}
	tp.finalTokens = make([]*Token, 0)
	tp.stateQueue = make([]*machine.State, 0)
	tp.stateQueue = append(tp.stateQueue, tp.finalNFA.GetStartState())
	tp.readingPosition = 0
}
func (tp *TokenParser) readingIsNotOver() bool {
	return tp.readingPosition != len(tp.text)
}
func (tp *TokenParser) updateFirstEndState() {
	tp.preEndState = getFirstEndState(tp.stateQueue)
}
func (tp *TokenParser) expandStateQueue() {
	readingChar := tp.text[tp.readingPosition]
	tmpQueue := make([]*machine.State, 0)
	for i := 0; i < len(tp.stateQueue); i++ {
		if tp.stateQueue[i].GetNext()[readingChar] != nil {
			tmpQueue = append(tmpQueue, tp.stateQueue[i].GetNext()[readingChar]...)
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
	if tp.preEndStateNotExist() {
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
func (tp *TokenParser) preEndStateNotExist() bool {
	return tp.preEndState == nil
}
func (tp *TokenParser) readingCharIsNotBlank() bool {
	return !isBlank(tp.text[tp.readingPosition])
}
func (tp *TokenParser) reParse() {
	tp.bufferOfChars = bytes.Buffer{}
	tp.stateQueue = nil
	tp.stateQueue = append(tp.stateQueue, tp.finalNFA.GetStartState())
}
func (tp *TokenParser) generateToken() {
	specialChar := tp.preEndState.GetSpecialChar()
	word := tp.bufferOfChars.String()
	_type := tp.specialCharTable.GetType(specialChar)
	kindCode := tp.specialCharTable.GetCode(specialChar, word)
	tp.finalTokens = append(tp.finalTokens, &Token{
		specialChar,
		kindCode,
		_type,
		word,
	})
}
func getFirstEndState(states []*machine.State) *machine.State {
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
