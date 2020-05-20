package stateMachine

import (
	"grammar"
)

type State struct {
	endFlag             bool
	toNextState         map[byte][]*State
	belongToSpecialChar byte
}

func NewState(endFlag bool) *State {
	return &State{endFlag: endFlag, toNextState: make(map[byte][]*State)}
}
func (s *State) Link(nextState *State) {
	s.LinkByChar(grammar.Eps, nextState)
}
func (s *State) LinkByChar(ch byte, nextState *State) {
	s.toNextState[ch] = append(s.toNextState[ch], nextState)
}


func (s *State) MarkSpecialChar(specialChar byte, hasVisited map[*State]bool) {
	if hasVisited[s] {
		return
	}
	hasVisited[s] = true
	s.setSpecialChar(specialChar)

	//对非空白态的子节点进行处理
	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		nextState.MarkSpecialChar(specialChar, hasVisited)
	}
	return
}
func (s *State) setSpecialChar(specialChar byte) {
	s.belongToSpecialChar = specialChar
}
func (s *State) getAllNextStates() []*State {
	result := make([]*State, 0)
	for char := range s.toNextState {
		result = append(result, s.getNextStates(char)...)
	}
	return result
}
func (s *State) getNextStates(char byte) []*State {
	return s.toNextState[char]
}
