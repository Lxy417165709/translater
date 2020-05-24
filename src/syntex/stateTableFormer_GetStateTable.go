package syntex

import (
	"bytes"
	"fmt"
)



func (stf *StateTableFormer)GetStateTable() {
	stf.StateTable = make(map[string]map[string]*sentence)
	for sentenc,terminators := range stf.Select{
		for _,terminator := range terminators{
			if stf.StateTable[stf.SentenceToNonTerminator[sentenc]]==nil{
				stf.StateTable[stf.SentenceToNonTerminator[sentenc]] = make(map[string]*sentence)
			}
			stf.StateTable[stf.SentenceToNonTerminator[sentenc]][terminator]=sentenc
		}
	}
}
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
	for index := range stf.GetNonTerminators(){
		result = append(result,stf.getNthNonTerminatorLines(index))
	}
	return result
}
func (stf *StateTableFormer) getNthNonTerminatorLines(nth int) string{
	lineBuffer := bytes.Buffer{}
	nonTerminator := stf.GetNonTerminators()[nth]
	lineBuffer.WriteString(nonTerminator)
	for _,terminator := range stf.terminators{
		lineBuffer.WriteString("|")
		if stf.StateTable[nonTerminator]!=nil && stf.StateTable[nonTerminator][terminator]!=nil{
			lineBuffer.WriteString(fmt.Sprintf("%v",*stf.StateTable[nonTerminator][terminator]))
		}
	}
	lineBuffer.WriteString("\n")
	return lineBuffer.String()
}


func (stf *StateTableFormer)GetNonTerminators() []string{
	result := make([]string,0)
	for _,production := range stf.productions{
		result = append(result,production.leftNonTerminator)
	}
	return result
}
