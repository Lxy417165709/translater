package LLONE



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
	stf.First = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
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
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]
	for pIndex, symbol := range handlingSentence.symbols {
		firstSetOfSymbol := stf.First[symbol]
		switch {
		case isTerminator(symbol):
			stf.bufferOfSet[handlingProduction.leftNonTerminator] = append(stf.bufferOfSet[handlingProduction.leftNonTerminator], symbol)
			return
		case !hasBlankSymbol(firstSetOfSymbol):
			stf.bufferOfSet[handlingProduction.leftNonTerminator] = append(stf.bufferOfSet[handlingProduction.leftNonTerminator], removeBlankSymbol(firstSetOfSymbol)...)
			return
		case hasBlankSymbol(firstSetOfSymbol):
			if pIndex == len(handlingSentence.symbols)-1 {
				stf.bufferOfSet[handlingProduction.leftNonTerminator] = append(stf.bufferOfSet[handlingProduction.leftNonTerminator], firstSetOfSymbol...)
			} else {
				stf.bufferOfSet[handlingProduction.leftNonTerminator] = append(stf.bufferOfSet[handlingProduction.leftNonTerminator],  removeBlankSymbol(firstSetOfSymbol)...)
			}
		default:
			panic("存在没有考虑的情况")
		}
	}
}














