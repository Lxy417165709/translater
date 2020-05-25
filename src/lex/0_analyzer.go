package lex

import (
	"grammar/machine"
)

type LexicalAnalyzer struct {
	tokenParser *machine.TokenParser
	symbolPairParser *SymbolPairParser
}

func NewLexicalAnalyzer()*LexicalAnalyzer {
	return &LexicalAnalyzer{
		tokenParser:machine.NewTokenParser(),
		symbolPairParser:NewSymbolPairParser(),
	}
}

func (la *LexicalAnalyzer)GetTerminatorPairs(text []byte) []*TerminatorPair{
	tokens := la.getTextTokens(text)
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





func (la *LexicalAnalyzer)getTextTokens(text []byte) []*machine.Token {
	la.tokenParser.SetText(text)
	la.tokenParser.ParseTextToFinalTokens()
	return la.tokenParser.GetTokens()
}
