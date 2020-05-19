package stateMachine

import "grammar"

func (nfa *NFA) IsMatch(pattern string) bool {
	return nfa.startState.IsMatch(pattern)
}

func (s *State) IsMatch(pattern string) bool {
	// 空匹配
	for _, nextState := range s.getNextStates(grammar.Eps) {
		if nextState.IsMatch(pattern) {
			return true
		}
	}
	if pattern == "" {
		return s.endFlag
	}

	char := pattern[0]

	// 实匹配
	for _, nextState := range s.getNextStates(char) {
		if nextState.IsMatch(pattern[1:]) {
			return true
		}
	}
	return false
}

func (s *State) stateIsLiving(char byte, beJudgeState *State) bool {
	for _, state := range s.getNextStates(char) {
		if state == beJudgeState {
			return true
		}
	}
	return false
}
