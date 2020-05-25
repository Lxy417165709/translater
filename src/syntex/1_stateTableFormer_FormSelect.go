package syntex

import (
	"conf"
)

// TODO : 应该还有些情况没考虑
func (stf *StateTableFormer)FormSelect() {
	stf._select = make(map[*sentence][]string)
	for _,production := range stf.productions{
		for _,sentence := range production.sentences{
			switch  {
			case stf.sentenceIsBlank(sentence):
				stf._select[sentence] = stf.follow[production.leftNonTerminator]
			case stf.sentenceFirstSymbolIsTerminator(sentence):
				stf._select[sentence] = []string{sentence.symbols[0]}
			case stf.sentenceFirstSymbolIsNotTerminator(sentence):
				// TODO:这可能还要考虑 stf.First[sentence.symbols[0]]存在空的情况
				stf._select[sentence] = stf.first[sentence.symbols[0]]
			default:
				stf.errorOfGettingSelect()
			}
		}
	}
}
func (stf *StateTableFormer)errorOfGettingSelect() {
	panic("获取select集时，存在没有考虑的情况")
}

func (stf *StateTableFormer)sentenceIsBlank(sentence *sentence) bool{
	return len(sentence.symbols)==1 && sentence.symbols[0]==conf.GetConf().SyntaxConf.BlankSymbol
}

func (stf *StateTableFormer)sentenceFirstSymbolIsTerminator(sentence *sentence) bool{
	return len(sentence.symbols)>=1 && stf.isTerminator(sentence.symbols[0])
}
func (stf *StateTableFormer)sentenceFirstSymbolIsNotTerminator(sentence *sentence) bool{
	return len(sentence.symbols)>=1 && !stf.isTerminator(sentence.symbols[0])
}
