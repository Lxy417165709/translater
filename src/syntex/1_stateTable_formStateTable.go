package syntex

import "fmt"

func (stf *StateTable) formTable() {
	stf.formSelect()
	stf.table = make(map[string]map[string]*sentence)
	sentenceToNonTerminator := stf.getSentenceToNonTerminator()
	for _sentence, terminators := range stf._select {
		fmt.Println(sentenceToNonTerminator[_sentence])
		for _, terminator := range terminators {
			if stf.table[sentenceToNonTerminator[_sentence]] == nil {
				stf.table[sentenceToNonTerminator[_sentence]] = make(map[string]*sentence)
			}

			stf.table[sentenceToNonTerminator[_sentence]][terminator] = _sentence
		}
	}
	stf.Show()
}


func (stf *StateTable) getSentenceToNonTerminator() map[*sentence]string {
	sentenceToNonTerminator := make(map[*sentence]string)
	for _, production := range stf.productions {
		for _, sentence := range production.sentences {
			sentenceToNonTerminator[sentence] = production.leftNonTerminator
		}
	}
	return sentenceToNonTerminator
}
