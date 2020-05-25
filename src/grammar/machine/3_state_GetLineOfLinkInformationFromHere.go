package machine

import "fmt"

func (s *State) GetLineOfLinkInformationFromHere(startId int, StateToId map[*State]int, StateIsVisit map[*State]bool) []string{
	if StateIsVisit[s] {
		return []string{}
	}
	result := make([]string,0)
	StateIsVisit[s] = true
	StateToId[s] = startId
	for bytes, nextStates := range s.next {
		for _, nextState := range nextStates {
			result = append(result, nextState.GetLineOfLinkInformationFromHere(len(StateToId), StateToId, StateIsVisit)...)
			option := string(bytes)
			result = append(result, formMermaidLine(StateToId,s,option,nextState))
		}
	}
	return result
}
func (s *State) getEndMark(id int) string {
	if s.isEnd {
		return fmt.Sprintf("((%d))", id)
	} else {
		return fmt.Sprintf("(%d)", id)
	}
}


func formMermaidLine(StateToId map[*State]int, sourceState *State,option string,destinationState *State) string{
	return fmt.Sprintf("id:%d%s -- .%s. --> id:%d%s\n",
		StateToId[sourceState],
		sourceState.getEndMark(StateToId[sourceState]),
		handleToSuitMermaid(option),
		StateToId[destinationState],
		destinationState.getEndMark(StateToId[destinationState]),
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

