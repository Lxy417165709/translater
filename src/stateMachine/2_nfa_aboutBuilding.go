package stateMachine

func (nfa *NFA) SetSpecialChar(char byte) {
	nfa.specialChar = char
}
func (nfa *NFA) GetSpecialChar() byte {
	return nfa.specialChar
}

func (nfa *NFA) EliminateBlankStates() *NFA{
	hasVisited := make(map[*State]bool)
	nfa.getStartState().EliminateNextBlankStatesFromHere(hasVisited)
	return nfa
}

func (nfa *NFA) MatchMoreThanOnce(beAddedNFA *NFA) *NFA{
	nfa.AddSeriesNFA(beAddedNFA.linkEndStateToStartState())
	return nfa
}
func (nfa *NFA) MatchMoreThanZeroTimes(beAddedNFA *NFA) *NFA{
	endStateOfShouldAddNFA := NewState(true)
	beAddedNFA.linkEndStateToStartState().linkStartStateTo(endStateOfShouldAddNFA).setEndState(endStateOfShouldAddNFA)
	return nfa.AddSeriesNFA(beAddedNFA)
}
func (nfa *NFA) MatchOnce(beAddedNFA *NFA)*NFA {
	return nfa.AddSeriesNFA(beAddedNFA)
}
func (nfa *NFA) AddParallelNFA(beAddedNFA *NFA) *NFA{
	beAddedNFA.breakDown()
	beAddedNFA.linkEndStateTo(nfa.getEndState())
	return nfa.linkStartStateTo(beAddedNFA.getStartState())
}
func (nfa *NFA) AddSeriesNFA(beAddedNFA *NFA) *NFA{
	nfa.breakDown()
	nfa.linkEndStateTo(beAddedNFA.getStartState()).setEndState(beAddedNFA.getEndState())
	return nfa
}
func (nfa *NFA) breakDown() {
	nfa.getEndState().endFlag = false
}

func (nfa *NFA) MarkSpecialChar() *NFA{
	nfa.startState.MarkSpecialChar(nfa.specialChar,make(map[*State]bool))
	return nfa
}

func (nfa *NFA) linkStartStateToEndState () *NFA{
	return nfa.linkStartStateTo(nfa.endState)
}
func (nfa *NFA) linkStartStateToEndStateByChar (char byte) *NFA{
	nfa.getStartState().LinkByChar(char,nfa.getEndState())
	return nfa
}
func (nfa *NFA) linkEndStateToStartState() *NFA{
	nfa.endState.Link(nfa.startState)
	return nfa
}
func (nfa *NFA) linkStartStateTo(state *State) *NFA{
	nfa.getStartState().Link(state)
	return nfa
}
func (nfa *NFA) linkEndStateTo(state *State) *NFA{
	nfa.getEndState().Link(state)
	return nfa
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


func getNextStates(states []*State, readingChar byte) []*State {
	tmpQueue := make([]*State, 0)
	for i := 0; i < len(states); i++ {
		if states[i].toNextState[readingChar] != nil {
			tmpQueue = append(tmpQueue, states[i].toNextState[readingChar]...)
		}
	}
	return tmpQueue
}
