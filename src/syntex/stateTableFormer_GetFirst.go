package syntex



func (stf *StateTableFormer) terminatorIsLivingInFirst(leftNonTerminator string, terminator string) bool {
	return arrayHasTerminator(stf.First[leftNonTerminator],terminator)
}
func (stf *StateTableFormer) syncBufferOfFirst() bool {
	firstSetHasBeenUpdated := false
	for leftNonTerminator,sentence:= range stf.bufferOfSet{
		for _,symbol := range sentence{
			if !stf.terminatorIsLivingInFirst(leftNonTerminator, symbol) {
				stf.First[leftNonTerminator] = append(stf.First[leftNonTerminator], symbol)
				firstSetHasBeenUpdated= true
			}
		}
	}
	stf.flushBufferOfSet()
	return firstSetHasBeenUpdated
}


func (stf *StateTableFormer) GetFirst() {
	stf.initGetFirst()
	for  {
		for stf.initHandlingProductionPosition(); stf.handlingProductionsIsNotOver(); stf.goToHandleNextProduction() {
			for stf.initHandleProductionSentencePosition(); stf.handlingProductionSentenceIsNotOver(); stf.goToHandleNextProductionSentence() {
				stf.handleGettingFirst()
			}
		}
		if !stf.syncBufferOfFirst(){
			break
		}
	}
}

func (stf *StateTableFormer) initGetFirst() {
	stf.First = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
}
func (stf *StateTableFormer) handleGettingFirst() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]
	if handlingSentence.isBlank() {
		stf.handleGettingFirstOfSentenceIsBlank()
	} else {
		stf.handleGettingFirstOfSentenceIsNotBlank()
	}
}

func (stf *StateTableFormer) handleGettingFirstOfSentenceIsBlank() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	stf.bufferOfSet[handlingProduction.leftNonTerminator] = append(stf.bufferOfSet[handlingProduction.leftNonTerminator], blankSymbol)
}
func (stf *StateTableFormer) handleGettingFirstOfSentenceIsNotBlank() {
	handlingSymbols := stf.getHandlingSymbols()
	for index, symbol := range handlingSymbols {
		firstSetOfSymbol := stf.First[symbol]
		switch {
		case stf.isTerminator(symbol):
			stf.appendSymbolToBufferOfSet(symbol)
			return
		case !hasBlankSymbol(firstSetOfSymbol):
			stf.appendSymbolToBufferOfSet(removeBlankSymbol(firstSetOfSymbol)...)
			return
		case hasBlankSymbol(firstSetOfSymbol):
			if index == len(handlingSymbols)-1 {
				stf.appendSymbolToBufferOfSet(firstSetOfSymbol...)
			} else {
				stf.appendSymbolToBufferOfSet(removeBlankSymbol(firstSetOfSymbol)...)
			}
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


func (stf *StateTableFormer)appendSymbolToBufferOfSet(symbols ...string) {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	stf.bufferOfSet[handlingProduction.leftNonTerminator] = append(stf.bufferOfSet[handlingProduction.leftNonTerminator], symbols...)
}


func (stf *StateTableFormer)error() {
	panic("存在没有考虑的情况")
}














