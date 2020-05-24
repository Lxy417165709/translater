package syntex

import (
	"conf"
	"file"
)

const additionCharBeginChar = byte('a')

type StateTableFormer struct {
	filePath string

	originProductions []*production // 可能有左递归
	productions       []*production // 消除了左递归
	terminators []string


	First  map[string][]string
	Follow map[string][]string
	Select map[*sentence][]string
	StateTable map[string]map[string]*sentence
	SentenceToNonTerminator map[*sentence]string

	positionOfHandlingProduction         int
	positionOfHandlingProductionSentence int
	bufferOfSet                          map[string][]string // First,Follow 集合求取过程的缓存
}

// 这里的参数可以进行配置
func NewStateTableFormer(cf *conf.Conf) *StateTableFormer {
	stf := &StateTableFormer{filePath: cf.SyntaxConf.SyntaxFilePath}
	stf.initTerminators()
	stf.initOriginProductions()
	stf.initProductions()
	stf.initSentenceToNonTerminator()
	stf.GetFirst()
	stf.GetFollow()
	stf.GetSelect()
	stf.GetStateTable()
	return stf
}

func (stf *StateTableFormer)initTerminators() {
	// TODO: 这个可以进行获取，不用手动配置
	stf.terminators = []string{
		"LEFT_PAR", "RIGHT_PAR", "IDE", "ADD", "SUB","ZS","EQR","DIV","FAC",endSymbol,
	}
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

func (stf *StateTableFormer) GetSentence(nonTerminator,terminator string) *sentence{
	return stf.StateTable[nonTerminator][terminator]
}
func (stf *StateTableFormer) HasSentence(nonTerminator,terminator string) bool{
	return stf.StateTable[nonTerminator][terminator]!=nil
}

func (stf *StateTableFormer)isTerminator(symbol string) bool {
	for _,terminator := range stf.terminators{
		if symbol == terminator{
			return true
		}
	}
	return false
}
