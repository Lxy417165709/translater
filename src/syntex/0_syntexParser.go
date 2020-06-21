package syntex

import (
	"lex"
	"lex/terminator"
	"syntex/table"
)


type SyntaxParser struct {
	lexicalAnalyzer *lex.LexicalAnalyzer
	stateTable *table.StateTable

	syntaxTreeRoot *TreeNode
	treeNodeStack []*TreeNode
	terminatorPairs []*terminator.Pair
	readingPosition int			// TODO: 命名不好
}

func NewSyntaxParser() *SyntaxParser{
	lexicalAnalyzer := lex.NewLexicalAnalyzer()
	sp := &SyntaxParser{
		lexicalAnalyzer:lexicalAnalyzer,
		stateTable:table.NewStateTable(lexicalAnalyzer.GetAllTerminators()),
	}
	return sp
}
