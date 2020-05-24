package syntex

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

func arrayHasTerminator(array []string,terminator string) bool{
	for _,element := range array{
		if element==terminator{
			return true
		}
	}
	return false
}

