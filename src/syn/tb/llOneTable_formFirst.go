package tb



func (one *LLOneTable) formFirst() {
	one.templateFunctionOfForming(
		one.initGetFirst,
		one.handleGettingFirst,
		one.syncBufferOfFirst,
	)
}

func (one *LLOneTable) initGetFirst() {
	one.first = make(map[string][]string)
	one.bufferOfSet = make(map[string][]string)
}


func (one *LLOneTable) handleGettingFirst() {
	handlingProduction := one.productions[one.indexOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[one.indexOfHandlingProductionSentence]
	if one.SentenceIsBlank(handlingSentence) {
		one.handleGettingFirstOfSentenceIsBlank()
	} else {
		one.handleGettingFirstOfSentenceIsNotBlank()
	}
}
func (one *LLOneTable) syncBufferOfFirst() bool {
	firstSetHasBeenUpdated := false
	for nonTerminator,sentence:= range one.bufferOfSet{
		for _,symbol := range sentence{
			if !one.terminatorIsLivingInFirst(nonTerminator, symbol) {
				one.first[nonTerminator] = append(one.first[nonTerminator], symbol)
				firstSetHasBeenUpdated= true
			}
		}
	}
	one.flushBufferOfSet()
	return firstSetHasBeenUpdated
}

func (one *LLOneTable) handleGettingFirstOfSentenceIsBlank() {
	handlingProduction := one.productions[one.indexOfHandlingProduction]
	nonTerminator := handlingProduction.nonTerminator
	one.appendToBufferOfSet(nonTerminator,one.conf.BlankSymbol)
}
func (one *LLOneTable) handleGettingFirstOfSentenceIsNotBlank() {
	handlingProduction := one.productions[one.indexOfHandlingProduction]
	handlingSymbols := one.getHandlingSymbols()
	for index, symbol := range handlingSymbols {
		nonTerminator := handlingProduction.nonTerminator
		switch {
		case one.isTerminator(symbol):
			one.appendToBufferOfSet(nonTerminator,symbol)
			return
		case !one.hasBlankSymbol(one.first[symbol]):
			one.appendToBufferOfSet(nonTerminator,one.removeBlankSymbol(one.first[symbol])...)
			return
		case one.hasBlankSymbol(one.first[symbol]):
			isHandlingLastSymbol := index == len(handlingSymbols)-1
			if isHandlingLastSymbol {
				one.appendToBufferOfSet(nonTerminator,one.first[symbol]...)
				continue
			}
			one.appendToBufferOfSet(nonTerminator,one.removeBlankSymbol(one.first[symbol])...)
		default:
			one.error()
		}
	}
}

func (one *LLOneTable)getHandlingSymbols() []string{
	return one.getHandlingSentence().Symbols
}
func (one *LLOneTable)getHandlingSentence() *Sentence{
	handlingProduction := one.productions[one.indexOfHandlingProduction]
	handlingSentence := handlingProduction.sentences[one.indexOfHandlingProductionSentence]
	return handlingSentence
}

// TODO: 这个命名不太准确
func (one *LLOneTable) terminatorIsLivingInFirst(nonTerminator string, terminator string) bool {
	return arrayHasTerminator(one.first[nonTerminator],terminator)
}



func (one *LLOneTable)error() {
	panic("存在没有考虑的情况")
}





