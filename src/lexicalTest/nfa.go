package lexicalTest

import "fmt"

type NFA struct {
	startState *State
	endState   *State
	isDFA      bool
}
func (nfa *NFA) IsDFA() bool{
	return nfa.isDFA
}
// 重构

// 空NFA能匹配空字符串
func NewNFA() *NFA {
	startState := NewState(false)
	endState := NewState(true)
	startState.Link(endState)
	return &NFA{startState, endState, false}
}

// 创建字符对应的状态机
func NewCharNFA(char byte) *NFA {
	startState := NewState(false)
	endState := NewState(true)
	startState.LinkByChar(char, endState)
	return &NFA{startState, endState, false}
}

func NewTestNFA() *NFA {
	nextState0 := NewState(false)
	nextState1 := NewState(false)
	nextState2 := NewState(false)
	nextState3 := NewState(false)
	nextState4 := NewState(false)
	nextState5 := NewState(false)
	nextState6 := NewState(false)
	nextState7 := NewState(false)
	nextState8 := NewState(false)
	nextState10 := NewState(false)

	//nextState3.LinkByChar('d', nextState4)
	//
	//nextState4.LinkByChar('f', startState)
	//

	nextState0.LinkByChar(eps, nextState1)
	nextState1.LinkByChar(eps, nextState2)
	nextState2.LinkByChar(eps, nextState3)
	nextState3.LinkByChar(eps, nextState4)
	nextState4.LinkByChar(eps, nextState5)
	nextState5.LinkByChar(eps, nextState6)
	nextState6.LinkByChar('a', nextState7)
	nextState7.LinkByChar(eps, nextState8)
	nextState8.LinkByChar(eps, nextState10)
	nextState8.LinkByChar(eps, nextState3)
	return &NFA{nextState0, nil, false}
}

func (nfa *NFA) IsMatch(pattern string) bool {
	if nfa.isDFA {
		return nfa.startState.IsMatchByDFA(pattern)
	}
	return nfa.startState.IsMatch(pattern)
}
func (nfa *NFA) Show() {
	ids := make(map[*State]int)
	nfa.startState.Show(0, ids, make(map[*State]bool))
	fmt.Println(ids)
}

// TODO: 可以统一
func (nfa *NFA) RepeatPlus(char byte) {
	var beAddedNFA *NFA
	if !GlobalNFAManager.IdentityIsExist(char) {
		beAddedNFA = NewCharNFA(char)
	} else {
		beAddedNFA = NewNFABuilder(GlobalNFAManager.Get(char)).BuildDFA()
	}

	beAddedNFA.getEndState().Link(beAddedNFA.getStartState())
	nfa.AddSeriesNFA(beAddedNFA)
}
func (nfa *NFA) RepeatZero(char byte) {
	var beAddedNFA *NFA
	if !GlobalNFAManager.IdentityIsExist(char) {
		beAddedNFA = NewCharNFA(char)
	} else {
		beAddedNFA = NewNFABuilder(GlobalNFAManager.Get(char)).BuildDFA()
	}

	// 次序不能倒
	beAddedNFA.getEndState().Link(beAddedNFA.getStartState())

	// 次序不能倒
	beAddedNFAEndState := NewState(true)
	beAddedNFA.startState.Link(beAddedNFAEndState)
	beAddedNFA.setEndState(beAddedNFAEndState)

	nfa.AddSeriesNFA(beAddedNFA)
}
func (nfa *NFA) Once(char byte) {
	var beAddedNFA *NFA
	if !GlobalNFAManager.IdentityIsExist(char) {
		beAddedNFA = NewCharNFA(char)
	} else {
		beAddedNFA = NewNFABuilder(GlobalNFAManager.Get(char)).BuildDFA()
		//fmt.Printf("%p\n",GlobalNFAManager.Get(char))
	}
	nfa.AddSeriesNFA(beAddedNFA)
}

func (nfa *NFA) AddParallelNFA(beAddedNFA *NFA) {
	nfa.startState.Link(beAddedNFA.startState)
	beAddedNFA.breakDown()
	beAddedNFA.endState.Link(nfa.endState)
}
func (nfa *NFA) AddSeriesNFA(beAddedNFA *NFA) {
	nfa.breakDown()
	nfa.endState.Link(beAddedNFA.startState)
	nfa.endState = beAddedNFA.endState
}

func (nfa *NFA) EndTo(state *State) {
	nfa.endState.Link(state)
	//nfa.endState = state
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

func (nfa *NFA) breakDown() {
	nfa.isDFA = false
	nfa.endState.endFlag = false
}
func (nfa *NFA) open() {
	nfa.endState.endFlag = true
}

func (nfa *NFA) ChangeToDFA() *NFA {
	nfa.startState.DFA(make(map[*State]bool))
	return nfa
}
