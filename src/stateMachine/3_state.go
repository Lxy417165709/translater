package stateMachine

import "fmt"



// TODO: 重构！！！！！
type State struct {
	endFlag     bool
	markFlag    byte
	code int
	toNextState map[byte][]*State
}

func NewState(endFlag bool) *State {
	return &State{endFlag, eps,  0,make(map[byte][]*State)}
}
func (s *State) GetWordEndPair(word string,hasVisit map[*State]bool) []*WordEndPair{
	if hasVisit[s]{
		return []*WordEndPair{}
	}
	hasVisit[s] = true

	if s==nil{
		return []*WordEndPair{}
	}
	result := make([]*WordEndPair,0)
	if s.endFlag{
		result = append(result,&WordEndPair{s,word})
	}
	for char, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			result = append(result,nextState.GetWordEndPair(word + string(char),hasVisit)...)
		}
	}
	return result
}
func (s *State) SetCode(code int) {
	s.code = code
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




func (s *State) AddNextStates(addedMap map[byte][]*State) {
	for char, states := range addedMap {
		s.toNextState[char] = append(s.toNextState[char], states...)
	}
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




func (s *State) setEndFlag(value bool) {
	s.endFlag = value
}


// TODO: 用于标记该状态属于哪个自动机
func (s *State) MarkDown(specialChar byte, stateIsVisit map[*State]bool) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	s.markFlag = specialChar
	for _, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			nextState.MarkDown(specialChar, stateIsVisit)
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
				HandleToSuitMermaid(option),
				stateToId[nextState],
				nextState.getEndMark(stateToId[nextState]),
			))

		}
	}
}
func (s *State) getEndMark(value int) string {
	//if s.endFlag == true {
	//	return fmt.Sprintf("{%d_%s-%d}",value,string(s.markFlag),s.code)
	//}
	//return fmt.Sprintf("((%d))",value)
	if s.endFlag == true {
		partOne,partTwo := "",""
		if s.markFlag!=0{
			partOne = "_" + string(s.markFlag)
		}
		if s.code!=0{
			partTwo = fmt.Sprintf("-%d",s.code)
		}

		return fmt.Sprintf("{%d%s%s}",value,partOne,partTwo)
	}
	return fmt.Sprintf("((%d))",value)
}
func HandleToSuitMermaid(str string) string{
	if str =="-"{
		return "减号"
	}
	if str =="."{
		return "逗号"
	}
	if str =="("{
		return "左括号"
	}
	if str ==")"{
		return "右括号"
	}
	if str == ";"{
		return "分号"
	}
	return str
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


