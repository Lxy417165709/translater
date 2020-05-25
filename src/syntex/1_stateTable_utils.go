package syntex

import (
	"conf"
	"file"
)

func (stf *StateTable) templateFunctionOfForming(initFunction func(),handleFunction func(),syncBufferFunction func() bool) {
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

func (stf *StateTable) flushBufferOfSet() {
	stf.bufferOfSet = map[string][]string{}
}
func (stf *StateTable) initHandlingProductionPosition() {
	stf.positionOfHandlingProduction = 0
}
func (stf *StateTable) initHandleProductionSentencePosition() {
	stf.positionOfHandlingProductionSentence = 0
}
func (stf *StateTable) goToHandleNextProduction() {
	stf.positionOfHandlingProduction++
}
func (stf *StateTable) goToHandleNextProductionSentence() {
	stf.positionOfHandlingProductionSentence++
}
func (stf *StateTable) handlingProductionsIsNotOver() bool {
	return stf.positionOfHandlingProduction < len(stf.productions)
}
func (stf *StateTable) handlingProductionSentenceIsNotOver() bool {
	handlingProduction := stf.productions[stf.positionOfHandlingProduction]
	return stf.positionOfHandlingProductionSentence < len(handlingProduction.sentences)
}


func(stf *StateTable)appendToBufferOfSet(key string,symbols ...string) {
	stf.bufferOfSet[key] = append(stf.bufferOfSet[key], symbols...)
}


func (stf *StateTable)getNonTerminators() []string{
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


func getProductions(filePath string) []*production {
	lines := file.NewFileReader(filePath).GetFileLines()
	originProductions := make([]*production,0)
	for _, line := range lines {
		production := NewProduction("",nil)
		production.Parse(line)
		originProductions = append(originProductions, production)
	}
	return originProductions
}


func removeBlankSymbol(symbols []string) []string {
	result := make([]string, 0)
	for _, symbol := range symbols {
		if symbol == conf.GetConf().SyntaxConf.BlankSymbol {
			continue
		}
		result = append(result, symbol)
	}
	return result
}
func hasBlankSymbol(symbols []string) bool {
	for _, symbol := range symbols {
		if symbol == conf.GetConf().SyntaxConf.BlankSymbol {
			return true
		}
	}
	return false
}




