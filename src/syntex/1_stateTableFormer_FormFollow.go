package syntex

import "conf"

// 用到了模板方法模式
func (stf *StateTableFormer) FormFollow() {
	stf.TemplateFunctionOfForming(
		stf.initGetFollow,
		stf.handleGettingFollow,
		stf.syncBufferOfFollow,
	)
}

func (stf *StateTableFormer) initGetFollow() {
	stf.follow = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
	stf.follow[conf.GetConf().SyntaxConf.StartSymbol] = append(
		stf.follow[conf.GetConf().SyntaxConf.StartSymbol],
		conf.GetConf().SyntaxConf.EndSymbol,
	) // 添加终止符
}
func (stf *StateTableFormer) handleGettingFollow() {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	handlingProductionSentence := handlingProduction.sentences[stf.positionOfHandlingProductionSentence]

	for i := 0; i < len(handlingProductionSentence.symbols); i++ {
		nowSymbol := handlingProductionSentence.symbols[i]
		if stf.isTerminator(nowSymbol) {
			continue
		}
		isHandlingLastSymbol := i == len(handlingProductionSentence.symbols)-1
		isHandlingLastTwoSymbol := i == len(handlingProductionSentence.symbols)-2

		switch {
		case isHandlingLastSymbol :
			stf.appendToBufferOfSet(nowSymbol, stf.follow[handlingProduction.leftNonTerminator]...)
		case isHandlingLastTwoSymbol:
			nextSymbol := handlingProductionSentence.symbols[i+1]
			if stf.hasBlankSymbol(stf.first[nextSymbol]) {
				stf.appendToBufferOfSet(nowSymbol, stf.follow[handlingProduction.leftNonTerminator]...)
			}
			fallthrough
		default:
			nextSymbol := handlingProductionSentence.symbols[i+1]
			if stf.isTerminator(nextSymbol) {
				stf.appendToBufferOfSet(nowSymbol, nextSymbol)
				continue
			}
			symbolsOfNotBlankSymbol := stf.removeBlankSymbol(stf.first[nextSymbol])
			stf.appendToBufferOfSet(nowSymbol, symbolsOfNotBlankSymbol...)
		}
	}
}
func (stf *StateTableFormer) syncBufferOfFollow() bool {
	followSetHasBeenUpdated := false
	for leftNonTerminator, sentence := range stf.bufferOfSet {
		for _, symbol := range sentence {
			if !stf.terminatorIsLivingInFollow(leftNonTerminator, symbol) {
				stf.follow[leftNonTerminator] = append(stf.follow[leftNonTerminator], symbol)
				followSetHasBeenUpdated = true
			}
		}
	}
	stf.flushBufferOfSet()
	return followSetHasBeenUpdated
}

func (stf *StateTableFormer) terminatorIsLivingInFollow(leftNonTerminator string, terminator string) bool {
	return arrayHasTerminator(stf.follow[leftNonTerminator], terminator)
}
func (stf *StateTableFormer) removeBlankSymbol(symbols []string) []string {
	result := make([]string, 0)
	for _, symbol := range symbols {
		if symbol == conf.GetConf().SyntaxConf.BlankSymbol {
			continue
		}
		result = append(result, symbol)
	}
	return result
}
func (stf *StateTableFormer) hasBlankSymbol(symbols []string) bool {
	for _, symbol := range symbols {
		if symbol == conf.GetConf().SyntaxConf.BlankSymbol {
			return true
		}
	}
	return false
}
