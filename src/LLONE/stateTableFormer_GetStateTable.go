package LLONE

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

func (stf *StateTableFormer) FormMarkDownStateTable() string{

	tableBuffer := bytes.Buffer{}
	tableBuffer.WriteString("非终结符")
	for _,terminator := range terminators{
		tableBuffer.WriteString("|"+terminator)
	}
	tableBuffer.WriteString("\n--")
	for i:=0;i<len(terminators);i++{
		tableBuffer.WriteString("|--")
	}
	tableBuffer.WriteString("\n")

	nonTerminators := stf.GetNonTerminators()
	for _,nonTerminator := range nonTerminators{
		tableBuffer.WriteString(nonTerminator)
		for _,terminator := range terminators{
			tableBuffer.WriteString("|")
			if stf.StateTable[nonTerminator]!=nil && stf.StateTable[nonTerminator][terminator]!=nil{
				tableBuffer.WriteString(fmt.Sprintf("%v",*stf.StateTable[nonTerminator][terminator]))
			}
		}
		tableBuffer.WriteString("\n")
	}
	return tableBuffer.String()
}


func (stf *StateTableFormer)GetNonTerminators() []string{
	result := make([]string,0)
	for _,production := range stf.productions{
		result = append(result,production.leftNonTerminator)
	}
	return result
}
