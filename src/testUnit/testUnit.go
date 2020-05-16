package testUnit

import (
	"fmt"
	"stateMachine"
	"strconv"
	"strings"
)

const (
	testUnitDelimiter = "||"
)

type TestUnit struct {
	regex   string
	pattern string
	isMatch bool
}

func NewTestUnit(regex string, pattern string, isMatch bool) *TestUnit {
	return &TestUnit{regex, pattern, isMatch}
}

func (unit *TestUnit) nfaTest() bool {
	nfaBuilder := stateMachine.NewNFABuilder(unit.regex)
	nfa := nfaBuilder.BuildNFA()
	return nfa.IsMatch(unit.pattern) == unit.isMatch
}

func (unit *TestUnit) dfaTest() bool {
	nfaBuilder := stateMachine.NewNFABuilder(unit.regex)
	dfa := nfaBuilder.BuildDFA()
	if !dfa.IsDFA() {
		panic(fmt.Sprintf("DFA算法有误 %v", *unit))
	}
	return dfa.IsMatch(unit.pattern) == unit.isMatch
}

func (unit *TestUnit) Parse(line string) {
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
