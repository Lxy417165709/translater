package lex

import (
	"lex/terminator"
	"lex/token"
)

type LexicalAnalyzer struct {
	tokenParser *token.Parser
	terminatorPairParser *terminator.Parser
}

func NewLexicalAnalyzer()*LexicalAnalyzer {
	return &LexicalAnalyzer{
		tokenParser:          token.NewParser(),
		terminatorPairParser: terminator.NewParser(),
	}
}

func (la *LexicalAnalyzer)GetPairs(text []byte) []*terminator.Pair{
	tokens := la.tokenParser.GetTokens(text)
	result := make([]*terminator.Pair,0)
	for _,tk := range tokens{
		terminatorPair := la.terminatorPairParser.ChangeTokenToTerminatorPair(tk)
		result = append(result,terminatorPair)
	}
	return result
}

func (la *LexicalAnalyzer)GetAllTerminators() []string{
	return la.terminatorPairParser.GetAllTerminators()
}

