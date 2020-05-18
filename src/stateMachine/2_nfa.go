package stateMachine

import (
	"fmt"
	"os"
	"regexpsManager"
)

type NFA struct {
	startState            *State
	endState              *State
	regexpsManager        *regexpsManager.RegexpsManager
	respondingSpecialChar byte
}

func NewEmptyNFA(regexpsManager *regexpsManager.RegexpsManager) *NFA {
	return &NFA{NewState(false), NewState(true), regexpsManager, eps}
}
func NewNFA(char byte, regexpsManager *regexpsManager.RegexpsManager) *NFA {
	if !regexpsManager.CharIsSpecial(char) {
		nfa := NewEmptyNFA(regexpsManager)
		nfa.getStartState().LinkByChar(char, nfa.getEndState())
		return nfa
	}
	regexp := regexpsManager.GetRegexp(char)
	nfa := NewNFABuilder(regexp, regexpsManager).BuildNFA()
	return nfa
}

func (nfa *NFA) SetRespondingSpecialChar(char byte) {
	nfa.respondingSpecialChar = char
}
func (nfa *NFA) GetRespondingSpecialChar() byte {
	return nfa.respondingSpecialChar
}

func (nfa *NFA) EliminateBlankStates() {
	hasVisited := make(map[*State]bool)
	nfa.getStartState().EliminateNextBlankStatesFromHere(hasVisited)
}



const endSymbol = '#'

func (nfa *NFA) GetTokenByNFA(pattern string) []*Token {
	pattern += string(endSymbol)
	buffer := ""
	tokens, queue := make([]*Token, 0), make([]*State, 0)
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
			tokens = append(tokens, &Token{
				lastEndState.markFlag,
				lastEndState.code,
				buffer,
			})
		case isBlank(pattern[readingPosition]):
			readingPosition++
		}
		buffer = ""
		queue = nil
		queue = append(queue, nfa.startState)
	}
	return tokens
}

func isBlank(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
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


func (nfa *NFA) IsMatch(pattern string) bool {
	return nfa.startState.IsMatch(pattern)
}

func (nfa *NFA) RepeatPlus(char byte) {
	shouldAddNFA := NewNFA(char, nfa.regexpsManager)
	shouldAddNFA.linkEndStateToStartState()
	nfa.AddSeriesNFA(shouldAddNFA)
}
func (nfa *NFA) RepeatZero(char byte) {
	shouldAddNFA := NewNFA(char, nfa.regexpsManager)
	shouldAddNFA.linkEndStateToStartState()

	endStateOfShouldAddNFA := NewState(true)
	shouldAddNFA.linkStartStateTo(endStateOfShouldAddNFA)
	shouldAddNFA.setEndState(endStateOfShouldAddNFA)

	nfa.AddSeriesNFA(shouldAddNFA)
}
func (nfa *NFA) Once(char byte) {
	beAddedNFA := NewNFA(char, nfa.regexpsManager)
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

func (nfa *NFA) MarkDown() *NFA {
	nfa.startState.MarkDown(nfa.respondingSpecialChar, make(map[*State]bool))
	return nfa
}
func (nfa *NFA) FormMermaid(file *os.File) {
	ids := make(map[*State]int)
	lines := new(int)
	result := make([]string, 0)
	nfa.getStartState().GetShowDataFromHere(0, ids, make(map[*State]bool), lines, &result)
	_, err := file.WriteString("```mermaid\ngraph LR\n")
	for i := 0; i < len(result); i++ {
		_, err = file.WriteString(result[i])
		if err != nil {
			panic(err)
		}
	}
	_, err = file.WriteString("```\n")
	if err != nil {
		panic(err)
	}
}
func (nfa *NFA) FormTheMermaidGraphOfNFA(filePath string) {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	nfa.EliminateBlankStates()
	nfa.FormMermaid(file)
}

func (nfa *NFA) GetWordEndPair() []*WordEndPair {
	return nfa.startState.GetWordEndPair("", make(map[*State]bool))
}

func (nfa *NFA) GetStartState() *State {
	return nfa.startState
}


