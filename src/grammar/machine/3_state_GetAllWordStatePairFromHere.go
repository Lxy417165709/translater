package machine

func (s *State) GetAllWordPairsFromHere(nowStr string, hasVisited map[*State]bool) []*WordPair {
	if hasVisited[s] == true {
		return []*WordPair{}
	}
	hasVisited[s] = true
	result := make([]*WordPair, 0)
	if s.isEnd {
		result = append(result, &WordPair{
			s.specialChar,
			nowStr,
		})
	}

	for char, nextStates := range s.next {
		nextStr := nowStr + string(char)
		for _, nextState := range nextStates {
			result = append(result, nextState.GetAllWordPairsFromHere(nextStr, hasVisited)...)
		}
	}
	return result
}
