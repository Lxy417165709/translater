package stateMachine

func (s *State) GetAllTokensFromHere(nowStr string,hasVisited map[*State]bool) []*Token {
	if hasVisited[s]==true{
		return  []*Token{}
	}
	hasVisited[s]=true
	result := make([]*Token,0)
	if s.endFlag{
		result = append(result, NewToken(s.belongToSpecialChar,nowStr))
	}

	for char,nextStates := range s.toNextState {
		nextStr := nowStr+string(char)
		for _,nextState := range nextStates{
			result = append(result,nextState.GetAllTokensFromHere(nextStr,hasVisited)...)
		}
	}
	return result
}

