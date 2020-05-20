package stateMachine

import (
	"conf"
	"grammar"
	"strings"
)

type NFABuilder struct {
	buildRegexp     string

	readingRegex    string
	readingPosition int
	finalNFA        *NFA
}

func NewNFABuilder(buildRegexp string) *NFABuilder {
	nfa := NewNFA()
	specialChar := grammar.GetRegexpsManager().GetSpecialCharFormRegexp(buildRegexp)
	nfa.SetSpecialChar(specialChar)

	return &NFABuilder{
		buildRegexp, // 为了方便越界判断
		"",
		0,
		nfa,
	}
}
func (nb *NFABuilder) BuildNotBlankStateNFA() *NFA {
	return nb.BuildNFA().EliminateBlankStates()
}
func (nb *NFABuilder) BuildNFA() *NFA {
	if nb.buildRegexpIsRespondToSingleNFA() {
		return nb.buildFinalNFAByParsingSingleBuildRegexp()
	}
	return nb.buildFinalNFAByParsingNotSingleBuildRegexp()
}
func (nb *NFABuilder) buildRegexpIsRespondToSingleNFA() bool {
	regexps := strings.Split(nb.buildRegexp, grammar.GetRegexpsManager().GetRegexpDelimiter())
	return len(regexps) == 1
}
func (nb *NFABuilder) buildFinalNFAByParsingNotSingleBuildRegexp() *NFA {
	regexps := strings.Split(nb.buildRegexp, grammar.GetRegexpsManager().GetRegexpDelimiter())
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
	// 这里要注意，依赖到了conf了，可以进行修改
	case nextChar == conf.GetConf().GrammarConf.MatchMoreThanOnceSymbol[0]:
		nb.finalNFA.MatchMoreThanOnce(beAddedNFA)
		nb.readingPositionMoveTwice()
	case nextChar == conf.GetConf().GrammarConf.MatchMoreThanZeroTimesSymbol[0]:
		nb.finalNFA.MatchMoreThanZeroTimes(beAddedNFA)
		nb.readingPositionMoveTwice()
	default:
		nb.finalNFA.MatchOnce(beAddedNFA)
		nb.readingPositionMoveOnce()
	}
}
func (nb *NFABuilder) getBaseChar() byte {
	return nb.readingRegex[nb.readingPosition]
}
func (nb *NFABuilder) getNextChar() byte {
	if len(nb.readingRegex) <= nb.readingPosition+1 {
		return grammar.Eps
	}
	return nb.readingRegex[nb.readingPosition+1]
}
func (nb *NFABuilder) charToNFA(char byte) *NFA {
	if !grammar.GetRegexpsManager().CharIsSpecial(char) {
		nfa := NewNFA()
		nfa.SetSpecialChar(char)
		return nfa.linkStartStateToEndStateByChar(char)
	}
	regexp := grammar.GetRegexpsManager().GetRegexp(char)
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

