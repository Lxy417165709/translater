package syntex

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

}


func (stf *StateTable) handleGettingFirst() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]
	if handlingSentence.IsBlank() {
		stf.handleGettingFirstOfSentenceIsBlank()
	} else {
		stf.handleGettingFirstOfSentenceIsNotBlank()
	}
}
func (stf *StateTable) syncBufferOfFirst() bool {
	firstSetHasBeenUpdated := false
	for leftNonTerminator,sentence:= range stf.bufferOfSet{
		for _,symbol := range sentence{
			if !stf.terminatorIsLivingInfirst(leftNonTerminator, symbol) {
				stf.first[leftNonTerminator] = append(stf.first[leftNonTerminator], symbol)
				firstSetHasBeenUpdated= true
			}
		}
	}
	stf.flushBufferOfSet()
	return firstSetHasBeenUpdated
}

func (stf *StateTable) handleGettingFirstOfSentenceIsBlank() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,conf.GetConf().SyntaxConf.BlankSymbol)
}
func (stf *StateTable) handleGettingFirstOfSentenceIsNotBlank() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSymbols := stf.getHandlingSymbols()
	for index, symbol := range handlingSymbols {
		switch {
		case stf.isTerminator(symbol):
			stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,symbol)
			return
		case !hasBlankSymbol(stf.first[symbol]):
			stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,removeBlankSymbol(stf.first[symbol])...)
			return
		case hasBlankSymbol(stf.first[symbol]):
			isHandlingLastSymbol := index == len(handlingSymbols)-1
			if isHandlingLastSymbol {
				stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,stf.first[symbol]...)
				continue
			}
			stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,removeBlankSymbol(stf.first[symbol])...)
		default:
			stf.error()
		}
	}
}

func (stf *StateTable)getHandlingSymbols() []string{
	return stf.getHandlingSentence().symbols
}
func (stf *StateTable)getHandlingSentence() *sentence{
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]
	return handlingSentence
}

func (stf *StateTable) terminatorIsLivingInfirst(leftNonTerminator string, terminator string) bool {
	return arrayHasTerminator(stf.first[leftNonTerminator],terminator)
}



func (stf *StateTable)error() {
	panic("存在没有考虑的情况")
}














