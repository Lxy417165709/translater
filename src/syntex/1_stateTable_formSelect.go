package syntex

// TODO : 应该还有些情况没考虑
func (stf *StateTable)formSelect() {
	stf.formFollow()
	stf._select = make(map[*sentence][]string)
	for _,production := range stf.productions{
		for _,sentence := range production.sentences{
			switch  {
			case sentence.IsBlank():
				stf._select[sentence] = stf.follow[production.leftNonTerminator]
			case sentence.FirstSymbolIsTerminator(stf.terminators):
				stf._select[sentence] = []string{sentence.symbols[0]}
			case sentence.FirstSymbolIsNotTerminator(stf.terminators):
				// TODO:这可能还要考虑 stf.First[sentence.symbols[0]]存在空的情况
				stf._select[sentence] = stf.first[sentence.symbols[0]]
			default:
				stf.errorOfGettingSelect()
			}
		}
	}
}
func (stf *StateTable)errorOfGettingSelect() {
	panic("获取select集时，存在没有考虑的情况")
}

