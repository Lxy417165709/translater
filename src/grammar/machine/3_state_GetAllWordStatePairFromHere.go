package machine

func (s *state) GetAllWordPairsFromHere(nowStr string,hasVisited map[*state]bool) []*wordPair{
	if hasVisited[s]==true{
		return  []*wordPair{}
	}
	hasVisited[s]=true
	result := make([]*wordPair,0)
	if s.isEnd{
		result = append(result,&wordPair{
			s.specialChar,
			nowStr,
		})
	}

	for char,nextStates := range s.next {
		nextStr := nowStr+string(char)
		for _,nextState := range nextStates{
			result = append(result,nextState.GetAllWordPairsFromHere(nextStr,hasVisited)...)
		}
	}
	return result
}

