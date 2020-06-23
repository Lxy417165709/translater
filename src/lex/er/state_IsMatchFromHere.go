package er


func (s *State) IsMatchFromHere(pattern string) bool {
	// 空匹配
	for _, nextState := range s.getNextStates(blankChar) {
		if nextState.IsMatchFromHere(pattern) {
			return true
		}
	}
	if pattern == "" {
		return s.isEnd
	}
	char := pattern[0]

	// 实匹配
	for _, nextState := range s.getNextStates(char) {
		if nextState.IsMatchFromHere(pattern[1:]) {
			return true
		}
	}
	return false
}


