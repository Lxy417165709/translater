package syntex

import (
	"bytes"
	"fmt"
)



func (stf *StateTableFormer)Show(){
	lines := stf.getStateTableLines()
	for _,line := range lines{
		fmt.Print(line)
	}
}
func (stf *StateTableFormer) getStateTableLines() []string{
	result := make([]string,0)
	result = append(result,stf.getStateTableFirstLine())
	result = append(result,stf.getStateTableOtherLines()...)
	return result
}
func (stf *StateTableFormer)getStateTableFirstLine() string{
	firstLineBuffer := bytes.Buffer{}
	firstLineBuffer.WriteString("非终结符")
	for _,terminator := range stf.terminators{
		firstLineBuffer.WriteString("|"+terminator)
	}
	firstLineBuffer.WriteString("\n--")
	for i:=0;i<len(stf.terminators);i++{
		firstLineBuffer.WriteString("|--")
	}
	firstLineBuffer.WriteString("\n")
	return firstLineBuffer.String()
}
func (stf *StateTableFormer)getStateTableOtherLines() []string{
	result := make([]string,0)
	for index := range stf.getNonTerminators(){
		result = append(result,stf.getNthNonTerminatorLines(index))
	}
	return result
}
func (stf *StateTableFormer) getNthNonTerminatorLines(index int) string{
	lineBuffer := bytes.Buffer{}
	nonTerminator := stf.getNonTerminators()[index]
	lineBuffer.WriteString(nonTerminator)
	for _,terminator := range stf.terminators{
		lineBuffer.WriteString("|")
		if stf.stateTable[nonTerminator]!=nil && stf.stateTable[nonTerminator][terminator]!=nil{
			lineBuffer.WriteString(fmt.Sprintf("%v",*stf.stateTable[nonTerminator][terminator]))
		}
	}
	lineBuffer.WriteString("\n")
	return lineBuffer.String()
}


