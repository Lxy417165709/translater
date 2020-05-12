package lexical

import (
	"fmt"
)

type stateType int

const (
	blanking stateType = iota
	readingDigit
	readingDigitOrLetter
	readingOperatorChar
)

type stateMachine struct {
	operatorsTrie  *trie
	segmentManager *segmentManager

	currentTrie  *trie
	currentState stateType
	handleBytes  []byte
	buffer       []byte
	segments     []*pair
}

func NewStateMachine(blankChars, delimiters []byte, reservedWords, operators []string) *stateMachine {
	operatorsTrie := NewTrie()
	for _, operator := range operators {
		operatorsTrie.Insert(operator)
	}
	segmentManager := NewSegmentManager(blankChars, delimiters, reservedWords, operators)

	return &stateMachine{
		operatorsTrie,
		segmentManager,
		operatorsTrie,
		blanking,
		nil,
		nil,
		nil,
	}

}

func (sm *stateMachine) SetHandleBytes(handleBytes []byte) {
	sm.handleBytes = handleBytes
}

func (sm *stateMachine) Handle() []*pair {
	for _, char := range sm.handleBytes {
		sm.eatChar(byte(char))
	}
	return sm.segments
}

func (sm *stateMachine) eatChar(ch byte) {
	if ch == '#' || ch == '~' {
		sm.flash()
		return
	}
	if sm.segmentManager.getCharType(ch) == invalid {
		panic(fmt.Sprintf("字符: %c 为非法字符", ch))
	}
	if sm.segmentManager.getCharType(ch) == blank {
		sm.flash()
		return
	}
	if sm.segmentManager.getCharType(ch) == delimiter {
		sm.flash()
		sm.bufferAppendChar(ch)
		sm.flash()
		return
	}

	switch sm.currentState {
	case blanking:
		switch sm.segmentManager.getCharType(ch) {
		case digit:
			sm.bufferAppendChar(ch)
			sm.changeCurrentState(readingDigit)
		case letter:
			sm.bufferAppendChar(ch)
			sm.changeCurrentState(readingDigitOrLetter)
		case operatorChar:
			sm.bufferAppendChar(ch)
			sm.initCurrentTrie()
			sm.currentTrieToNext(ch)
			sm.changeCurrentState(readingOperatorChar)
		}
	case readingDigitOrLetter, readingDigit, readingOperatorChar:
		if sm.charIsInNotSuitCharTypes(ch) {
			sm.flash()
			sm.eatChar(ch)
			return
		}
		if sm.currentState == readingOperatorChar {
			sm.currentTrieToNext(ch)
		}
		sm.bufferAppendChar(ch)
	}
}
func (sm *stateMachine) flash() {
	if !sm.bufferIsEmpty() {
		sm.segments = append(sm.segments, sm.handleBuffer())
	}

	sm.cleanBuffer()
	sm.changeCurrentState(blanking)
	sm.initCurrentTrie()
}

func (sm *stateMachine) handleBuffer() *pair {
	code := sm.getCodeOfBufferBytes()
	if sm.currentState == readingDigit {
		return NewPair("整型常量", code, sm.bufferBytesToNumber())
	}
	if sm.currentState == readingDigitOrLetter {
		if sm.bufferBytesIsReservedWord() {
			return NewPair("保留字", code, string(sm.buffer))
		} else {
			return NewPair("标识符", code, string(sm.buffer))
		}
	}
	if sm.currentState == readingOperatorChar {
		return NewPair("操作符", code, string(sm.buffer))
	}
	if sm.bufferBytesIsDelimiter() {
		return NewPair("界符", code, sm.buffer[0])
	}
	panic(fmt.Sprintf("buffer:%s 存在错误", string(sm.buffer)))

}
func (sm *stateMachine) getCodeOfBufferBytes() int {
	return sm.segmentManager.getCode(sm.buffer)
}
func (sm *stateMachine) cleanBuffer() {
	sm.buffer = sm.buffer[:0]
}
func (sm *stateMachine) bufferIsEmpty() bool {
	return len(sm.buffer) == 0
}
func (sm *stateMachine) bufferAppendChar(ch byte) {
	sm.buffer = append(sm.buffer, ch)
}
func (sm *stateMachine) bufferBytesToNumber() int {
	return bytesToNumber(sm.buffer)
}
func (sm *stateMachine) bufferBytesIsReservedWord() bool {
	return sm.segmentManager.bytesIsReservedWord(sm.buffer)
}
func (sm *stateMachine) bufferBytesIsDelimiter() bool {
	return sm.segmentManager.bytesIsDelimiter(sm.buffer)
}

func (sm *stateMachine) initCurrentTrie() {
	sm.currentTrie = sm.operatorsTrie
}
func (sm *stateMachine) currentTrieToNext(ch byte) {
	if !sm.operatorsTrie.NextTrieIsExist(ch) {
		panic("词法分析错误，存在错误的操作符")
	}
	sm.currentTrie = sm.currentTrie.GetNextTrie(ch)
}

func (sm *stateMachine) changeCurrentState(targetState stateType) {
	sm.currentState = targetState
}

func (sm *stateMachine) charIsInNotSuitCharTypes(ch byte) bool {
	charType := sm.segmentManager.getCharType(ch)
	NotSuitCharTypes := sm.getNotSuitCharTypes()
	for _, ct := range NotSuitCharTypes {
		if ct == charType {
			return true
		}
	}
	return false
}
func (sm *stateMachine) getNotSuitCharTypes() []charType {
	switch sm.currentState {
	case blanking:
		return nil
	case readingDigit:
		return []charType{letter, operatorChar}
	case readingDigitOrLetter:
		return []charType{operatorChar}
	case readingOperatorChar:
		return []charType{letter, digit}
	}
	return []charType{}
}
