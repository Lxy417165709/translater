package syntex

import (
	"lex"
)


type SyntaxParser struct {
	lexicalAnalyzer *lex.LexicalAnalyzer
	stateTable *StateTable

	symbolsStack []string
	terminatorPairs []*lex.TerminatorPair
	readingPosition int
}

func NewSyntaxParser() *SyntaxParser{
	lexicalAnalyzer := lex.NewLexicalAnalyzer()
	sp := &SyntaxParser{
		lexicalAnalyzer:lexicalAnalyzer,
		stateTable:NewStateTable(lexicalAnalyzer.GetAllTerminators()),
	}
	//fmt.Println("-----")
	//sp.stateTable.Show()
	//fmt.Println("-----")
	return sp
}
