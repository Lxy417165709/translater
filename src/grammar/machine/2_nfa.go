package machine

import "bytes"

type NFA struct {
	startState            *State
	endState              *State

	text            []byte
	preEndState     *State
	stateQueue      []*State
	bufferOfChars   bytes.Buffer
	readingPosition int
	finalWordPairs     []*WordPair
}

func NewNFA(ordinaryChar byte) *NFA {
	startState:=NewState(false)
	endState:=NewState(true)
	startState.next[ordinaryChar] = append(startState.next[ordinaryChar],endState)
	return &NFA{
		startState:startState,
		endState:endState,
	}
}
func NewEmptyNFA() *NFA{
	return &NFA{
		startState:NewState(false),
		endState:NewState(true),
	}
}

func (nfa *NFA)GetStartState() *State{
	return nfa.startState
}
















