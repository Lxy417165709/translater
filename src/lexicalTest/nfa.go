package lexicalTest

import "fmt"

type NFA struct {
	startState *State
	endState   *State
}

func NewNFA(char byte) *NFA {
	startState := NewState(false)
	endState := NewState(true)
	startState.LinkByChar(char,endState)
	return &NFA{startState, endState}
}
//func NewTestNFA() *NFA {
//	s0 := NewState(false)
//	s1 := NewState(false)
//	s2 := NewState(false)
//	s3 := NewState(false)
//	s4 := NewState(false)
//	//s0.LinkByChar('a',s1)
//	s0.Link(s1)
//	s1.Link(s2)
//	s1.LinkByChar('a',s4)
//	s2.LinkByChar('a',s3)
//	s3.LinkByChar('b',s1)
//	s4.endFlag=true
//
//
//
//
//
//	return &NFA{s0,nil}
//}
//func NewTestNFA2() *NFA {
//	s0 := NewState(false)
//	s1 := NewState(false)
//	s2 := NewState(false)
//	s3 := NewState(false)
//	//s4 := NewState(false)
//	//id:0     --a--> id:1
//	//id:3     --a--> id:1
//	//id:3     --a--> id:2
//	//id:2     --b--> id:3
//	//id:0     --a--> id:2
//
//	s0.LinkByChar('a',s1)
//	s3.LinkByChar('a',s1)
//	s3.LinkByChar('a',s2)
//	s2.LinkByChar('b',s3)
//	s0.LinkByChar('a',s2)
//
//
//	return &NFA{s0,nil}
//}
//func NewTestNFA3() *NFA {
//	s0 := NewState(false)
//	s1 := NewState(false)
//
//	//s4 := NewState(false)
//	//id:0     --a--> id:1
//	//id:3     --a--> id:1
//	//id:3     --a--> id:2
//	//id:2     --b--> id:3
//	//id:0     --a--> id:2
//
//	s0.LinkByChar('d',s1)
//	s1.LinkByChar('d',s1)
//	s0.LinkByChar('d',s0)
//	// TODO: 要处理这种情况..
//
//	return &NFA{s0,nil}
//}



func (nfa *NFA) Merge(){
	hasVisited := make(map[*State]bool)
	nfa.getStartState().Merge(hasVisited)
}


func (nfa *NFA)MarkDown(specialChar byte) {
	nfa.startState.MarkDown(specialChar,make(map[*State]bool))
}


func (nfa *NFA) Show() {
	ids := make(map[*State]int)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("是否DFA:",nfa.IsDFA())
	nfa.getStartState().Show(0, ids, make(map[*State]bool))
	fmt.Println(ids)
	fmt.Println("-------------------------------------------------------------")
}
func (nfa *NFA) ChangeToDFA() {
	// TODO: 这可能有些问题，可能nfa.endState会发生改变
	hasVisited := make(map[*State]bool)
	nfa.getStartState().DFA(hasVisited)
}
func (nfa *NFA) Get(pattern string) []string{

	result := make([]string,0)
	begin := nfa.startState
	buffer := ""
	for position := 0;position<len(pattern);position++{
		if pattern[position]=='#'{
			break
		}
		//fmt.Printf("%d-%v-%v-\n",position,pattern[position],string(pattern[position]))
		// 不匹配
		if len(begin.toNextState[pattern[position]])==0{
			if begin.endFlag{
				result = append(result,buffer)
				fmt.Print(buffer)
			}
			begin = nfa.startState
			buffer = ""

			if len(begin.toNextState[pattern[position]])!=0{
				position--
			}else{
				fmt.Print(string(pattern[position]))
			}
			continue
		}
		// 成功匹配
		buffer+=string(pattern[position])
		begin = begin.toNextState[pattern[position]][0]
	}
	if buffer!=""{
		result = append(result,buffer)
	}
	return result
}
func (nfa *NFA) IsMatch(pattern string) bool {
	//if nfa.IsDFA(){
	//	fmt.Println("		「DFA匹配」		")
	//}
	return nfa.startState.IsMatch(pattern)
}
func (nfa *NFA) IsDFA() bool {
	hasVisited := make(map[*State]bool)
	return nfa.getStartState().CanBeStartOfDFA(hasVisited)
}

func (nfa *NFA) RepeatPlus(char byte) {
	shouldAddNFA := charToNFA(char)
	shouldAddNFA.linkEndStateToStartState()
	nfa.AddSeriesNFA(shouldAddNFA)
}
func (nfa *NFA) RepeatZero(char byte) {
	shouldAddNFA  := charToNFA(char)
	shouldAddNFA .linkEndStateToStartState()

	endStateOfShouldAddNFA := NewState(true)
	shouldAddNFA.linkStartStateTo(endStateOfShouldAddNFA)
	shouldAddNFA.setEndState(endStateOfShouldAddNFA)

	nfa.AddSeriesNFA(shouldAddNFA)
}
func (nfa *NFA) Once(char byte) {
	beAddedNFA := charToNFA(char)
	nfa.AddSeriesNFA(beAddedNFA)
}

func (nfa *NFA) AddParallelNFA(beAddedNFA *NFA) {
	beAddedNFA.breakDown()
	nfa.getStartState().Link(beAddedNFA.getStartState())
	beAddedNFA.getEndState().Link(nfa.getEndState())
}
func (nfa *NFA) AddSeriesNFA(beAddedNFA *NFA) {
	nfa.breakDown()
	nfa.getEndState().Link(beAddedNFA.getStartState())
	nfa.setEndState(beAddedNFA.getEndState())
}

func (nfa *NFA) linkEndStateToStartState() {
	nfa.getEndState().Link(nfa.getStartState())
}
func (nfa *NFA) linkStartStateTo(state *State) {
	nfa.getStartState().Link(state)
}

func (nfa *NFA) breakDown() {
	nfa.getEndState().endFlag = false
}
func (nfa *NFA) getStartState() *State {
	return nfa.startState
}
func (nfa *NFA) getEndState() *State {
	return nfa.endState
}
func (nfa *NFA) setStartState(state *State) {
	nfa.startState = state
}
func (nfa *NFA) setEndState(state *State) {
	nfa.endState = state
}
