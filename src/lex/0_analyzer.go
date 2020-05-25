package lex

import (
	"grammar/token"
)

type LexicalAnalyzer struct {
	tokenParser *token.TokenParser
	symbolPairParser *SymbolPairParser
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
	for _,token := range tokens{
		terminatorPair := la.symbolPairParser.changeTokenToTerminatorPair(token)
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

