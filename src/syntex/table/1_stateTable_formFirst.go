package table

import (
	"conf"
)

func (stf *StateTable) formFirst() {
	stf.templateFunctionOfForming(
		stf.initGetFirst,
		stf.handleGettingFirst,
		stf.syncBufferOfFirst,
	)
}

func (stf *StateTable) initGetFirst() {
	stf.first = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
	stf.initProductions()
}

func (stf *StateTable) initProductions()  {
	for _, originProduction := range getProductions(conf.GetConf().SyntaxConf.SyntaxFilePath) {
		stf.productions = append(stf.productions, originProduction.ChangeToNonLeftRecursionProductions()...)
	}
	// 输出无左递归的产生式
	//for i:=0;i<len(stf.productions);i++{
	//	fmt.Println(stf.productions[i].nonTerminator,"->")
	//	for t:=0;t<len(stf.productions[i].sentences);t++{
	//		fmt.Printf("\t")
	//		fmt.Println(stf.productions[i].sentences[t].symbols)
	//	}
	//}

}


func (stf *StateTable) handleGettingFirst() {
	handlingProduction := stf.productions[stf.indexOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.indexOfHandlingProductionSentence]
	if handlingSentence.IsBlank() {
		stf.handleGettingFirstOfSentenceIsBlank()
	} else {
		stf.handleGettingFirstOfSentenceIsNotBlank()
	}
}
func (stf *StateTable) syncBufferOfFirst() bool {
	firstSetHasBeenUpdated := false
	for nonTerminator,sentence:= range stf.bufferOfSet{
		for _,symbol := range sentence{
			if !stf.terminatorIsLivingInFirst(nonTerminator, symbol) {
				stf.first[nonTerminator] = append(stf.first[nonTerminator], symbol)
				firstSetHasBeenUpdated= true
			}
		}
	}
	stf.flushBufferOfSet()
	return firstSetHasBeenUpdated
}

func (stf *StateTable) handleGettingFirstOfSentenceIsBlank() {
	handlingProduction := stf.productions[stf.indexOfHandlingProduction]
	blankSymbol := conf.GetConf().SyntaxConf.BlankSymbol
	nonTerminator := handlingProduction.nonTerminator
	stf.appendToBufferOfSet(nonTerminator,blankSymbol)
}
func (stf *StateTable) handleGettingFirstOfSentenceIsNotBlank() {
	handlingProduction := stf.productions[stf.indexOfHandlingProduction]
	handlingSymbols := stf.getHandlingSymbols()
	for index, symbol := range handlingSymbols {
		nonTerminator := handlingProduction.nonTerminator
		switch {
		case stf.isTerminator(symbol):
			stf.appendToBufferOfSet(nonTerminator,symbol)
			return
		case !hasBlankSymbol(stf.first[symbol]):
			stf.appendToBufferOfSet(nonTerminator,removeBlankSymbol(stf.first[symbol])...)
			return
		case hasBlankSymbol(stf.first[symbol]):
			isHandlingLastSymbol := index == len(handlingSymbols)-1
			if isHandlingLastSymbol {
				stf.appendToBufferOfSet(nonTerminator,stf.first[symbol]...)
				continue
			}
			stf.appendToBufferOfSet(nonTerminator,removeBlankSymbol(stf.first[symbol])...)
		default:
			stf.error()
		}
	}
}

func (stf *StateTable)getHandlingSymbols() []string{
	return stf.getHandlingSentence().symbols
}
func (stf *StateTable)getHandlingSentence() *Sentence{
	handlingProduction := stf.productions[stf.indexOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.indexOfHandlingProductionSentence]
	return handlingSentence
}

// TODO: 这个命名不太准确
func (stf *StateTable) terminatorIsLivingInFirst(nonTerminator string, terminator string) bool {
	return arrayHasTerminator(stf.first[nonTerminator],terminator)
}



func (stf *StateTable)error() {
	panic("存在没有考虑的情况")
}














