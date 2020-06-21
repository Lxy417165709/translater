package table

func (stf *StateTable) formTable() {
	stf.formSelect()
	stf.table = make(map[string]map[string]*Sentence)
	sentenceToNonTerminator := stf.getSentenceToNonTerminator()
	for _sentence, terminators := range stf._select {
		for _, terminator := range terminators {
			if stf.table[sentenceToNonTerminator[_sentence]] == nil {
				stf.table[sentenceToNonTerminator[_sentence]] = make(map[string]*Sentence)
			}
			stf.table[sentenceToNonTerminator[_sentence]][terminator] = _sentence
		}
	}
}


func (stf *StateTable) getSentenceToNonTerminator() map[*Sentence]string {
	sentenceToNonTerminator := make(map[*Sentence]string)
	for _, production := range stf.productions {
		for _, sentence := range production.sentences {
			sentenceToNonTerminator[sentence] = production.nonTerminator
		}
	}
	return sentenceToNonTerminator
}
