package syntex

import (
	"conf"
	"fmt"
	"lex"
)


func (sp *SyntaxParser)IsValid(text []byte) (result bool){
	defer func() {
		result = recover()==nil
	}()

	sp.GetSyntaxTree(text)
	return
}

// TODO: 这里只完成了语法分析，还没获得语法树
func (sp *SyntaxParser)GetSyntaxTree(text []byte) {
	sp.initGetSyntaxTree(text)
	for sp.readingIsNotOver(){
		sp.execGetSyntaxTree()
	}
}

func (sp *SyntaxParser)initGetSyntaxTree(text []byte){
	sp.readingPosition = 0
	sp.symbolsStack = nil
	sp.terminatorPairs = sp.lexicalAnalyzer.GetTerminatorPairs(text)

	sp.terminatorPairs = append(sp.terminatorPairs,lex.NewTerminatorPair(conf.GetConf().SyntaxConf.EndSymbol,nil))
	sp.symbolsStack = append(sp.symbolsStack, conf.GetConf().SyntaxConf.EndSymbol)
	sp.symbolsStack = append(sp.symbolsStack, conf.GetConf().SyntaxConf.StartSymbol)
}
func (sp *SyntaxParser)readingIsNotOver() bool{
	return sp.readingPosition!=len(sp.terminatorPairs)
}
func (sp *SyntaxParser)execGetSyntaxTree() {
	//sp.showParsingMiddleWares()
	switch  {
	case sp.symbolOfStackTopIsBlank():
		sp.symbolStackPop()
	case sp.canOffset():
		sp.offset()
	case sp.canContinueParsing():
		sp.continueParsing()
	default:
		sp.error()
	}
}


func (sp *SyntaxParser)showParsingMiddleWares() {
	fmt.Println(sp.symbolsStack,"|",terminatorPairsToString(sp.getNotReadSymbolPairs()))
}
func (sp *SyntaxParser)canOffset() bool{
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	return stackTopSymbol == readingSymbol
}
func (sp *SyntaxParser)offset() {
	sp.symbolStackPop()
	sp.readNextSymbolPair()
}
func(sp *SyntaxParser)canContinueParsing() bool{
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	return sp.stateTable.HasSentence(stackTopSymbol,readingSymbol)
}
func (sp *SyntaxParser)continueParsing() {
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	sentence := sp.stateTable.GetSentence(stackTopSymbol,readingSymbol)
	sp.symbolStackPop()
	sp.reversePushTheSymbolOfSentenceIntoSymbolStack(sentence)
}
func (sp *SyntaxParser)error() {
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	panic(fmt.Sprintf("出错，%s %s", stackTopSymbol ,readingSymbol))
}

func (sp *SyntaxParser)getNotReadSymbolPairs() []*lex.TerminatorPair{
	return sp.terminatorPairs[sp.readingPosition:]
}
func (sp *SyntaxParser)readNextSymbolPair() {
	sp.readingPosition++
}
func (sp *SyntaxParser)symbolStackIsEmpty() bool{
	return len(sp.symbolsStack)==0
}
func (sp *SyntaxParser) getSymbolOfReadingSymbolPair() string{
	return sp.terminatorPairs[sp.readingPosition].GetSymbol()
}
func (sp *SyntaxParser) symbolStackPop() {
	sp.symbolsStack = sp.symbolsStack[:len(sp.symbolsStack)-1]
}
func (sp *SyntaxParser) getSymbolOfSymbolStack() string{
	return sp.symbolsStack[len(sp.symbolsStack)-1]
}
func (sp *SyntaxParser) symbolOfStackTopIsBlank() bool{
	return sp.getSymbolOfSymbolStack()== conf.GetConf().SyntaxConf.BlankSymbol
}
func (sp *SyntaxParser)reversePushTheSymbolOfSentenceIntoSymbolStack(sentence *sentence) {
	for i:=len(sentence.symbols)-1;i>=0;i-- {
		sp.symbolsStack = append(sp.symbolsStack,sentence.symbols[i])
	}
}

func terminatorPairsToString(terminatorPairs []*lex.TerminatorPair) string {
	result := ""
	for i := 0; i < len(terminatorPairs); i++ {
		result += fmt.Sprintf("%v ",terminatorPairs[i].GetValue())
	}
	return result
}





