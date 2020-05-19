package stateMachine

import (
	"regexpsManager"
	"strings"
)

// TODO: 重构
type NFABuilder struct {
	buildRegexp     string // 存在 RegexSplitString 的正则表达式，可以分割为多个NFA
	readingRegex    string // 不存在 RegexSplitString 的正则表达式，只能形成一个NFA
	readingPosition int
	finalNFA        *NFA
}

func NewNFABuilder(buildRegexp string) *NFABuilder {
	return &NFABuilder{
		buildRegexp, // 为了方便越界判断
		"",
		0,
		NewNFA(regexpsManager.GetRegexpsManager().GetSpecialCharFormRegexp(buildRegexp)),
	}
}
func (nb *NFABuilder) BuildNotBlankStateNFA() *NFA {
	nfa := nb.BuildNFA()
	nfa.EliminateBlankStates()
	return nfa
}
func (nb *NFABuilder) BuildNFA() *NFA {
	if nb.buildRegexpIsRespondToSingleNFA() {
		return nb.buildFinalNFAByParsingSingleBuildRegexp()
	}
	return nb.buildFinalNFAByParsingNotSingleBuildRegexp()
}
func (nb *NFABuilder) buildRegexpIsRespondToSingleNFA() bool {
	regexps := strings.Split(nb.buildRegexp, regexpsManager.RegexSplitString)
	return len(regexps) == 1
}
func (nb *NFABuilder) buildFinalNFAByParsingNotSingleBuildRegexp() *NFA {
	regexps := strings.Split(nb.buildRegexp, regexpsManager.RegexSplitString)
	// 这要去除空格（这职责应该不是由它担任）
	for i := 0; i < len(regexps); i++ {
		addedNfa := NewNFABuilder(strings.TrimSpace(regexps[i])).BuildNFA()
		nb.finalNFA.AddParallelNFA(addedNfa)
	}
	return nb.finalNFA
}
func (nb *NFABuilder) buildFinalNFAByParsingSingleBuildRegexp() *NFA {
	nb.finalNFA.linkStartStateToEndState()
	nb.setReadingRegexp(nb.buildRegexp)
	for !nb.readingIsOver() {
		nb.parseChar()
	}
	return nb.finalNFA
}
func (nb *NFABuilder) parseChar() {
	baseChar := nb.getBaseChar()
	nextChar := nb.getNextChar()
	beAddedNFA := nb.charToNFA(baseChar)
	switch {
	case nextChar == regexpsManager.RepeatPlusSymbol:
		nb.finalNFA.RepeatPlus(beAddedNFA)
		nb.readingPositionMoveTwice()
	case nextChar == regexpsManager.RepeatZeroSymbol:
		nb.finalNFA.RepeatZero(beAddedNFA)
		nb.readingPositionMoveTwice()
	default:
		nb.finalNFA.Once(beAddedNFA)
		nb.readingPositionMoveOnce()
	}
}
func (nb *NFABuilder) getBaseChar() byte {
	return nb.readingRegex[nb.readingPosition]
}
func (nb *NFABuilder) getNextChar() byte {
	if len(nb.readingRegex) <= nb.readingPosition+1 {
		return regexpsManager.Eps
	}
	return nb.readingRegex[nb.readingPosition+1]
}
func (nb *NFABuilder) charToNFA(char byte) *NFA {
	if !regexpsManager.GetRegexpsManager().CharIsSpecial(char) {
		nfa := NewNFA(char)
		return nfa.linkStartStateToEndStateByChar(char)
	}
	regexp := regexpsManager.GetRegexpsManager().GetRegexp(char)
	return NewNFABuilder(regexp).BuildNFA()
}
func (nb *NFABuilder) readingPositionMoveOnce() {
	nb.readingPosition++
}
func (nb *NFABuilder) readingPositionMoveTwice() {
	nb.readingPosition += 2
}
func (nb *NFABuilder) readingIsOver() bool {
	return nb.readingPosition == len(nb.readingRegex)
}
func (nb *NFABuilder) setReadingRegexp(regexp string) {
	nb.readingRegex = regexp
}

