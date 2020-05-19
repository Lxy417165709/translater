package stateMachine

import (
	"fmt"
	"regexpsManager"
)



// TODO: 重构！！！！！
type State struct {
	endFlag     bool
	toNextState map[byte][]*State
	token *regexpsManager.Token
}

func NewState(endFlag bool) *State {
	return &State{endFlag, make(map[byte][]*State),nil}
}
func (s *State) Link(nextState *State) {
	s.LinkByChar(regexpsManager.Eps, nextState)
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
	for _, nextState := range s.getNextStates(regexpsManager.Eps) {
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
	s.cleanNextStates(regexpsManager.Eps)
}
func (s *State) cleanNextStates(char byte) {
	delete(s.toNextState, char)
}


func (s *State) isNextBlankStatesHaveEndState() bool {

	return s.isNextStatesHaveEndState(regexpsManager.Eps)
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
	return s.getNextStates(regexpsManager.Eps)
}
func (s *State) getNextStates(char byte) []*State {
	return s.toNextState[char]
}




func (s *State) setEndFlag(value bool) {
	s.endFlag = value
}


// TODO: 用于标记该状态属于哪个自动机
func (s *State) InsertToken(nowWord string, specialChar byte,stateIsVisit map[*State]bool) {
	currentState := s
	if stateIsVisit[currentState] {
		return
	}
	stateIsVisit[currentState] = true
	if s.endFlag{
		s.token = regexpsManager.GetRegexpsManager().GetToken(specialChar,nowWord)
	}
	for char, nextStates := range s.toNextState {
		for _, nextState := range nextStates {
			nextState.InsertToken(nowWord+string(char),specialChar, stateIsVisit)
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
	if s.endFlag{
		partOne,partTwo := "",""
		if s.token.GetSpecialChar()!=0{

			partOne = "_" + string(s.token.GetSpecialChar())
		}
		if s.token.GetKindCode()!=0{
			partTwo = fmt.Sprintf("-%d",s.token.GetKindCode())
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
	if str =="{"{
		return "左大括号"
	}
	if str =="}"{
		return "右大括号"
	}
	if str =="["{
		return "左中括号"
	}
	if str =="]"{
		return "右中括号"
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


