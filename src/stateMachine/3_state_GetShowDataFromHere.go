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



func formMermaidLine(stateToId map[*State]int, sourceState *State,option string,destinationState *State) string{
	return fmt.Sprintf("id:%d%s -- .%s. --> id:%d%s\n",
		stateToId[sourceState],
		sourceState.getEndMark(stateToId[sourceState]),
		handleToSuitMermaid(option),
		stateToId[destinationState],
		destinationState.getEndMark(stateToId[destinationState]),
	)
}
func handleToSuitMermaid(str string) string {
	strToSuitMermaid := make(map[string]string)
	strToSuitMermaid["-"] = "减号"
	strToSuitMermaid[","] = "逗号"
	strToSuitMermaid["("] = "左括号"
	strToSuitMermaid[")"] = "右括号"
	strToSuitMermaid["["] = "左中括号"
	strToSuitMermaid["]"] = "右中括号"
	strToSuitMermaid["{"] = "左大括号"
	strToSuitMermaid["}"] = "右大括号"
	strToSuitMermaid[";"] = "分号"
	strToSuitMermaid[`"`] = "双引号"
	strToSuitMermaid[`.`] = "小点"
	if strToSuitMermaid[str] == "" {
		return str
	}
	return strToSuitMermaid[str]
}

