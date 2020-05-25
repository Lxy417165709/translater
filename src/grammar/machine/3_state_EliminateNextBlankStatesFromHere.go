package machine



func (s *State) EliminateNextBlankStatesFromHere(hasVisited map[*State]bool) {
	if hasVisited[s] {
		return
	}
	hasVisited[s] = true
	// 消除空白态
	for s.haveBlankStates() {
		mapOfReachableStateOfBlankStates := s.formMapOfReachableStateOfBlankStates()
		if s.isNextBlankStatesHaveEndState() {
			s.isEnd=true
		}
		s.cleanBlankStates()
		s.addNextStates(mapOfReachableStateOfBlankStates)
	}

	//对非空白态的子节点进行处理
	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		nextState.EliminateNextBlankStatesFromHere(hasVisited)
	}
	return
}

func (s *State) formMapOfReachableStateOfBlankStates() map[byte][]*State {
	blankStates := s.getNextBlankStates()
	return getStatesToNext(blankStates)
}
func (s *State) haveBlankStates() bool {
	return len(s.getNextBlankStates()) != 0
}
func (s *State) getNextBlankStates() []*State {
	return s.getNextStates(Eps)
}
func (s *State) isNextBlankStatesHaveEndState() bool {
	return s.isNextStatesHaveEndState(Eps)
}
func (s *State) isNextStatesHaveEndState(char byte) bool {
	for _, State := range s.next[char] {
		if State.isEnd == true {
			return true
		}
	}
	return false
}

func (s *State) cleanBlankStates() {
	s.cleanNextStates(Eps)
}
func (s *State) cleanNextStates(char byte) {
	delete(s.next, char)
}
func (s *State) addNextStates(addedMap map[byte][]*State) {
	for char, States := range addedMap {
		s.next[char] = append(s.next[char], States...)
	}
}
