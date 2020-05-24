package syntex



// 感觉可以用模板方法模式
func (stf *StateTableFormer) GetFollow() {
	stf.initGetFollow()
	for  {
		for stf.initHandlingProductionPosition(); stf.handlingProductionsIsNotOver(); stf.goToHandleNextProduction() {
			for stf.initHandleProductionSentencePosition(); stf.handlingProductionSentenceIsNotOver(); stf.goToHandleNextProductionSentence() {
				stf.handleGettingFollow()
			}
		}
		if !stf.syncBufferOfFollow(){
			break
		}
	}
}
func (stf *StateTableFormer) initGetFollow() {
	stf.Follow = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
	stf.Follow[startSymbol] = append(stf.Follow[startSymbol], endSymbol) // 添加终止符
}

func (stf *StateTableFormer) handleGettingFollow() {

	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingProductionSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]

	for i := 0; i < len(handlingProductionSentence.symbols); i++ {
		if stf.isTerminator(handlingProductionSentence.symbols[i]) {
			continue
		}
		nextPosition := i + 1
		switch {

		// 位于最后
		case i == len(handlingProductionSentence.symbols)-1:
			stf.appendToBufferOfSet(handlingProductionSentence.symbols[i],stf.Follow[handlingProduction.leftNonTerminator]...)
		case i == len(handlingProductionSentence.symbols)-2:
			// 位于倒数第二，而且倒数第一存在空符
			if hasBlankSymbol(stf.First[handlingProductionSentence.symbols[nextPosition]]) {
				stf.appendToBufferOfSet(handlingProductionSentence.symbols[i],stf.Follow[handlingProduction.leftNonTerminator]...)
			}
			fallthrough
		default:
			// 冗余
			if stf.isTerminator(handlingProductionSentence.symbols[nextPosition]) {
				stf.appendToBufferOfSet(handlingProductionSentence.symbols[i],handlingProductionSentence.symbols[nextPosition])
			} else {
				symbolsOfNotBlankSymbol := removeBlankSymbol(stf.First[handlingProductionSentence.symbols[nextPosition]])
				stf.appendToBufferOfSet(handlingProductionSentence.symbols[i],symbolsOfNotBlankSymbol...)
			}
		}
	}
}

func(stf *StateTableFormer)appendToBufferOfSet(key string,symbols ...string) {
	stf.bufferOfSet[key] = append(stf.bufferOfSet[key], symbols...)
}


func (stf *StateTableFormer) syncBufferOfFollow() bool {
	followSetHasBeenUpdated := false
	for leftNonTerminator,sentence:= range stf.bufferOfSet{
		for _, symbol := range sentence {
			if !stf.terminatorIsLivingInFollow(leftNonTerminator, symbol) {
				stf.Follow[leftNonTerminator] = append(stf.Follow[leftNonTerminator], symbol)
				followSetHasBeenUpdated = true
			}
		}
	}
	stf.flushBufferOfSet()
	return followSetHasBeenUpdated
}
func (stf *StateTableFormer) terminatorIsLivingInFollow(leftNonTerminator string, terminator string) bool {
	return arrayHasTerminator(stf.Follow[leftNonTerminator],terminator)
}
