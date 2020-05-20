package stateMachine

import (
	"grammar"
)

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
