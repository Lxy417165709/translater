package tb


func (one *LLOneTable) templateFunctionOfForming(initFunction func(), handleFunction func(), syncBufferFunction func() bool) {
	initFunction()
	for {
		for one.initHandlingProductionPosition(); one.handlingProductionsIsNotOver(); one.goToHandleNextProduction() {
			for one.initHandleProductionSentencePosition(); one.handlingProductionSentenceIsNotOver(); one.goToHandleNextProductionSentence() {
				handleFunction()
			}
		}
		if !syncBufferFunction() {
			break
		}
	}
}

func (one *LLOneTable) flushBufferOfSet() {
	one.bufferOfSet = map[string][]string{}
}
func (one *LLOneTable) initHandlingProductionPosition() {
	one.indexOfHandlingProduction = 0
}
func (one *LLOneTable) initHandleProductionSentencePosition() {
	one.indexOfHandlingProductionSentence = 0
}
func (one *LLOneTable) goToHandleNextProduction() {
	one.indexOfHandlingProduction++
}
func (one *LLOneTable) goToHandleNextProductionSentence() {
	one.indexOfHandlingProductionSentence++
}
func (one *LLOneTable) handlingProductionsIsNotOver() bool {
	return one.indexOfHandlingProduction < len(one.productions)
}
func (one *LLOneTable) handlingProductionSentenceIsNotOver() bool {
	handlingProduction := one.productions[one.indexOfHandlingProduction]
	return one.indexOfHandlingProductionSentence < len(handlingProduction.sentences)
}
func (one *LLOneTable) appendToBufferOfSet(key string, symbols ...string) {
	one.bufferOfSet[key] = append(one.bufferOfSet[key], symbols...)
}



// TODO: 这些函数或许可以提取出来
func arrayHasTerminator(array []string, terminator string) bool {
	for _, element := range array {
		if element == terminator {
			return true
		}
	}
	return false
}


func (one *LLOneTable)SentenceIsBlank(sentence *Sentence) bool{
	return len(sentence.Symbols)==1 && sentence.Symbols[0]==one.conf.BlankSymbol
}



func (one *LLOneTable)removeBlankSymbol(symbols []string) []string {
	result := make([]string, 0)
	for _, symbol := range symbols {
		if symbol == one.conf.BlankSymbol {
			continue
		}
		result = append(result, symbol)
	}
	return result
}
func (one *LLOneTable)hasBlankSymbol(symbols []string) bool {
	for _, symbol := range symbols {
		if symbol == one.conf.BlankSymbol {
			return true
		}
	}
	return false
}

