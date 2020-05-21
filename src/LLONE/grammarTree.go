package LLONE

import (
	"file"
	"fmt"
	"lexical"
	"stateMachine"
)

const llOneFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\LL1`

type GrammarTreeBuilder struct {
	stateTable      *StateTableFormer
	lexicalAnalyzer *lexical.LexicalAnalyzer

	symbolsStack []string
	tokens []*stateMachine.Token
	readingPosition int
}

func NewGrammarTree(stateTable *StateTableFormer, lexicalAnalyzer *lexical.LexicalAnalyzer) *GrammarTreeBuilder {
	return &GrammarTreeBuilder{stateTable: stateTable, lexicalAnalyzer: lexicalAnalyzer}
}

func (gt *GrammarTreeBuilder) initSymbolsStack() {
	gt.symbolsStack = make([]string,0)
	gt.symbolsStack = append(gt.symbolsStack, "END")
	gt.symbolsStack = append(gt.symbolsStack, "EXP")
}
func (gt *GrammarTreeBuilder) initTokens() {
	gt.tokens = append(gt.tokens, stateMachine.NewEmptyToken())
}
func (gt *GrammarTreeBuilder) initReadingPosition() {
	gt.readingPosition=0
}
func (gt *GrammarTreeBuilder) readingIsNotOver()bool {
	return gt.readingPosition != len(gt.tokens)
}
func (gt *GrammarTreeBuilder) getSymbolOfReadingToken() string{
	return transfer(gt.tokens[gt.readingPosition].GetSpecialChar())
}
func (gt *GrammarTreeBuilder) getSymbolOfSymbolStackTop() string{
	return gt.symbolsStack[len(gt.symbolsStack)-1]
}
func (gt *GrammarTreeBuilder) popSymbolStackTop() {
	gt.symbolsStack = gt.symbolsStack[:len(gt.symbolsStack)-1]
}

func (gt *GrammarTreeBuilder)ReversePushSentenceIntoSymbolStack(sentence *sentence) {
	for i:=len(sentence.symbols)-1;i>=0;i-- {
		gt.symbolsStack = append(gt.symbolsStack,sentence.symbols[i])
	}
}
func (gt *GrammarTreeBuilder) Do(path string) {
	tokens := gt.lexicalAnalyzer.GetTokens(file.NewFileReader(path).GetFileBytes())
	gt.handle(tokens)
}

func (gt *GrammarTreeBuilder) handle(tokens []*stateMachine.Token) {

	gt.initSymbolsStack()
	gt.tokens = append(gt.tokens,tokens...)
	gt.initTokens()
	gt.initReadingPosition()

	for gt.readingIsNotOver() {
		if gt.getSymbolOfSymbolStackTop() == "BLA" {
			gt.popSymbolStackTop()
			continue
		}
		fmt.Println(gt.symbolsStack,showTokens(gt.tokens[gt.readingPosition:]))
		if gt.getSymbolOfReadingToken() == gt.getSymbolOfSymbolStackTop(){
			gt.popSymbolStackTop()
			gt.readingPosition++
		} else {
			sentence := gt.stateTable.GetSentence(gt.getSymbolOfSymbolStackTop(),gt.getSymbolOfReadingToken())
			if sentence==nil {
				panic(fmt.Sprintf("出错，%s %s", gt.getSymbolOfSymbolStackTop(),gt.getSymbolOfReadingToken()))
			}
			gt.popSymbolStackTop()
			gt.ReversePushSentenceIntoSymbolStack(sentence)
		}
	}
}

func transfer(char byte) string {
	switch char {
	case 'I':
		return "IDE"
	case 'A':
		return "ASO"
	case 'F':
		return "FDO"
	case 'Y':
		return "LEFT_PAR"
	case 'U':
		return "RIGHT_PAR"
	case 'Z':
		return "ZS"
	case 0:
		return "END"
	}
	panic("error!!!!!!!!!!")
}
func showTokens(tokens []*stateMachine.Token) string {
	str := ""
	for i := 0; i < len(tokens); i++ {
		str += tokens[i].GetValue().(string) + " "
	}
	return str
}
