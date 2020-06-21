package table

import (
	"conf"
)

// 用到了模板方法模式
func (stf *StateTable) formFollow() {
	stf.formFirst()
	stf.templateFunctionOfForming(
		stf.initGetFollow,
		stf.handleGettingFollow,
		stf.syncBufferOfFollow,
	)
}

func (stf *StateTable) initGetFollow() {
	stf.follow = make(map[string][]string)
	stf.bufferOfSet = make(map[string][]string)
	stf.follow[conf.GetConf().SyntaxConf.StartSymbol] = append(
		stf.follow[conf.GetConf().SyntaxConf.StartSymbol],
		conf.GetConf().SyntaxConf.EndSymbol,
	) // 添加终止符

}
func (stf *StateTable) handleGettingFollow() {
	handlingProduction := stf.productions[stf.indexOfHandlingProduction]
	handlingProductionSentence := handlingProduction.sentences[stf.indexOfHandlingProductionSentence]

	for i := 0; i < len(handlingProductionSentence.symbols); i++ {
		nowSymbol := handlingProductionSentence.symbols[i]
		if stf.isTerminator(nowSymbol) {
			continue
		}
		isHandlingLastSymbol := i == len(handlingProductionSentence.symbols)-1
		isHandlingLastTwoSymbol := i == len(handlingProductionSentence.symbols)-2
		nonTerminator := handlingProduction.nonTerminator
		switch {
		case isHandlingLastSymbol :
			stf.appendToBufferOfSet(nowSymbol, stf.follow[nonTerminator]...)
		case isHandlingLastTwoSymbol:
			nextSymbol := handlingProductionSentence.symbols[i+1]
			if hasBlankSymbol(stf.first[nextSymbol]) {
				stf.appendToBufferOfSet(nowSymbol, stf.follow[nonTerminator]...)
			}
			fallthrough
		default:
			nextSymbol := handlingProductionSentence.symbols[i+1]
			if stf.isTerminator(nextSymbol) {
				stf.appendToBufferOfSet(nowSymbol, nextSymbol)
				continue
			}
			stf.appendToBufferOfSet(nowSymbol, removeBlankSymbol(stf.first[nextSymbol])...)
		}
	}
}
func (stf *StateTable) syncBufferOfFollow() bool {
	followSetHasBeenUpdated := false
	for nonTerminator, sentence := range stf.bufferOfSet {
		for _, symbol := range sentence {
			if !stf.terminatorIsLivingInFollow(nonTerminator, symbol) {
				stf.follow[nonTerminator] = append(stf.follow[nonTerminator], symbol)
				followSetHasBeenUpdated = true
			}
		}
	}
	stf.flushBufferOfSet()
	return followSetHasBeenUpdated
}

func (stf *StateTable) terminatorIsLivingInFollow(nonTerminator string, terminator string) bool {
	return arrayHasTerminator(stf.follow[nonTerminator], terminator)
}

