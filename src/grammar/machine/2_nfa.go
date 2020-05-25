package machine

type NFA struct {
	startState            *State
	endState              *State
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
















