package lexicalTest

import "fmt"

const eps = ' '

// TODO: 重构！！！！！
type State struct {
	endFlag     bool
	markFlag byte
	toNextState map[byte][]*State
}

func NewState(endFlag bool) *State {
	return &State{endFlag, eps,make(map[byte][]*State)}
}

func (s *State) getFather(fatherSet map[*State]*State) *State{
	return fatherSet[s]
}
func (s *State) Merge(hasVisited map[*State]bool){
	if hasVisited[s]{
		return
	}
	hasVisited[s]=true
	// 消除空白态
	for len(s.getNextBlankStates())!=0{
		nextBlankStates := s.getNextBlankStates()
		s.cleanNextBlankStates()
		s.AddNextStates(s.getStatesToNext(nextBlankStates))
		for _,nextBlankState:=range nextBlankStates{
			if nextBlankState.endFlag==true{
				s.endFlag = true
			}
		}
	}

	//对非空白态的子节点进行处理
	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		nextState.Merge(hasVisited)
	}
	return
}
func (s *State) getNextBlankStates() []*State{
	return s.getNextStates(eps)
}
func (s *State) getAllNextStates() []*State{
	result := make([]*State,0)
	for char := range s.toNextState{
		result = append(result,s.getNextStates(char)...)
	}
	return result
}


func (s *State) AddNextStates(addedMap map[byte][]*State) {
	for char,states := range addedMap{
		s.toNextState[char] = append(s.toNextState[char], states...)
	}
}


func (s *State) cleanNextBlankStates() {
	s.cleanNextStates(eps)
}
func (s *State) cleanNextStates(char byte) {
	delete(s.toNextState,char)
}

//var i=0
// TODO： N -> X | Z 有问题..  这个的DFA不正确...
func (s *State) DFA(hasVisited map[*State]bool) *State {
	//i++
	//if i==10{
	//	panic("taiduo")
	//}
	//fmt.Printf("%p %v\n",s,s)
	if hasVisited[s] {
		return s
	}
	hasVisited[s] = true

	for firstLayChar, firstLayStates := range s.toNextState {
		states := make([]*State,0)
		for _,state := range firstLayStates{
			states=append(states,state)
		}

		if len(states) == 0{
			continue
		}
		if len(states) == 1 {
			s.cleanNextStates(firstLayChar)
			s.toNextState[firstLayChar] = append(s.toNextState[firstLayChar],states[0].DFA(hasVisited))
		} else {
			compoundState := NewState(s.hasEndFlag(states))
			compoundState.toNextState = s.getStatesToNext(states)

			// 重复的取第一个就可以了
			if s.toNextIsSame(compoundState) && compoundState.toNextIsSame(s){
				onlyState := s.toNextState[firstLayChar][0]
				s.cleanNextStates(firstLayChar)
				s.toNextState[firstLayChar] = append(s.toNextState[firstLayChar], onlyState.DFA(hasVisited))
				return s
			}
			if s.hasSelf(states) {
				compoundState.LinkByChar(firstLayChar, compoundState)
			}
			s.cleanNextStates(firstLayChar)
			s.toNextState[firstLayChar] = append(s.toNextState[firstLayChar], compoundState.DFA(hasVisited))
		}
	}
	return s
}
func (s *State)toNextIsSame(reference *State) bool{

	for byte,nextStates := range reference.toNextState{
		hasVisit := make(map[*State]bool)
		for _,nextState := range nextStates{
			hasVisit[nextState]=true
		}
		for _,nextState := range s.toNextState[byte]{
			if hasVisit[nextState]==false {
				return false
			}
		}
	}
	return true
}


func (s *State) stateIsLiving(char byte,x *State) bool{
	for _,state := range s.toNextState[char]{
		if state == x{
			return true
		}
	}
	return false
}



func (s *State) CanBeStartOfDFA(hasVisited map[*State]bool) bool{
	if hasVisited[s]{
		return true
	}
	hasVisited[s] = true

	charsOfLinkingToNextStates := s.getTheCharsOfLinkingToNextStates()
	for _,char := range charsOfLinkingToNextStates{
		nextStates := s.getNextStates(char)
		if len(nextStates)!=1{
			return false
		}
		// 往后搜索
		for _,state := range nextStates{
			if state.CanBeStartOfDFA(hasVisited)==false{
				return false
			}
		}
	}
	return true
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

func (s *State) MarkDown(specialChar byte,stateIsVisit map[*State]bool) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	s.markFlag = specialChar
	for _, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			nextState.MarkDown(specialChar,stateIsVisit)
		}
	}
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
			nextState.Show(len(stateToId), stateToId, stateIsVisit)
			option := string(bytes)
			fmt.Printf("id:%d%s|%s --%s--> id:%d%s|%s\n",
				stateToId[currentState],
				currentState.getEndMark(),
				string(currentState.markFlag),
				option,
				stateToId[nextState],
				nextState.getEndMark(),
				string(currentState.markFlag),
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


func (s *State) getNextStates(char byte) []*State{
	return s.toNextState[char]
}
func (s *State) getTheCharsOfLinkingToNextStates() []byte{
	chars := make([]byte,0)
	for char := range s.toNextState{
		chars = append(chars, char)
	}
	return chars
}



func (s *State) getStatesToNext(states []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	for _, state := range states {
		for char, nextStates := range state.toNextState {
			for _,nextState := range nextStates{
				result[char] = append(result[char], nextState)
			}
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




func (s *State) getEndMark() string {
	if s.endFlag == true {
		return "(OK)"
	}
	return "    "
}
