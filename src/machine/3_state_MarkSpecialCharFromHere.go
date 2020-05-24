package machine


func (s *state) MarkSpecialCharFromHere(specialChar byte, hasVisited map[*state]bool) {
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
