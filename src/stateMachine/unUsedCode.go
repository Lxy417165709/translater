package stateMachine

// TODO: 用于标记该状态属于哪个自动机
func (s *State) MarkDown(specialChar byte, stateIsVisit map[*State]bool) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	s.markFlag = specialChar
	for _, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			nextState.MarkDown(specialChar, stateIsVisit)
		}
	}
}
