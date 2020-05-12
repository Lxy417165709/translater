package lexicalTest

type NFA struct {
	startState *State
	endState *State
}


// 空NFA能匹配空字符串
func NewNFA() *NFA{
	startState := NewState(false)
	endState := NewState(true)
	startState.addNextState(' ',endState)
	return &NFA{startState,endState}
}


func (nfa *NFA) IsMatch(pattern string) bool{
	return nfa.startState.IsMatch(pattern)
}
func (nfa *NFA) Show(){
	nfa.startState.Show(0,make(map[*State]int),make(map[*State]bool))
}

// TODO: 可以统一
func (nfa *NFA) RepeatPlus(ch byte) {
	nfa.breakDown()
	if !GlobalNFAManager.IdentityIsExist(ch) {
		newState := NewState(true)
		newState.addNextState(ch,newState)
		nfa.endState.addNextState(ch,newState)
		nfa.endState = newState
	}else{
		// Testing
		theNFAofChar := GlobalNFAManager.Get(ch).BuildNFA()
		theNFAofChar.breakDown()
		theNFAofChar.endState.addNextState(' ',theNFAofChar.startState)
		nfa.endState.addNextState(' ',theNFAofChar.startState)
		nfa.endState = theNFAofChar.endState
	}
	nfa.open()
}
func (nfa *NFA) RepeatZero(ch byte) {
	nfa.breakDown()
	if !GlobalNFAManager.IdentityIsExist(ch) {
		newState := NewState(true)
		newState.addNextState(ch,newState)
		nfa.endState.addNextState(' ',newState)
		nfa.endState.addNextState(ch,newState)
		nfa.endState = newState
	}else{
		// Testing
		theNFAofChar := GlobalNFAManager.Get(ch).BuildNFA()
		theNFAofChar.breakDown()
		theNFAofChar.endState.addNextState(' ',theNFAofChar.startState)
		nfa.endState.addNextState(' ',theNFAofChar.startState)
		nfa.endState.addNextState(' ',theNFAofChar.endState)
		nfa.endState = theNFAofChar.endState

	}
	nfa.open()
}
func (nfa *NFA) Once(ch byte) {
	nfa.breakDown()
	if !GlobalNFAManager.IdentityIsExist(ch) {
		newState := NewState(true)
		nfa.endState.addNextState(ch,newState)
		nfa.endState = newState
	}else{
		theNFAofChar := GlobalNFAManager.Get(ch).BuildNFA()
		nfa.AddSeriesGraph(theNFAofChar)
	}
	nfa.open()
}


func (nfa *NFA) AddParallelGraph(addedGraph *NFA) {
	addedGraph.breakDown()

	nfa.startState.addNextState(' ',addedGraph.startState)
	addedGraph.endState.addNextState(' ',nfa.endState)
}
func (nfa *NFA) AddSeriesGraph(addedGraph *NFA) {
	nfa.breakDown()
	nfa.endState.addNextState(' ',addedGraph.startState)
	nfa.endState = addedGraph.endState
}


func (nfa *NFA) breakDown() {
	nfa.endState.endFlag = false
}
func (nfa *NFA) open() {
	nfa.endState.endFlag = true
}
