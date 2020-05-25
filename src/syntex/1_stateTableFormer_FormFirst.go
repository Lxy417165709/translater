package syntex

import "conf"

func (stf *StateTableFormer) FormFirst() {
	stf.TemplateFunctionOfForming(
		stf.initGetFirst,
		stf.handleGettingFirst,
		stf.syncBufferOfFirst,
	)
}

func (stf *StateTableFormer) initGetFirst() {
	stf.first = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
}
func (stf *StateTableFormer) handleGettingFirst() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]
	if stf.sentenceIsBlank(handlingSentence) {
		stf.handleGettingFirstOfSentenceIsBlank()
	} else {
		stf.handleGettingFirstOfSentenceIsNotBlank()
	}
}
func (stf *StateTableFormer) syncBufferOfFirst() bool {
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




func (stf *StateTableFormer) handleGettingFirstOfSentenceIsBlank() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,conf.GetConf().SyntaxConf.BlankSymbol)
}
func (stf *StateTableFormer) handleGettingFirstOfSentenceIsNotBlank() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSymbols := stf.getHandlingSymbols()
	for index, symbol := range handlingSymbols {
		switch {
		case stf.isTerminator(symbol):
			stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,symbol)
			return
		case !stf.hasBlankSymbol(stf.first[symbol]):
			stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,stf.removeBlankSymbol(stf.first[symbol])...)
			return
		case stf.hasBlankSymbol(stf.first[symbol]):
			isHandlingLastSymbol := index == len(handlingSymbols)-1
			if isHandlingLastSymbol {
				stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,stf.first[symbol]...)
				continue
			}
			stf.appendToBufferOfSet(handlingProduction.leftNonTerminator,stf.removeBlankSymbol(stf.first[symbol])...)
		default:
			stf.error()
		}
	}
}

func (stf *StateTableFormer)getHandlingSymbols() []string{
	return stf.getHandlingSentence().symbols
}
func (stf *StateTableFormer)getHandlingSentence() *sentence{
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]
	return handlingSentence
}

func (stf *StateTableFormer) terminatorIsLivingInfirst(leftNonTerminator string, terminator string) bool {
	return arrayHasTerminator(stf.first[leftNonTerminator],terminator)
}



func (stf *StateTableFormer)error() {
	panic("存在没有考虑的情况")
}














