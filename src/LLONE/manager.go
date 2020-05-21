package LLONE

import (
	"file"
)

type StateTableFormer struct {
	filePath string

	originProductions []*production // 可能有左递归
	productions       []*production // 消除了左递归

	First  map[string][]string
	Follow map[string][]string
	Select map[*sentence][]string
	StateTable map[string]map[string]*sentence
	SentenceToNonTerminator map[*sentence]string

	positionOfHandlingProduction         int
	positionOfHandlingProductionSentence int
	bufferOfSet                          map[string][]string // First,Follow 集合求取过程的缓存
}

func NewStateTableFormer(filePath string) *StateTableFormer {
	stf := &StateTableFormer{filePath: filePath}
	stf.initOriginProductions()
	stf.initProductions()
	stf.initSentenceToNonTerminator()
	return stf
}

func (stf *StateTableFormer) initOriginProductions() {
	lines := file.NewFileReader(stf.filePath).GetFileLines()
	for _, line := range lines {
		production := &production{}
		production.Parse(line)
		stf.originProductions = append(stf.originProductions, production)
	}
}
func (stf *StateTableFormer) initProductions() {
	for _, originProduction := range stf.originProductions {
		stf.productions = append(stf.productions, originProduction.ChangeToNonLeftRecursionProductions()...)
	}
}
func (stf *StateTableFormer) initSentenceToNonTerminator() {
	stf.SentenceToNonTerminator = make(map[*sentence]string)
	for _, production := range stf.productions {
		for _,sentence := range production.sentences{
			stf.SentenceToNonTerminator[sentence] = production.leftNonTerminator
		}
	}
}
