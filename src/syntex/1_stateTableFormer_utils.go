package syntex

func (stf *StateTableFormer) TemplateFunctionOfForming(initFunction func(),handleFunction func(),syncBufferFunction func() bool) {
	initFunction()
	for  {
		for stf.initHandlingProductionPosition(); stf.handlingProductionsIsNotOver(); stf.goToHandleNextProduction() {
			for stf.initHandleProductionSentencePosition(); stf.handlingProductionSentenceIsNotOver(); stf.goToHandleNextProductionSentence() {
				handleFunction()
			}
		}
		if !syncBufferFunction(){
			break
		}
	}
}

func (stf *StateTableFormer) flushBufferOfSet() {
	stf.bufferOfSet = map[string][]string{}
}
func (stf *StateTableFormer) initHandlingProductionPosition() {
	stf.positionOfHandlingProduction = 0
}
func (stf *StateTableFormer) initHandleProductionSentencePosition() {
	stf.positionOfHandlingProductionSentence = 0
}
func (stf *StateTableFormer) goToHandleNextProduction() {
	stf.positionOfHandlingProduction++
}
func (stf *StateTableFormer) goToHandleNextProductionSentence() {
	stf.positionOfHandlingProductionSentence++
}
func (stf *StateTableFormer) handlingProductionsIsNotOver() bool {
	return stf.positionOfHandlingProduction < len(stf.productions)
}
func (stf *StateTableFormer) handlingProductionSentenceIsNotOver() bool {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	return stf.positionOfHandlingProductionSentence < len(handlingProduction.sentences)
}


func(stf *StateTableFormer)appendToBufferOfSet(key string,symbols ...string) {
	stf.bufferOfSet[key] = append(stf.bufferOfSet[key], symbols...)
}


func (stf *StateTableFormer)getNonTerminators() []string{
	result := make([]string,0)
	for _,production := range stf.productions{
		result = append(result,production.leftNonTerminator)
	}
	return result
}



func arrayHasTerminator(array []string,terminator string) bool{
	for _,element := range array{
		if element==terminator{
			return true
		}
	}
	return false
}






