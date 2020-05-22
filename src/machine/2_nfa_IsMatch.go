package machine


func (nfa *NFA) IsMatch(pattern string) bool {
	return nfa.startState.IsMatchFromHere(pattern)
}
