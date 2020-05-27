package lex

import (
	"lex/token"
)

type LexicalAnalyzer struct {
	tokenParser *token.TokenParser
	symbolPairParser *TerminatorPairParser
}

func NewLexicalAnalyzer()*LexicalAnalyzer {
	return &LexicalAnalyzer{
		tokenParser:token.NewTokenParser(),
		symbolPairParser:NewSymbolPairParser(),
	}
}

func (la *LexicalAnalyzer)GetTerminatorPairs(text []byte) []*TerminatorPair{
	tokens := la.tokenParser.GetTokens(text)
	result := make([]*TerminatorPair,0)
	for _,tk := range tokens{
		terminatorPair := la.symbolPairParser.changeTokenToTerminatorPair(tk)
		result = append(result,terminatorPair)
	}
	return result
}

func (la *LexicalAnalyzer)GetAllTerminators() []string{
	terminators := make([]string,0)
	for _,terminator := range la.symbolPairParser.kindCodeToTerminators{
		terminators = append(terminators,terminator)
	}
	return terminators
}

