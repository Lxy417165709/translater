package tb



// TODO : 应该还有些情况没考虑
func (one *LLOneTable)formSelect() {
	one.formFollow()
	one._select = make(map[*Sentence][]string)
	for _,production := range one.productions{
		for _,sentence := range production.sentences{
			switch  {
			case one.SentenceIsBlank(sentence):
				one._select[sentence] = one.follow[production.nonTerminator]
			case one.isTerminator(sentence.Symbols[0]):
				one._select[sentence] = []string{sentence.Symbols[0]}
			default:
				// 非终结符情况
				// TODO:这可能还要考虑 one.First[sentence.symbols[0]]存在空的情况
				one._select[sentence] = one.first[sentence.Symbols[0]]
			}
		}
	}
}
