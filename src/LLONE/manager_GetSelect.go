package LLONE

import "fmt"

// TODO : 应该还有些情况没考虑
func (stf *StateTableFormer)GetSelect() {
	stf.Select = make(map[*sentence][]string)
	for _,production := range stf.productions{
		for _,sentence := range production.sentences{
			switch  {
			case sentence.isBlank():
				stf.Select[sentence] = stf.Follow[production.leftNonTerminator]
			case sentence.firstSymbolIsTerminator():
				stf.Select[sentence] = []string{sentence.symbols[0]}
			case sentence.firstSymbolIsNotTerminator():
				// TODO:这可能还要考虑 stf.First[sentence.symbols[0]]存在空的情况
				stf.Select[sentence] = stf.First[sentence.symbols[0]]
			}
		}
	}
	for sentence,terminators := range stf.Select{
		fmt.Println(sentence,terminators)
	}
}
