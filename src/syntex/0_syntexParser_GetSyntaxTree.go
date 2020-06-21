package syntex

import (
	"conf"
	"fmt"
	"lex/terminator"
)

func (sp *SyntaxParser) IsValid(text []byte) (result bool) {
	defer func() {
		result = recover() == nil
	}()
	sp.GetSyntaxTree(text)
	return
}

func (sp *SyntaxParser) GetSyntaxTree(text []byte) {
	sp.initGetSyntaxTree(text)
	for sp.readingIsNotOver() {
		sp.execGetSyntaxTree()
	}
}

func (sp *SyntaxParser) initGetSyntaxTree(text []byte) {
	sp.readingPosition = 0
	sp.treeNodeStack = nil
	sp.syntaxTreeRoot = NewTreeNode(conf.GetConf().SyntaxConf.StartSymbol)
	sp.terminatorPairs = sp.lexicalAnalyzer.GetPairs(text)

	sp.terminatorPairs = append(sp.terminatorPairs, terminator.NewPair(conf.GetConf().SyntaxConf.EndSymbol, nil))
	sp.treeNodeStack = append(sp.treeNodeStack, NewTreeNode(conf.GetConf().SyntaxConf.EndSymbol))
	sp.treeNodeStack = append(sp.treeNodeStack, sp.syntaxTreeRoot)
}
func (sp *SyntaxParser) readingIsNotOver() bool {
	return sp.readingPosition != len(sp.terminatorPairs)
}
func (sp *SyntaxParser) execGetSyntaxTree() {
	sp.showParsingMiddleWares()
	switch {
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

func (sp *SyntaxParser) showParsingMiddleWares() {

	for i:=0;i<len(sp.treeNodeStack);i++{
		fmt.Printf("%s ",sp.treeNodeStack[i].symbol)
	}

	fmt.Println( "|", terminatorPairsToString(sp.getNotReadSymbolPairs()))
}
func (sp *SyntaxParser) canOffset() bool {
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	return stackTopSymbol == readingSymbol
}
func (sp *SyntaxParser) offset() {
	topNode := sp.getTopNode()
	topNode.pair = sp.terminatorPairs[sp.readingPosition]
	sp.symbolStackPop()
	sp.readNextSymbolPair()
}
func (sp *SyntaxParser) canContinueParsing() bool {
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	return sp.stateTable.HasSentence(stackTopSymbol, readingSymbol)
}
func (sp *SyntaxParser) continueParsing() {
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	sentence := sp.stateTable.GetSentence(stackTopSymbol, readingSymbol)


	topNode := sp.getTopNode()
	topNode.FormSon(sentence)
	sp.symbolStackPop()
	sp.appendSonToStack(topNode)

}

func (sp *SyntaxParser) appendSonToStack(topNode *TreeNode) {
	for i := len(topNode.son)-1; i >=0; i-- {
		sp.treeNodeStack = append(sp.treeNodeStack, topNode.son[i])
	}
}

func (sp *SyntaxParser) getTopNode() *TreeNode {
	return sp.treeNodeStack[len(sp.treeNodeStack)-1]
}

func (sp *SyntaxParser) error() {
	stackTopSymbol := sp.getSymbolOfSymbolStack()
	readingSymbol := sp.getSymbolOfReadingSymbolPair()
	panic(fmt.Sprintf("出错，%s %s", stackTopSymbol, readingSymbol))
}

func (sp *SyntaxParser) getNotReadSymbolPairs() []*terminator.Pair {
	return sp.terminatorPairs[sp.readingPosition:]
}
func (sp *SyntaxParser) readNextSymbolPair() {
	sp.readingPosition++
}

// TODO: 命名
func (sp *SyntaxParser) symbolStackIsEmpty() bool {
	return len(sp.treeNodeStack) == 0
}
func (sp *SyntaxParser) getSymbolOfReadingSymbolPair() string {
	return sp.terminatorPairs[sp.readingPosition].GetSymbol()
}

// TODO: 命名
func (sp *SyntaxParser) symbolStackPop() {
	sp.treeNodeStack = sp.treeNodeStack[:len(sp.treeNodeStack)-1]
}

// TODO: 命名
func (sp *SyntaxParser) getSymbolOfSymbolStack() string {
	return sp.treeNodeStack[len(sp.treeNodeStack)-1].symbol
}
func (sp *SyntaxParser) symbolOfStackTopIsBlank() bool {
	return sp.getSymbolOfSymbolStack() == conf.GetConf().SyntaxConf.BlankSymbol
}

func terminatorPairsToString(terminatorPairs []*terminator.Pair) string {
	result := ""
	for i := 0; i < len(terminatorPairs); i++ {
		result += fmt.Sprintf("%v ", terminatorPairs[i].GetValue())
	}
	return result
}
