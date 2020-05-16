package testLay

import (
	"fmt"
	"regexpsManager"
	"stateMachine"
	"strconv"
	"strings"
)

const (
	testUnitDelimiter = "||"
)

type NFATestUnit struct {
	regex   string
	pattern string
	isMatch bool
	regexpsManager *regexpsManager.RegexpsManager
}

func NewNFATestUnit(regexpsManager *regexpsManager.RegexpsManager) *NFATestUnit {
	return &NFATestUnit{regexpsManager:regexpsManager}
}
func (unit *NFATestUnit) Test() bool{
	return unit.nfaTest() && unit.nfaTest()
}
func (unit *NFATestUnit) nfaTest() bool {
	nfaBuilder := stateMachine.NewNFABuilder(unit.regex,unit.regexpsManager)
	nfa := nfaBuilder.BuildNFA()
	return nfa.IsMatch(unit.pattern) == unit.isMatch
}
func (unit *NFATestUnit) dfaTest() bool {
	nfaBuilder := stateMachine.NewNFABuilder(unit.regex,unit.regexpsManager)
	dfa := nfaBuilder.BuildDFA()
	if !dfa.IsDFA() {
		panic(fmt.Sprintf("DFA算法有误 %v", *unit))
	}
	return dfa.IsMatch(unit.pattern) == unit.isMatch
}
func (unit *NFATestUnit) Parse(line string) {
	parts := strings.Split(strings.TrimSpace(line), testUnitDelimiter)
	if len(parts) != 3 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于3", parts))
	}
	regex := strings.TrimSpace(parts[0])
	pattern := strings.TrimSpace(parts[1])
	matchFlag, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		panic(err)
	}
	unit.regex = regex
	unit.pattern = pattern
	unit.isMatch = intToBool(matchFlag)
}
func intToBool(a int) bool {
	return a != 0
}

