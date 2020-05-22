package machine

import (
	"grammar"
)



func (s *state) EliminateNextBlankStatesFromHere(hasVisited map[*state]bool) {
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

func (s *state) formMapOfReachableStateOfBlankStates() map[byte][]*state {
	blankStates := s.getNextBlankStates()
	return getStatesToNext(blankStates)
}
func (s *state) haveBlankStates() bool {
	return len(s.getNextBlankStates()) != 0
}
func (s *state) getNextBlankStates() []*state {
	return s.getNextStates(grammar.Eps)
}
func (s *state) isNextBlankStatesHaveEndState() bool {
	return s.isNextStatesHaveEndState(grammar.Eps)
}
func (s *state) isNextStatesHaveEndState(char byte) bool {
	for _, state := range s.next[char] {
		if state.isEnd == true {
			return true
		}
	}
	return false
}

func (s *state) cleanBlankStates() {
	s.cleanNextStates(grammar.Eps)
}
func (s *state) cleanNextStates(char byte) {
	delete(s.next, char)
}
func (s *state) addNextStates(addedMap map[byte][]*state) {
	for char, states := range addedMap {
		s.next[char] = append(s.next[char], states...)
	}
}
