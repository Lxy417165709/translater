package syntex

import (
	"lex"
)


type SyntaxParser struct {
	lexicalAnalyzer *lex.LexicalAnalyzer

	stateTable *StateTableFormer
	symbolsStack []string

	terminatorPairs []*lex.TerminatorPair
	readingPosition int
}

func NewSyntaxParser() *SyntaxParser{
	lexicalAnalyzer := lex.NewLexicalAnalyzer()
	sp := &SyntaxParser{
		lexicalAnalyzer:lexicalAnalyzer,
		stateTable:NewStateTableFormer(lexicalAnalyzer.GetAllTerminators()),
	}
	return sp
}
