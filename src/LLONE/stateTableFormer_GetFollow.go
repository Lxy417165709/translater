package LLONE

// 感觉可以用模板方法模式
func (stf *StateTableFormer) GetFollow() {
	stf.Follow = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
	stf.Follow["EXP"] = append(stf.Follow["EXP"], "END") // 添加终止符

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
	//for k,v := range stf.Follow{
	//	fmt.Println(k,v)
	//}
}
func (stf *StateTableFormer) handleGettingFollow() {

	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingProductionSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]

	for i := 0; i < len(handlingProductionSentence.symbols); i++ {
		if isTerminator(handlingProductionSentence.symbols[i]) {
			continue
		}
		nextPosition := i + 1
		switch {

		// 位于最后
		case i == len(handlingProductionSentence.symbols)-1:
			stf.bufferOfSet[handlingProductionSentence.symbols[i]] = append(stf.bufferOfSet[handlingProductionSentence.symbols[i]], stf.Follow[handlingProduction.leftNonTerminator]...)

		case i == len(handlingProductionSentence.symbols)-2:
			// 位于倒数第二，而且倒数第一存在空符
			if hasBlankSymbol(stf.First[handlingProductionSentence.symbols[nextPosition]]) {
				stf.bufferOfSet[handlingProductionSentence.symbols[i]] = append(stf.bufferOfSet[handlingProductionSentence.symbols[i]], stf.Follow[handlingProduction.leftNonTerminator]...)
			}
			fallthrough
		default:
			// 冗余
			if isTerminator(handlingProductionSentence.symbols[nextPosition]) {
				stf.bufferOfSet[handlingProductionSentence.symbols[i]] = append(
					stf.bufferOfSet[handlingProductionSentence.symbols[i]],
					handlingProductionSentence.symbols[nextPosition],
				)
			} else {
				stf.bufferOfSet[handlingProductionSentence.symbols[i]] = append(
					stf.bufferOfSet[handlingProductionSentence.symbols[i]],
					removeBlankSymbol(stf.First[handlingProductionSentence.symbols[nextPosition]])...
				)
			}
		}
	}
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
