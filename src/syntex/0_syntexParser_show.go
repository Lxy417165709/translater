package syntex







func (sp *SyntaxParser)ShowSyntaxTree() {
	sp.syntaxTreeRoot.Show()
}

func (sp *SyntaxParser)ParseSyntaxTree() {
	sp.syntaxTreeRoot.Parse()
}
//func (sp *SyntaxParser)GetSyntaxTreeResult() int{
//	return sp.syntaxTreeRoot.ParseValue()
//}
