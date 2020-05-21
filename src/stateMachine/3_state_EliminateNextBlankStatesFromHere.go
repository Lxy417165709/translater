package stateMachine

import (
	"grammar"
)




func (s *State) EliminateNextBlankStatesFromHere(hasVisited map[*State]bool) {
	if hasVisited[s] {
		return
	}
	hasVisited[s] = true
	// 消除空白态
	for s.haveBlankStates() {
		mapOfReachableStateOfBlankStates := s.formMapOfReachableStateOfBlankStates()
		if s.isNextBlankStatesHaveEndState() {
			s.setEndFlag(true)
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
	return s.getNextStates(grammar.Eps)
}
func (s *State) isNextBlankStatesHaveEndState() bool {
	return s.isNextStatesHaveEndState(grammar.Eps)
}
func (s *State) isNextStatesHaveEndState(char byte) bool {
	for _, state := range s.toNextState[char] {
		if state.endFlag == true {
			return true
		}
	}
	return false
}
func (s *State) setEndFlag(value bool) {
	s.endFlag = value
}
func (s *State) cleanBlankStates() {
	s.cleanNextStates(grammar.Eps)
}
func (s *State) cleanNextStates(char byte) {
	delete(s.toNextState, char)
}
func (s *State) addNextStates(addedMap map[byte][]*State) {
	for char, states := range addedMap {
		s.toNextState[char] = append(s.toNextState[char], states...)
	}
}

func getStatesToNext(states []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	hasExist := make(map[byte]map[*State]bool)
	for _, state := range states {
		for char, nextStates := range state.toNextState {
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

