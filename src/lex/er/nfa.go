package er

import (
	"fmt"
	"os"
)

type NFA struct {
	startState *State
	endState   *State

	innerStates []*State
	specialChar byte
}

func NewNFA(ordinaryChar byte) *NFA {
	startState:=NewState(false)
	endState:=NewState(true)
	startState.next[ordinaryChar] = append(startState.next[ordinaryChar],endState)
	return &NFA{
		startState:startState,
		endState:endState,
		innerStates:[]*State{startState},
	}
}
func NewEmptyNFA() *NFA{
	startState := NewState(false)
	return &NFA{
		startState:startState,
		endState:NewState(true),
		innerStates:[]*State{startState},
	}
}

func (nfa *NFA) SetSpecialChar(specialChar byte) {
	nfa.specialChar = specialChar
}
func (nfa *NFA) GetSpecialChar() byte{
	return nfa.specialChar
}

func (nfa *NFA) IsMatch(pattern string) bool {
	return nfa.startState.IsMatchFromHere(pattern)
}
func (nfa *NFA) EliminateBlankStates() *NFA {
	hasVisited := make(map[*State]bool)
	nfa.startState.EliminateNextBlankStatesFromHere(hasVisited)
	return nfa
}
func (nfa *NFA) StoreMermaidGraphOfThisNFA(filePath string) error {
	var file *os.File
	var err error
	if file, err = os.Create(filePath); err != nil {
		return fmt.Errorf("%s 路径，文件创建失败", filePath)
	}
	defer file.Close()
	for _, line := range nfa.getMermaidLines() {
		if _, err = file.WriteString(line); err != nil {
			return fmt.Errorf("%s 路径，向文件中写入内容(%s)失败", filePath, line)
		}
	}
	return err
}

func (nfa *NFA) getMermaidLines() []string {
	lines := make([]string, 0)
	lines = append(lines, "```mermaid\ngraph LR\n")
	lines = append(lines, nfa.getMetaMermaidData()...)
	lines = append(lines, "```\n")
	return lines
}
func (nfa *NFA) getMetaMermaidData() []string {
	return nfa.startState.GetLinesOfLinkInformationFromHere(0, make(map[*State]int), make(map[*State]bool))
}
