package table

type StateTable struct {
	productions []*production // 消除了左递归的产生式
	terminators []string

	first   map[string][]string
	follow  map[string][]string
	_select map[*Sentence][]string
	table   map[string]map[string]*Sentence

	indexOfHandlingProduction         int
	indexOfHandlingProductionSentence int
	bufferOfSet                       map[string][]string // first,follow集 求取过程的缓存
}

// 这里的参数可以进行配置
func NewStateTable(terminators []string) *StateTable {
	stf := &StateTable{
		terminators: terminators,
	}
	stf.formTable()
	return stf
}

func (stf *StateTable) GetSentence(nonTerminator, terminator string) *Sentence {
	return stf.table[nonTerminator][terminator]
}
func (stf *StateTable) HasSentence(nonTerminator, terminator string) bool {
	return stf.table[nonTerminator][terminator] != nil
}

func (stf *StateTable) isTerminator(symbol string) bool {
	for _, terminator := range stf.terminators {
		if symbol == terminator {
			return true
		}
	}
	return false
}
