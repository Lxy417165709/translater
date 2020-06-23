package tb

import "env"

type StateTable struct {
	llOneTable *LLOneTable
	dataMatrix map[string]map[string]*Sentence
}

func NewStateTable(conf *env.StateTableConf) (*StateTable, error) {
	table := &StateTable{}
	llOneTable, err := NewLLOneTable(conf.LLOneTableConf)
	if err != nil {
		return nil, err
	}
	table.llOneTable = llOneTable
	table.formDataMatrix()
	return table, nil
}

func (stf *StateTable) formDataMatrix() {
	stf.dataMatrix = make(map[string]map[string]*Sentence)
	sentenceToNonTerminator := stf.llOneTable.GetSentenceToNonTerminator()
	for _sentence, terminators := range stf.llOneTable.GetSelect() {
		for _, terminator := range terminators {
			if stf.dataMatrix[sentenceToNonTerminator[_sentence]] == nil {
				stf.dataMatrix[sentenceToNonTerminator[_sentence]] = make(map[string]*Sentence)
			}
			stf.dataMatrix[sentenceToNonTerminator[_sentence]][terminator] = _sentence
		}
	}
}

func (stf *StateTable) GetSentence(nonTerminator, terminator string) *Sentence {
	return stf.dataMatrix[nonTerminator][terminator]
}
func (stf *StateTable) HasSentence(nonTerminator, terminator string) bool {
	return stf.dataMatrix[nonTerminator]!=nil && stf.dataMatrix[nonTerminator][terminator] != nil
}

