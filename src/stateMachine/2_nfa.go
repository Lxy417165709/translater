package stateMachine

import (
	"fmt"
	"grammar"
	"os"
)

const endSymbol = '#'


type NFA struct {
	startState            *State
	endState              *State
	specialChar byte
}

func NewNFA() *NFA {
	startState :=NewState(false)
	endState :=NewState(true)
	return &NFA{
		startState:startState,
		endState:endState,
	}
}

func (nfa *NFA) SetSpecialChar(char byte) {
	nfa.specialChar = char
}
func (nfa *NFA) GetSpecialChar() byte {
	return nfa.specialChar
}

func (nfa *NFA) EliminateBlankStates() *NFA{
	hasVisited := make(map[*State]bool)
	nfa.getStartState().EliminateNextBlankStatesFromHere(hasVisited)
	return nfa
}





func (nfa *NFA) GetTokens(pattern string) []*grammar.Token {
	pattern += string(endSymbol)
	buffer := ""
	tokens, queue := make([]*grammar.Token, 0), make([]*State, 0)
	queue = append(queue, nfa.startState)
	readingPosition := 0
	for pattern[readingPosition] != endSymbol {
		lastEndState := getFirstEndState(queue)
		queue = getNextStates(queue, pattern[readingPosition])
		if len(queue) != 0 {
			buffer += string(pattern[readingPosition])
			readingPosition++
			continue
		}
		if lastEndState == nil && !isBlank(pattern[readingPosition]) {
			panic(fmt.Sprintf("源文件存在非法字符：%s 索引:%d", string(pattern[readingPosition]), readingPosition))
		}
		switch {
		case lastEndState != nil:
			token := grammar.GetRegexpsManager().GetToken(lastEndState.belongToSpecialChar,buffer).Copy()
			tokens = append(tokens, token)
		case isBlank(pattern[readingPosition]):
			readingPosition++
		}
		buffer = ""
		queue = nil
		queue = append(queue, nfa.startState)
	}
	return tokens
}

func (nfa *NFA) MatchMoreThanOnce(beAddedNFA *NFA) *NFA{
	nfa.AddSeriesNFA(beAddedNFA.linkEndStateToStartState())
	return nfa
}
func (nfa *NFA) MatchMoreThanZeroTimes(beAddedNFA *NFA) *NFA{
	endStateOfShouldAddNFA := NewState(true)
	beAddedNFA.linkEndStateToStartState().linkStartStateTo(endStateOfShouldAddNFA).setEndState(endStateOfShouldAddNFA)
	return nfa.AddSeriesNFA(beAddedNFA)
}
func (nfa *NFA) MatchOnce(beAddedNFA *NFA)*NFA {
	return nfa.AddSeriesNFA(beAddedNFA)
}
func (nfa *NFA) AddParallelNFA(beAddedNFA *NFA) *NFA{
	beAddedNFA.breakDown()
	beAddedNFA.linkEndStateTo(nfa.getEndState())
	return nfa.linkStartStateTo(beAddedNFA.getStartState())
}
func (nfa *NFA) AddSeriesNFA(beAddedNFA *NFA) *NFA{
	nfa.breakDown()
	nfa.linkEndStateTo(beAddedNFA.getStartState()).setEndState(beAddedNFA.getEndState())
	return nfa
}
func (nfa *NFA) breakDown() {
	nfa.getEndState().endFlag = false
}





func (nfa *NFA) FormTheMermaidGraphOfNFA(filePath string) error {
	var file *os.File
	var err error
	if file, err = os.Create(filePath);err!=nil{
		return err
	}
	defer file.Close()
	for _,line := range nfa.getMermaidLines() {
		if _,err = file.WriteString(line);err!=nil{
			return err
		}
	}
	return err
}
func (nfa *NFA) getMermaidLines() []string{
	lines := make([]string,0)
	lines = append(lines, "```mermaid\ngraph LR\n")
	lines = append(lines,nfa.getMetaMermaidData()...)
	lines = append(lines,"```\n")
	return lines
}
func (nfa *NFA) getMetaMermaidData() []string{
	metaMermaidData := make([]string,0)
	nfa.getStartState().GetShowDataFromHere(0, make(map[*State]int), make(map[*State]bool), &metaMermaidData)
	return metaMermaidData
}





func (nfa *NFA) linkStartStateToEndState () *NFA{
	return nfa.linkStartStateTo(nfa.endState)
}
func (nfa *NFA) linkStartStateToEndStateByChar (char byte) *NFA{
	nfa.getStartState().LinkByChar(char,nfa.getEndState())
	return nfa
}
func (nfa *NFA) linkEndStateToStartState() *NFA{
	nfa.endState.Link(nfa.startState)
	return nfa
}
func (nfa *NFA) linkStartStateTo(state *State) *NFA{
	nfa.getStartState().Link(state)
	return nfa
}
func (nfa *NFA) linkEndStateTo(state *State) *NFA{
	nfa.getEndState().Link(state)
	return nfa
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

func (nfa *NFA) MarkSpecialChar() *NFA{
	nfa.startState.MarkSpecialChar(nfa.specialChar,make(map[*State]bool))
	return nfa
}

func getNextStates(states []*State, readingChar byte) []*State {
	tmpQueue := make([]*State, 0)
	for i := 0; i < len(states); i++ {
		if states[i].toNextState[readingChar] != nil {
			tmpQueue = append(tmpQueue, states[i].toNextState[readingChar]...)
		}
	}
	return tmpQueue
}
func getFirstEndState(states []*State) *State {
	for _, state := range states {
		if state.endFlag {
			return state
		}
	}
	return nil
}



