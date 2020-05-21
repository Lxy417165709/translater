package LLONE


func (one *LLOne) EliminateDirectLeftRecursion() {
	// TODO:这里假设文法输入都是正确的
	for _, unit := range one.llUnits {
		one.handledUnits = append(one.handledUnits, unit.GetDelimiterLeftRecursionUnits()...)
	}
}
