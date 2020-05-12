package lexicalTest

import "fmt"

// TODO: 重构
type State struct {
	endFlag     bool
	toNextState map[byte][]*State
}

func NewState(endFlag bool) *State {
	return &State{endFlag, make(map[byte][]*State)}
}
func (s *State) Show(id int, stateToId map[*State]int, stateIsVisit map[*State]bool) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	stateToId[currentState] = id
	for bytes, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			id++
			nextState.Show(id, stateToId, stateIsVisit)
			option := string(bytes)
			fmt.Printf("id:%d%s --%s--> id:%d%s\n",
				stateToId[currentState],
				currentState.getEndMark(),
				option,
				stateToId[nextState],
				nextState.getEndMark(),
			)
		}
	}
}

func (s *State) addNextState(ch byte, nextState *State) {
	if !s.isNeedToAdd(ch,nextState){
		return
	}
	s.toNextState[ch] = append(s.toNextState[ch], nextState)
}

func (s *State) isNeedToAdd(ch byte,addedState *State) bool{
	for _,nextState := range s.toNextState[ch]{
		if nextState==addedState{
			return false
		}
	}
	return true
}

func (s *State) roadIsExist(ch byte) bool {
	return s.toNextState[ch] != nil
}

func (s *State) IsMatch(pattern string) bool {
	// 空匹配
	nextStates := s.toNextState[eps]
	for _, nextState := range nextStates {
		if nextState.IsMatch(pattern) {
			return true
		}
	}
	if pattern == "" {
		return s.endFlag
	}

	ch := pattern[0]

	// 实匹配
	nextStates = s.toNextState[ch]
	for _, nextState := range nextStates {
		if nextState.IsMatch(pattern[1:]) {
			return true
		}
	}

	return false
}

func (s *State) getEndMark() string {
	if s.endFlag == true {
		return "(OK)"
	}
	return "    "
}
