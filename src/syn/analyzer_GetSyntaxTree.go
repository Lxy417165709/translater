package syn

import (
	"dto"
	"fmt"
)

func (az *Analyzer) GetSyntaxTree(tokens []*dto.Token) error {
	az.initGetSyntaxTree(tokens)
	for az.readingIsNotOver() {
		if err := az.execGetSyntaxTree(); err != nil {
			return err
		}
	}
	return nil
}

func (az *Analyzer) initGetSyntaxTree(tokens []*dto.Token) {
	az.tokens = append([]*dto.Token{}, tokens...)
	az.tokens = append(az.tokens, dto.NewEmptyToken(az.conf.StateTableConf.LLOneTableConf.EndSymbol))

	az.readingPosition = 0
	az.treeNodeStack = nil
	az.syntaxTreeRoot = NewTreeNode(
		dto.NewEmptyToken(az.conf.StateTableConf.LLOneTableConf.StartSymbol),
	)

	az.treeNodeStack = append(az.treeNodeStack, NewTreeNode(
		dto.NewEmptyToken(az.conf.StateTableConf.LLOneTableConf.EndSymbol),
	))
	az.treeNodeStack = append(az.treeNodeStack, az.syntaxTreeRoot)
}
func (az *Analyzer) readingIsNotOver() bool {
	return az.readingPosition != len(az.tokens)
}
func (az *Analyzer) execGetSyntaxTree() error {
	//az.showParsingMiddleWares()
	switch {
	case az.symbolOfStackTopIsBlank():
		az.symbolStackPop()
	case az.canOffset():
		az.offset()
	case az.canContinueParsing():
		az.continueParsing()
	default:
		return fmt.Errorf("语法有误，请检查代码")
	}
	return nil
}

func (az *Analyzer) showParsingMiddleWares() {

	for i := 0; i < len(az.treeNodeStack); i++ {
		fmt.Printf("%s ", az.treeNodeStack[i].token.Symbol)
	}

	fmt.Println("|", tokensToString(az.getNotReadSymbolTokens()))
}
func (az *Analyzer) canOffset() bool {
	stackTopSymbol := az.getSymbolOfSymbolStack()
	readingSymbol := az.getSymbolOfReadingSymbolToken()
	return stackTopSymbol == readingSymbol
}
func (az *Analyzer) offset() {
	topNode := az.getTopNode()
	topNode.token = az.tokens[az.readingPosition]
	az.symbolStackPop()
	az.readNextSymbolPair()
}
func (az *Analyzer) canContinueParsing() bool {
	stackTopSymbol := az.getSymbolOfSymbolStack()
	readingSymbol := az.getSymbolOfReadingSymbolToken()
	return az.stateTable.HasSentence(stackTopSymbol, readingSymbol)
}
func (az *Analyzer) continueParsing() {
	stackTopSymbol := az.getSymbolOfSymbolStack()
	readingSymbol := az.getSymbolOfReadingSymbolToken()
	sentence := az.stateTable.GetSentence(stackTopSymbol, readingSymbol)

	topNode := az.getTopNode()
	topNode.FormSon(sentence)
	az.symbolStackPop()
	az.appendSonToStack(topNode)

}

func (az *Analyzer) appendSonToStack(topNode *TreeNode) {
	for i := len(topNode.son) - 1; i >= 0; i-- {
		az.treeNodeStack = append(az.treeNodeStack, topNode.son[i])
	}
}

func (az *Analyzer) getTopNode() *TreeNode {
	return az.treeNodeStack[len(az.treeNodeStack)-1]
}


func (az *Analyzer) getNotReadSymbolTokens() []*dto.Token {
	return az.tokens[az.readingPosition:]
}
func (az *Analyzer) readNextSymbolPair() {
	az.readingPosition++
}

// TODO: 命名
func (az *Analyzer) symbolStackIsEmpty() bool {
	return len(az.treeNodeStack) == 0
}
func (az *Analyzer) getSymbolOfReadingSymbolToken() string {
	return az.tokens[az.readingPosition].Symbol
}

// TODO: 命名
func (az *Analyzer) symbolStackPop() {
	az.treeNodeStack = az.treeNodeStack[:len(az.treeNodeStack)-1]
}

// TODO: 命名
func (az *Analyzer) getSymbolOfSymbolStack() string {
	return az.treeNodeStack[len(az.treeNodeStack)-1].token.Symbol
}
func (az *Analyzer) symbolOfStackTopIsBlank() bool {
	return az.getSymbolOfSymbolStack() == az.conf.StateTableConf.LLOneTableConf.BlankSymbol
}

func tokensToString(tokens []*dto.Token) string {
	result := ""
	for i := 0; i < len(tokens); i++ {
		result += fmt.Sprintf("%v ", tokens[i].Symbol)
	}
	return result
}
