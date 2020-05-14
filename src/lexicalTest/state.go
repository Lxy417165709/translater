package lexicalTest

import "fmt"

const eps = ' '

// TODO: 重构
type State struct {
	endFlag     bool
	toNextState map[byte][]*State
}

func NewState(endFlag bool) *State {
	return &State{endFlag, make(map[byte][]*State)}
}
func (s *State) Show(startId int, stateToId map[*State]int, stateIsVisit map[*State]bool) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	stateToId[currentState] = startId
	for bytes, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			startId++
			nextState.Show(startId, stateToId, stateIsVisit)
			option := string(bytes)
			fmt.Printf("id:%d%s --%s--> id:%d%s\n",
				stateToId[currentState],
				currentState.getEndMark(),
				option,
				stateToId[nextState],
				nextState.getEndMark(),
			)
		}
	}
}

func (s *State) LinkByChar(ch byte, nextState *State) {
	s.toNextState[ch] = append(s.toNextState[ch], nextState)
}
func (s *State) Link(nextState *State) {
	s.LinkByChar(eps, nextState)
}

func (s *State) roadIsExist(ch byte) bool {
	return s.toNextState[ch] != nil
}

func (s *State) IsMatch(pattern string) bool {
	// 空匹配
	nextStates := s.toNextState[eps]
	for _, nextState := range nextStates {
		if nextState.IsMatch(pattern) {
			return true
		}
	}
	if pattern == "" {
		return s.endFlag
	}

	ch := pattern[0]

	// 实匹配
	nextStates = s.toNextState[ch]
	for _, nextState := range nextStates {
		if nextState.IsMatch(pattern[1:]) {
			return true
		}
	}

	return false
}

func (s *State) IsMatchByDFA(pattern string) bool {

	// 空匹配
	nextStates := s.toNextState[eps]

	// 判断DFA是否正确
	if len(nextStates) >= 2 {
		panic("DFA存在错误")
	}

	if len(nextStates) == 1 {
		if nextStates[0].IsMatchByDFA(pattern) {
			return true
		}
	}
	if pattern == "" {
		return s.endFlag
	}
	ch := pattern[0]

	// 实匹配
	nextStates = s.toNextState[ch]

	// 判断DFA是否正确
	if len(nextStates) >= 2 {
		panic("DFA存在错误")
	}

	if len(nextStates) == 1 {
		if nextStates[0].IsMatchByDFA(pattern[1:]) {
			return true
		}
	}

	return false
}

func (s *State) getEndMark() string {
	if s.endFlag == true {
		return "(OK)"
	}
	return "    "
}

func (s *State) DFA(hasVisited map[*State]bool) *State {
	if hasVisited[s] {
		return s
	}
	hasVisited[s] = true

	for firstLayChar, firstLayStates := range s.toNextState {
		if len(firstLayStates) == 1 {
			s.toNextState[firstLayChar][0] = s.toNextState[firstLayChar][0].DFA(hasVisited)
		} else {
			compoundState := NewState(s.hasEndFlag(firstLayStates))
			compoundState.toNextState = s.getStatesToNext(firstLayStates)
			if s.hasSelf(firstLayStates) {
				compoundState.LinkByChar(firstLayChar, compoundState)
			}

			s.toNextState[firstLayChar] = s.toNextState[firstLayChar][:0] //清空
			s.toNextState[firstLayChar] = append(s.toNextState[firstLayChar], compoundState.DFA(hasVisited))
		}
	}
	return s
}

func (s *State) getStatesToNext(states []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	for _, state := range states {
		if state == s {
			continue
		}
		for char, nextStates := range state.toNextState {
			//fmt.Printf("wwwww         %p %s %v\n",state,string(char),nextStates)
			result[char] = append(result[char], nextStates...)
		}
	}
	return result
}

func (s *State) hasSelf(states []*State) bool {
	for _, state := range states {
		if state == s {
			return true
		}
	}
	return false
}

func (s *State) hasEndFlag(states []*State) bool {
	for _, state := range states {
		if state.endFlag {
			return true
		}
	}
	return false
}

// 要求终止状态不成环
func (s *State) GetFinalStates(hasVisited map[*State]bool) []*State {
	if hasVisited[s] || s == nil {
		return []*State{}
	}
	hasVisited[s] = true

	result := make([]*State, 0)
	if len(s.toNextState) == 0 {
		result = append(result, s)
		return result
	}

	for _, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			result = append(result, nextState.GetFinalStates(hasVisited)...)
		}
	}
	return result
}
