package grammarTree

import (
	"file"
	"fmt"
	"lexical"
	"stateMachine"
)

type GrammarTree struct {
	table *llTable
	lexicalAnalyzer *lexical.LexicalAnalyzer
}

func NewGrammarTree(lexicalAnalyzer *lexical.LexicalAnalyzer) *GrammarTree{
	return &GrammarTree{NewLLTable(),lexicalAnalyzer}
}
func showTokens(tokens []*stateMachine.Token) string{
	str := ""
	for i:=0;i<len(tokens);i++{
		str += tokens[i].GetValue().(string)+" "
	}
	return str
}

func (gt *GrammarTree) Do(path string) {
	tokens := gt.lexicalAnalyzer.GetTokens(file.NewFileReader(path).GetFileBytes())
	gt.handle(tokens)
}

func (gt *GrammarTree)handle(tokens []*stateMachine.Token) {
	stack := make([]string,0)
	stack = append(stack,"END")
	stack = append(stack,"EXP")
	tokens = append(tokens,stateMachine.NewEmptyToken())
	readingPosition := 0
	for readingPosition!=len(tokens){
		fmt.Println(stack,showTokens(tokens[readingPosition:]))
		symbol := transfer(tokens[readingPosition].GetSpecialChar())
		topIndex := len(stack)-1

		if stack[topIndex]=="BLA"{
			stack = stack[:topIndex]
			continue
		}

		if symbol==stack[topIndex]{
			stack = stack[:topIndex]
			readingPosition++
		}else{
			result := gt.table.Get(stack[topIndex],symbol)
			if len(result)==0{
				panic(fmt.Sprintf("出错，%s %s %v",stack[topIndex],symbol,result))
			}
			stack = stack[:topIndex]
			for i:=len(result)-1;i>=0;i--{
				stack = append(stack, result[i])
			}
		}
	}
}

func transfer(char byte) string{
	switch char {
	case 'I':
		return "IDE"
	case 'A':
		return "ADD"
	case 'F':
		return "FOPT"
	case 'Y':
		return "LEF"
	case 'U':
		return "RIT"
	case 0:
		return "END"
	}
	panic("error!!!!!!!!!!")
}
