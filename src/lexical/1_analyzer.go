package lexical

import (
	"conf"
	"stateMachine"
)



var globalLexicalAnalyzer = &LexicalAnalyzer{}

type LexicalAnalyzer struct {
	lexicalConf *conf.LexicalConf
	nfas        []*stateMachine.NFA
	finalNfa    *stateMachine.NFA
	lexicalDocumentGenerator *lexicalDocumentGenerator
}

func GetLexicalAnalyzer() *LexicalAnalyzer{
	return globalLexicalAnalyzer
}





