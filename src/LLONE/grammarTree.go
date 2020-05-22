package LLONE

import (
	"file"
	"fmt"
	"lexical"
	"stateMachine"
)


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
	return gt.tokens[gt.readingPosition].GetType()
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
			gt.readNextToken()
		} else {
			if !gt.stateTable.HasSentence(gt.getSymbolOfSymbolStackTop(),gt.getSymbolOfReadingToken()){
				panic(fmt.Sprintf("出错，%s %s", gt.getSymbolOfSymbolStackTop(),gt.getSymbolOfReadingToken()))
			}
			sentence := gt.stateTable.GetSentence(gt.getSymbolOfSymbolStackTop(),gt.getSymbolOfReadingToken())
			gt.popSymbolStackTop()
			gt.ReversePushSentenceIntoSymbolStack(sentence)
		}
	}
}



func (gt *GrammarTreeBuilder) readNextToken() {
	gt.readingPosition++
}
func showTokens(tokens []*stateMachine.Token) string {
	str := ""
	for i := 0; i < len(tokens); i++ {
		str += tokens[i].GetType() + " "
	}
	return str
}
