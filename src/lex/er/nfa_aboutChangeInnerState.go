package er

func (nfa *NFA) GetInner() []*State{
	return nfa.innerStates
}
func (nfa *NFA) ResetInnerStates() {
	nfa.innerStates = []*State{nfa.startState}
}

// 这里冗余了
func (nfa *NFA) ChangeInnerStates(char byte) {
	nextInnerStates := make([]*State, 0)
	for _, innerState := range nfa.innerStates {
		nextInnerStates = append(nextInnerStates, innerState.getNextStates(char)...)
	}
	nfa.innerStates = nextInnerStates
}

// 这里冗余了
func (nfa *NFA) CanChangeInnerStates(char byte) bool {
	nextInnerStates := make([]*State, 0)
	for _, innerState := range nfa.innerStates {
		nextInnerStates = append(nextInnerStates, innerState.getNextStates(char)...)
	}
	return len(nextInnerStates) != 0
}

func (nfa *NFA) IsEnd() bool {
	for _, innerState := range nfa.innerStates {
		if innerState.isEnd == true {
			return true
		}
	}
	return false
}
