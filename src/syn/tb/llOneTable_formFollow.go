package tb



// 用到了模板方法模式
func (one *LLOneTable) formFollow() {
	one.formFirst()
	one.templateFunctionOfForming(
		one.initGetFollow,
		one.handleGettingFollow,
		one.syncBufferOfFollow,
	)
}

func (one *LLOneTable) initGetFollow() {
	one.follow = make(map[string][]string)
	one.bufferOfSet = make(map[string][]string)
	one.follow[one.conf.StartSymbol] = append(
		one.follow[one.conf.StartSymbol],
		one.conf.EndSymbol,
	) // 添加终止符

}
func (one *LLOneTable) handleGettingFollow() {
	handlingProduction := one.productions[one.indexOfHandlingProduction]
	handlingProductionSentence := handlingProduction.sentences[one.indexOfHandlingProductionSentence]

	for i := 0; i < len(handlingProductionSentence.Symbols); i++ {
		nowSymbol := handlingProductionSentence.Symbols[i]
		if one.isTerminator(nowSymbol) {
			continue
		}
		isHandlingLastSymbol := i == len(handlingProductionSentence.Symbols)-1
		isHandlingLastTwoSymbol := i == len(handlingProductionSentence.Symbols)-2
		nonTerminator := handlingProduction.nonTerminator
		switch {
		case isHandlingLastSymbol :
			one.appendToBufferOfSet(nowSymbol, one.follow[nonTerminator]...)
		case isHandlingLastTwoSymbol:
			nextSymbol := handlingProductionSentence.Symbols[i+1]
			if one.hasBlankSymbol(one.first[nextSymbol]) {
				one.appendToBufferOfSet(nowSymbol, one.follow[nonTerminator]...)
			}
			fallthrough
		default:
			nextSymbol := handlingProductionSentence.Symbols[i+1]
			if one.isTerminator(nextSymbol) {
				one.appendToBufferOfSet(nowSymbol, nextSymbol)
				continue
			}
			one.appendToBufferOfSet(nowSymbol, one.removeBlankSymbol(one.first[nextSymbol])...)
		}
	}
}
func (one *LLOneTable) syncBufferOfFollow() bool {
	followSetHasBeenUpdated := false
	for nonTerminator, sentence := range one.bufferOfSet {
		for _, symbol := range sentence {
			if !one.terminatorIsLivingInFollow(nonTerminator, symbol) {
				one.follow[nonTerminator] = append(one.follow[nonTerminator], symbol)
				followSetHasBeenUpdated = true
			}
		}
	}
	one.flushBufferOfSet()
	return followSetHasBeenUpdated
}

func (one *LLOneTable) terminatorIsLivingInFollow(nonTerminator string, terminator string) bool {
	return arrayHasTerminator(one.follow[nonTerminator], terminator)
}

