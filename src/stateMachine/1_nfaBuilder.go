package stateMachine

import (
	"fmt"
	"regexpsManager"
	"strings"
)

const (
	repeatPlusSymbol = '@'
	repeatZeroSymbol = '$'
	RegexSplitString = "|"
	eps = 0
)


// TODO: 重构
type NFABuilder struct {
	buildRegexp     string // 存在 RegexSplitString 的正则表达式，可以分割为多个NFA
	readingRegex    string // 不存在 RegexSplitString 的正则表达式，只能形成一个NFA
	readingPosition int
	finalNFA        *NFA
	respondingSpecialChar byte
	regexpsManager *regexpsManager.RegexpsManager

}

func NewNFABuilder(buildRegexp string,regexpsManager *regexpsManager.RegexpsManager) *NFABuilder {
	specialChar := byte(0)
	if len(buildRegexp)!=0{
		specialChar = buildRegexp[0]
	}
	return &NFABuilder{
		buildRegexp, // 为了方便越界判断
		buildRegexp,
		0,
		NewEmptyNFA(regexpsManager),
		specialChar,
		regexpsManager,
	}
}
func (nb *NFABuilder) BuildNotBlankStateNFA() *NFA{
	nfa := nb.BuildNFA()
	nfa.EliminateBlankStates()
	return nfa
}
func (nb *NFABuilder) BuildNFA() *NFA {
	regexps := strings.Split(nb.buildRegexp, RegexSplitString)
	if len(regexps) == 0 {
		return nil
	}
	if nb.buildRegexpIsRespondToSingleNFA() {
		nb.finalNFA.startState.LinkByChar(eps, nb.finalNFA.getEndState())
		nb.setReadingRegexp(regexps[0])
		for !nb.readingIsOver() {
			nb.parseChar()
		}

		// 标记
		nb.finalNFA.SetRespondingSpecialChar(nb.respondingSpecialChar)
		nb.finalNFA.MarkDown()

		return nb.finalNFA
	}
	// 这要去除空格（这职责应该不是由它担任）
	for i := 0; i < len(regexps); i++ {
		addedNfa := NewNFABuilder(strings.TrimSpace(regexps[i]),nb.regexpsManager).BuildNFA()
		nb.finalNFA.AddParallelNFA(addedNfa)
	}
	return nb.finalNFA
}
func (nb *NFABuilder) BuildDFA() *NFA {
	nfa := nb.BuildNFA()
	nfa.EliminateBlankStates()
	nfa.ToBeDFA()
	if !nfa.IsDFA() {
		panic(fmt.Sprintf("DFA算法有误"))
	}
	return nfa
}

func (nb *NFABuilder) parseChar() {
	baseChar := nb.getBaseChar()
	nextChar := nb.getNextChar()
	switch {
	case nextChar == repeatPlusSymbol:
		nb.finalNFA.RepeatPlus(baseChar)
		nb.readingPositionMoveTwice()
	case nextChar == repeatZeroSymbol:
		nb.finalNFA.RepeatZero(baseChar)
		nb.readingPositionMoveTwice()
	default:
		nb.finalNFA.Once(baseChar)
		nb.readingPositionMoveOnce()
	}
}
func (nb *NFABuilder) getBaseChar() byte {
	return nb.readingRegex[nb.readingPosition]
}
func (nb *NFABuilder) getNextChar() byte {
	if len(nb.readingRegex) <= nb.readingPosition+1 {
		return eps
	}
	return nb.readingRegex[nb.readingPosition+1]
}
func (nb *NFABuilder) buildRegexpIsRespondToSingleNFA() bool {
	regexps := strings.Split(nb.buildRegexp, RegexSplitString)
	return len(regexps) == 1
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

