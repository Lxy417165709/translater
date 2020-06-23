package tb

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
	terminators := stf.llOneTable.GetTerminators()
	for _, terminator := range terminators {
		firstLineBuffer.WriteString("|" + terminator)
	}
	firstLineBuffer.WriteString("\n--")
	for i := 0; i < len(terminators); i++ {
		firstLineBuffer.WriteString("|--")
	}
	firstLineBuffer.WriteString("\n")
	return firstLineBuffer.String()
}
func (stf *StateTable) getStateTableOtherLines() []string {
	result := make([]string, 0)
	nonTerminators := stf.llOneTable.GetNonTerminators()
	for index := range nonTerminators {
		result = append(result, stf.getNthNonTerminatorLines(index))
	}
	return result
}
func (stf *StateTable) getNthNonTerminatorLines(index int) string {
	lineBuffer := bytes.Buffer{}
	nonTerminators := stf.llOneTable.GetNonTerminators()
	terminators := stf.llOneTable.GetTerminators()

	nonTerminator := nonTerminators[index]
	lineBuffer.WriteString(nonTerminator)
	for _, terminator := range terminators {
		lineBuffer.WriteString("|")
		if stf.dataMatrix[nonTerminator] != nil && stf.dataMatrix[nonTerminator][terminator] != nil {
			lineBuffer.WriteString(fmt.Sprintf("%v", *stf.dataMatrix[nonTerminator][terminator]))
		}
	}
	lineBuffer.WriteString("\n")
	return lineBuffer.String()
}





