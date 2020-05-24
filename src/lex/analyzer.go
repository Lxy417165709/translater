package lex

import (
	"conf"
	"machine"
)

type LexicalAnalyzer struct {
	tokenParser *machine.TokenParser
	symbolPairParser *SymbolPairParser
}

func NewLexicalAnalyzer(cf *conf.Conf)*LexicalAnalyzer {
	return &LexicalAnalyzer{
		// (cf *conf.Conf)
		tokenParser:machine.NewTokenParser(cf),
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
func (la *LexicalAnalyzer)getTextTokens(text []byte) []*machine.Token {
	la.tokenParser.SetText(text)
	la.tokenParser.ParseTextToTokens()
	return la.tokenParser.GetTokens()
}
