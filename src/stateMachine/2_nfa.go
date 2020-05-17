package stateMachine

import (
	"fmt"
	"regexpsManager"
)

type NFA struct {
	startState *State
	endState   *State
	regexpsManager *regexpsManager.RegexpsManager
	respondingSpecialChar byte
}

func NewEmptyNFA(regexpsManager *regexpsManager.RegexpsManager) *NFA {
	return &NFA{NewState(false), NewState(true),regexpsManager,eps}
}
func NewNFA(char byte,regexpsManager *regexpsManager.RegexpsManager) *NFA {
	if !regexpsManager.CharIsSpecial(char) {
		nfa :=  NewEmptyNFA(regexpsManager)
		nfa.getStartState().LinkByChar(char,nfa.getEndState())
		return nfa
	}
	regexp := regexpsManager.GetRegexp(char)
	nfa := NewNFABuilder(regexp,regexpsManager).BuildNFA()
	return nfa
}

func (nfa *NFA) SetRespondingSpecialChar(char byte) {
	nfa.respondingSpecialChar = char
}
func (nfa *NFA) GetRespondingSpecialChar() byte{
	return nfa.respondingSpecialChar
}

func (nfa *NFA) EliminateBlankStates() {
	hasVisited := make(map[*State]bool)
	nfa.getStartState().EliminateNextBlankStatesFromHere(hasVisited)
}

func (nfa *NFA) Show() {
	ids := make(map[*State]int)
	lines := new(int)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("是否DFA:", nfa.IsDFA())
	nfa.getStartState().ShowFromHere(0, ids, make(map[*State]bool), lines)
	fmt.Println("总边数:", *lines)
	fmt.Println(ids)
	fmt.Println("-------------------------------------------------------------")
}

func (nfa *NFA) ToBeDFA() {
	// TODO: 这可能有些问题，可能nfa.endState会发生改变
	hasVisited := make(map[*State]bool)
	nfa.getStartState().MultiWayMergeFromHere(hasVisited)
}
func (nfa *NFA) Get(pattern string) []string {
	result := make([]string, 0)
	begin := nfa.startState
	buffer := ""
	for position := 0; position < len(pattern); position++ {
		char := pattern[position]
		if char == '#' {
			break
		}
		// 不匹配
		if len(begin.toNextState[char]) == 0 {
			if begin.endFlag {
				result = append(result, buffer)
			}
			begin = nfa.startState
			buffer = ""
			if len(begin.toNextState[char]) != 0 {
				position--
			}
			continue
		}
		// 成功匹配
		buffer += string(char)
		begin = begin.toNextState[char][0]
	}
	if buffer != "" && begin.endFlag {
		result = append(result, buffer)
	}
	return result
}
func (nfa *NFA) IsMatch(pattern string) bool {
	return nfa.startState.IsMatch(pattern)
}
func (nfa *NFA) IsDFA() bool {
	hasVisited := make(map[*State]bool)
	return nfa.getStartState().CanBeStartOfDFA(hasVisited)
}
func (nfa *NFA) RepeatPlus(char byte) {
	shouldAddNFA := NewNFA(char,nfa.regexpsManager)
	shouldAddNFA.linkEndStateToStartState()
	nfa.AddSeriesNFA(shouldAddNFA)
}
func (nfa *NFA) RepeatZero(char byte) {
	shouldAddNFA := NewNFA(char,nfa.regexpsManager)
	shouldAddNFA.linkEndStateToStartState()

	endStateOfShouldAddNFA := NewState(true)
	shouldAddNFA.linkStartStateTo(endStateOfShouldAddNFA)
	shouldAddNFA.setEndState(endStateOfShouldAddNFA)

	nfa.AddSeriesNFA(shouldAddNFA)
}
func (nfa *NFA) Once(char byte) {
	beAddedNFA := NewNFA(char,nfa.regexpsManager)
	nfa.AddSeriesNFA(beAddedNFA)
}

func (nfa *NFA) AddParallelNFA(beAddedNFA *NFA) {
	beAddedNFA.breakDown()
	nfa.getStartState().Link(beAddedNFA.getStartState())
	beAddedNFA.getEndState().Link(nfa.getEndState())
}
func (nfa *NFA) AddSeriesNFA(beAddedNFA *NFA) {
	nfa.breakDown()
	nfa.getEndState().Link(beAddedNFA.getStartState())
	nfa.setEndState(beAddedNFA.getEndState())
}

func (nfa *NFA) linkEndStateToStartState() {
	nfa.getEndState().Link(nfa.getStartState())
}
func (nfa *NFA) linkStartStateTo(state *State) {
	nfa.getStartState().Link(state)
}

func (nfa *NFA) breakDown() {
	nfa.getEndState().endFlag = false
}
func (nfa *NFA) getStartState() *State {
	return nfa.startState
}
func (nfa *NFA) getEndState() *State {
	return nfa.endState
}
func (nfa *NFA) setStartState(state *State) {
	nfa.startState = state
}
func (nfa *NFA) setEndState(state *State) {
	nfa.endState = state
}

