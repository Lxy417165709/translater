package machine


func (s *State) MarkSpecialCharFromHere(specialChar byte, hasVisited map[*State]bool) {
	if hasVisited[s] {
		return
	}
	hasVisited[s] = true
	s.specialChar = specialChar

	//对非空白态的子节点进行处理
	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		nextState.MarkSpecialCharFromHere(specialChar, hasVisited)
	}
	return
}
