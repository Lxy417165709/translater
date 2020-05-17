package stateMachine

import "fmt"



// TODO: 重构！！！！！
type State struct {
	endFlag     bool
	markFlag    byte
	toNextState map[byte][]*State
}

func NewState(endFlag bool) *State {
	return &State{endFlag, eps, make(map[byte][]*State)}
}

func (s *State) Link(nextState *State) {
	s.LinkByChar(eps, nextState)
}
func (s *State) LinkByChar(ch byte, nextState *State) {
	if s.stateIsLiving(ch,nextState){
		return
	}
	s.toNextState[ch] = append(s.toNextState[ch], nextState)
}
func (s *State) stateIsLiving(char byte,beJudgeState *State) bool{
	for _,state := range s.getNextStates(char){
		if state==beJudgeState{
			return true
		}
	}
	return false
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
		s.AddNextStates(mapOfReachableStateOfBlankStates)
	}

	//对非空白态的子节点进行处理
	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		nextState.EliminateNextBlankStatesFromHere(hasVisited)
	}
	return
}
func (s *State) MultiWayMergeFromHere(hasVisited map[*State]bool) *State {
	if hasVisited[s] {
		return s
	}
	hasVisited[s] = true
	for char := range s.toNextState {
		dfaState := s.getDFAState(char)
		s.cleanNextStates(char)
		s.LinkByChar(char, dfaState.MultiWayMergeFromHere(hasVisited))
	}
	return s
}





func (s *State) AddNextStates(addedMap map[byte][]*State) {
	for char, states := range addedMap {
		s.toNextState[char] = append(s.toNextState[char], states...)
	}
}

// TODO： N -> X | Z 有问题..  这个的DFA不正确...
// TODO: 这个函数还是有问题的     D+.D+|D+ 这种情况不能判断
func (s *State) getDFAState(char byte) *State {
	states := s.toNextState[char]
	if len(states) == 1 {
		return states[0]
	}

	dfaState := NewState(s.isNextStatesHaveEndState(char))
	dfaState.toNextState = s.formMapOfReachableStateOfAllNextStates()
	//fmt.Println("----",dfaState.toNextState)
	if s.toNextIsSame(dfaState) {
		//fmt.Printf("%p %v %p %v\n",s,s,dfaState,dfaState)
		dfaState = s.toNextState[char][0]
		return dfaState
	}
	if s.hasSelf(char) {
		dfaState.LinkByChar(char, dfaState)
	}
	return dfaState
}

func (s *State) toNextIsSame(reference *State) bool {
	if len(s.toNextState) != len(reference.toNextState) {
		return false
	}

	for char, nextStates := range reference.toNextState {
		HasVisitOfS := make(map[*State]bool)
		HasVisitOfRef := make(map[*State]bool)
		for _, nextState1 := range nextStates {
			HasVisitOfRef[nextState1] = true
		}
		for _, nextState1 := range s.toNextState[char] {
			HasVisitOfS[nextState1] = true
		}
		//if len(HasVisitOfS)!=len(HasVisitOfRef){
		//	return false
		//}
		for state := range HasVisitOfS {
			if HasVisitOfRef[state] == false {
				return false
			}
		}
		for state := range HasVisitOfRef {
			if HasVisitOfS[state] == false {
				return false
			}
		}

	}
	return true
}



func (s *State) CanBeStartOfDFA(hasVisited map[*State]bool) bool {
	if hasVisited[s] {
		return true
	}
	hasVisited[s] = true
	for char := range s.toNextState {
		if len(s.toNextState[char])!=1{
			return false
		}
	}

	allNextStates := s.getAllNextStates()
	for _, nextState := range allNextStates {
		if nextState.CanBeStartOfDFA(hasVisited)==false{
			return false
		}
	}
	return true
}
func (s *State) IsMatch(pattern string) bool {
	// 空匹配
	for _, nextState := range s.getNextStates(eps) {
		if nextState.IsMatch(pattern) {
			return true
		}
	}
	if pattern == "" {
		return s.endFlag
	}

	char := pattern[0]

	// 实匹配
	for _, nextState := range s.getNextStates(char) {
		if nextState.IsMatch(pattern[1:]) {
			return true
		}
	}
	return false
}

// 生成 mermaid 图
func (s *State) ShowFromHere(startId int, stateToId map[*State]int, stateIsVisit map[*State]bool,line *int) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	stateToId[currentState] = startId
	for bytes, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			*line++
			nextState.ShowFromHere(len(stateToId), stateToId, stateIsVisit,line)
			option := string(bytes)
			fmt.Printf("id:%d%s -- .%s. --> id:%d%s\n",
				stateToId[currentState],
				currentState.getEndMark(stateToId[currentState]),
				option,
				stateToId[nextState],
				nextState.getEndMark(stateToId[nextState]),
			)
		}
	}
}


func (s *State) GetShowDataFromHere(startId int, stateToId map[*State]int, stateIsVisit map[*State]bool,line *int,result *[]string){
	currentState := s

	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	stateToId[currentState] = startId
	for bytes, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			*line++
			nextState.GetShowDataFromHere(len(stateToId), stateToId, stateIsVisit,line,result)
			option := string(bytes)
			*result = append(*result,fmt.Sprintf("id:%d%s -- .%s. --> id:%d%s\n",
				stateToId[currentState],
				currentState.getEndMark(stateToId[currentState]),
				option,
				stateToId[nextState],
				nextState.getEndMark(stateToId[nextState]),
			))

		}
	}
}



func (s *State) hasSelf(char byte) bool {
	for _, state := range s.toNextState[char] {
		if state == s {
			return true
		}
	}
	return false
}




func (s *State) cleanBlankStates() {
	s.cleanNextStates(eps)
}
func (s *State) cleanNextStates(char byte) {
	delete(s.toNextState, char)
}


func (s *State) isNextBlankStatesHaveEndState() bool {

	return s.isNextStatesHaveEndState(eps)
}
func (s *State) isNextStatesHaveEndState(char byte) bool {
	for _, state := range s.toNextState[char] {
		if state.endFlag == true {
			return true
		}
	}
	return false
}

func (s *State) formMapOfReachableStateOfBlankStates() map[byte][]*State {
	blankStates := s.getNextBlankStates()
	return getStatesToNext(blankStates)
}
func (s *State) formMapOfReachableStateOfAllNextStates() map[byte][]*State {
	allNextStates := s.getAllNextStates()
	return getStatesToNext(allNextStates)
}




func (s *State) haveBlankStates() bool {
	return len(s.getNextBlankStates()) != 0
}


func (s *State) getAllNextStates() []*State {
	result := make([]*State, 0)
	for char := range s.toNextState {
		result = append(result, s.getNextStates(char)...)
	}
	return result
}
func (s *State) getNextBlankStates() []*State {
	return s.getNextStates(eps)
}
func (s *State) getNextStates(char byte) []*State {
	return s.toNextState[char]
}





func (s *State) getEndMark(value int) string {
	if s.endFlag == true {
		return fmt.Sprintf("{%d}",value)
	}
	return fmt.Sprintf("((%d))",value)
}
func (s *State) setEndFlag(value bool) {
	s.endFlag = value
}


func getStatesToNext(states []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	hasExist := make(map[byte]map[*State]bool)
	for _, state := range states {
		for char, nextStates := range state.toNextState {
			for _,nextState := range nextStates{
				if hasExist[char]==nil{
					hasExist[char]=make(map[*State]bool)
				}
				if hasExist[char][nextState]{
					continue
				}
				hasExist[char][nextState]=true
				result[char] = append(result[char], nextState)
			}
		}
	}
	return result
}
