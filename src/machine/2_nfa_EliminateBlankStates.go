package machine





func (nfa *NFA) EliminateBlankStates() *NFA{
	hasVisited := make(map[*state]bool)
	nfa.startState.EliminateNextBlankStatesFromHere(hasVisited)
	return nfa
}
