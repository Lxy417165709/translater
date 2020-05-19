package stateMachine


func (nfa *NFA) IsMatch(pattern string) bool {
	return nfa.startState.IsMatch(pattern)
}
