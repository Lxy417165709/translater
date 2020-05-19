package stateMachine

import (
	"fmt"
	"grammar"
)

type State struct {
	endFlag             bool
	toNextState         map[byte][]*State
	belongToSpecialChar byte
}

func NewState(endFlag bool) *State {
	return &State{endFlag: endFlag, toNextState: make(map[byte][]*State)}
}
func (s *State) Link(nextState *State) {
	s.LinkByChar(grammar.Eps, nextState)
}
func (s *State) LinkByChar(ch byte, nextState *State) {
	s.toNextState[ch] = append(s.toNextState[ch], nextState)
}

func (s *State) EliminateNextBlankStatesFromHere(hasVisited map[*State]bool) {
	if hasVisited[s] {
		return
	}
	hasVisited[s] = true
	// 消除空白态
	for s.haveBlankStates() {
		mapOfReachableStateOfBlankStates := s.formMapOfReachableStateOfBlankStates()
		if s.isNextBlankStatesHaveEndState() {
			s.setEndFlag(true)
		}
		s.cleanBlankStates()
		s.addNextStates(mapOfReachableStateOfBlankStates)
	}

	//对非空白态的子节点进行处理
	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		nextState.EliminateNextBlankStatesFromHere(hasVisited)
	}
	return
}
func (s *State) formMapOfReachableStateOfBlankStates() map[byte][]*State {
	blankStates := s.getNextBlankStates()
	return getStatesToNext(blankStates)
}
func (s *State) haveBlankStates() bool {
	return len(s.getNextBlankStates()) != 0
}
func (s *State) getNextBlankStates() []*State {
	return s.getNextStates(grammar.Eps)
}
func (s *State) isNextBlankStatesHaveEndState() bool {
	return s.isNextStatesHaveEndState(grammar.Eps)
}
func (s *State) isNextStatesHaveEndState(char byte) bool {
	for _, state := range s.toNextState[char] {
		if state.endFlag == true {
			return true
		}
	}
	return false
}
func (s *State) setEndFlag(value bool) {
	s.endFlag = value
}
func (s *State) cleanBlankStates() {
	s.cleanNextStates(grammar.Eps)
}
func (s *State) cleanNextStates(char byte) {
	delete(s.toNextState, char)
}
func (s *State) addNextStates(addedMap map[byte][]*State) {
	for char, states := range addedMap {
		s.toNextState[char] = append(s.toNextState[char], states...)
	}
}


func (s *State) MarkSpecialChar(specialChar byte, hasVisited map[*State]bool) {
	if hasVisited[s] {
		return
	}
	hasVisited[s] = true
	s.setSpecialChar(specialChar)

	//对非空白态的子节点进行处理
	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		nextState.MarkSpecialChar(specialChar, hasVisited)
	}
	return
}
func (s *State) setSpecialChar(specialChar byte) {
	s.belongToSpecialChar = specialChar
}
func (s *State) getAllNextStates() []*State {
	result := make([]*State, 0)
	for char := range s.toNextState {
		result = append(result, s.getNextStates(char)...)
	}
	return result
}
func (s *State) getNextStates(char byte) []*State {
	return s.toNextState[char]
}


func (s *State) GetShowDataFromHere(startId int, stateToId map[*State]int, stateIsVisit map[*State]bool, line *int, result *[]string) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	stateToId[currentState] = startId
	for bytes, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			*line++
			nextState.GetShowDataFromHere(len(stateToId), stateToId, stateIsVisit, line, result)
			option := string(bytes)
			*result = append(*result, fmt.Sprintf("id:%d%s -- .%s. --> id:%d%s\n",
				stateToId[currentState],
				currentState.getEndMark(stateToId[currentState]),
				handleToSuitMermaid(option),
				stateToId[nextState],
				nextState.getEndMark(stateToId[nextState]),
			))

		}
	}
}
func (s *State) getEndMark(id int) string {
	if s.endFlag {
		return fmt.Sprintf("((%d))", id)
	} else {
		return fmt.Sprintf("(%d)", id)
	}
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
func getStatesToNext(states []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	hasExist := make(map[byte]map[*State]bool)
	for _, state := range states {
		for char, nextStates := range state.toNextState {
			for _, nextState := range nextStates {
				if hasExist[char] == nil {
					hasExist[char] = make(map[*State]bool)
				}
				if hasExist[char][nextState] {
					continue
				}
				hasExist[char][nextState] = true
				result[char] = append(result[char], nextState)
			}
		}
	}
	return result
}
