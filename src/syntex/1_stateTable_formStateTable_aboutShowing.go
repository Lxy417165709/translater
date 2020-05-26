package syntex

import (
	"bytes"
	"fmt"
)

func (stf *StateTable) Show() {
	lines := stf.getStateTableLines()
	for _, line := range lines {
		fmt.Print(line)
	}
}
func (stf *StateTable) getStateTableLines() []string {
	result := make([]string, 0)
	result = append(result, stf.getStateTableBeginLines())
	result = append(result, stf.getStateTableOtherLines()...)
	return result
}
func (stf *StateTable) getStateTableBeginLines() string {
	firstLineBuffer := bytes.Buffer{}
	firstLineBuffer.WriteString("非终结符")
	for _, terminator := range stf.terminators {
		firstLineBuffer.WriteString("|" + terminator)
	}
	firstLineBuffer.WriteString("\n--")
	for i := 0; i < len(stf.terminators); i++ {
		firstLineBuffer.WriteString("|--")
	}
	firstLineBuffer.WriteString("\n")
	return firstLineBuffer.String()
}
func (stf *StateTable) getStateTableOtherLines() []string {
	result := make([]string, 0)
	for index := range stf.getNonTerminators() {
		result = append(result, stf.getNthNonTerminatorLines(index))
	}
	return result
}
func (stf *StateTable) getNthNonTerminatorLines(index int) string {
	lineBuffer := bytes.Buffer{}
	nonTerminator := stf.getNonTerminators()[index]
	lineBuffer.WriteString(nonTerminator)
	for _, terminator := range stf.terminators {
		lineBuffer.WriteString("|")
		if stf.table[nonTerminator] != nil && stf.table[nonTerminator][terminator] != nil {
			lineBuffer.WriteString(fmt.Sprintf("%v", *stf.table[nonTerminator][terminator]))
		}
	}
	lineBuffer.WriteString("\n")
	return lineBuffer.String()
}
