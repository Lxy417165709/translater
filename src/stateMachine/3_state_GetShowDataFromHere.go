package stateMachine

import (
	"fmt"
)


func (s *State) GetLineOfLinkInformationFromHere(startId int, stateToId map[*State]int, stateIsVisit map[*State]bool) []string{
	if stateIsVisit[s] {
		return []string{}
	}
	result := make([]string,0)
	stateIsVisit[s] = true
	stateToId[s] = startId
	for bytes, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			result = append(result, nextState.GetLineOfLinkInformationFromHere(len(stateToId), stateToId, stateIsVisit)...)
			option := string(bytes)
			result = append(result, formMermaidLine(stateToId,s,option,nextState))
		}
	}
	return result
}
func (s *State) getEndMark(id int) string {
	if s.endFlag {
		return fmt.Sprintf("((%d))", id)
	} else {
		return fmt.Sprintf("(%d)", id)
	}
}




