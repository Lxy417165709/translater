package syntex

import (
	"conf"
	"file"
)



type StateTableFormer struct {
	originProductions []*production // 可能有左递归
	productions       []*production // 消除了左递归
	terminators []string

	first  map[string][]string
	follow map[string][]string
	_select map[*sentence][]string
	stateTable map[string]map[string]*sentence
	sentenceToNonTerminator map[*sentence]string

	positionOfHandlingProduction         int
	positionOfHandlingProductionSentence int
	bufferOfSet                          map[string][]string // First,Follow 集合求取过程的缓存
}

// 这里的参数可以进行配置
func NewStateTableFormer(terminators []string) *StateTableFormer {
	stf := &StateTableFormer{
		terminators:terminators,
	}
	stf.initOriginProductions()
	stf.initProductions()
	stf.initSentenceToNonTerminator()
	stf.FormFirst()
	stf.FormFollow()
	stf.FormSelect()
	stf.FormStateTable()

	return stf
}


func (stf *StateTableFormer) initOriginProductions() {
	lines := file.NewFileReader(conf.GetConf().SyntaxConf.SyntaxFilePath).GetFileLines()
	for _, line := range lines {
		production := NewProduction("",nil)
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
	stf.sentenceToNonTerminator = make(map[*sentence]string)
	for _, production := range stf.productions {
		for _,sentence := range production.sentences{
			stf.sentenceToNonTerminator[sentence] = production.leftNonTerminator
		}
	}
}

func (stf *StateTableFormer) GetSentence(nonTerminator,terminator string) *sentence{
	return stf.stateTable[nonTerminator][terminator]
}
func (stf *StateTableFormer) HasSentence(nonTerminator,terminator string) bool{
	return stf.stateTable[nonTerminator][terminator]!=nil
}

func (stf *StateTableFormer)isTerminator(symbol string) bool {
	for _,terminator := range stf.terminators{
		if symbol == terminator{
			return true
		}
	}
	return false
}
