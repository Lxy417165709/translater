package er

func (s *State) EliminateNextBlankStatesFromHere(hasVisited map[*State]bool) {
	if hasVisited[s] {
		return
	}
	hasVisited[s] = true
	// 消除空白态
	for s.haveStates(blankChar) {
		mapOfReachableStateOfBlankStates := s.formMapOfReachableStateOfBlankStates()
		if s.isNextStatesHaveEndState(blankChar) {
			s.isEnd = true
		}
		s.cleanNextStates(blankChar)
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
	blankStates := s.getNextStates(blankChar)
	return getStatesToNext(blankStates)
}
func (s *State) haveStates(char byte) bool {
	return len(s.getNextStates(char)) != 0
}
func (s *State) isNextStatesHaveEndState(char byte) bool {
	for _, State := range s.next[char] {
		if State.isEnd == true {
			return true
		}
	}
	return false
}
func (s *State) cleanNextStates(char byte) {
	delete(s.next, char)
}
func (s *State) addNextStates(addedMap map[byte][]*State) {
	for char, States := range addedMap {
		s.next[char] = append(s.next[char], States...)
	}
}

func (s *State) getNextStates(char byte) []*State {
	return s.next[char]
}
func (s *State) getAllNextStates() []*State {
	result := make([]*State, 0)
	for char := range s.next {
		result = append(result, s.getNextStates(char)...)
	}
	return result
}

func getStatesToNext(States []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	hasExist := make(map[byte]map[*State]bool)
	for _, stat := range States {
		for char, nextStates := range stat.next {
			for _, nextState := range nextStates {
				if hasExist[char] == nil {
					hasExist[char] = make(map[*State]bool)
				}
				if hasExist[char][nextState] {
					continue
				}
				hasExist[char][nextState] = true
				result[char] = append(result[char], nextState)
			}
		}
	}
	return result
}
