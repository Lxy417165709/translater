package syntex

import (
	"conf"
	"strings"
)

type sentence struct{
	symbols []string
}

func NewSentence(symbols []string) *sentence{
	return &sentence{
		symbols:symbols,
	}
}
func (s *sentence)Parse(line string) {
	line = strings.TrimSpace(line)
	symbols := strings.Split(line,conf.GetConf().SyntaxConf.DelimiterOfSymbols)
	for _,symbol := range symbols{
		s.symbols = append(s.symbols,strings.TrimSpace(symbol))
	}
}

func  (s *sentence)IsBlank() bool{
	return len(s.symbols)==1 && s.symbols[0]==conf.GetConf().SyntaxConf.BlankSymbol
}

func (s *sentence)FirstSymbolIsTerminator(terminators []string) bool{
	if len(s.symbols)==0{
		return false
	}
	for _,terminator := range terminators{
		if terminator == s.symbols[0]{
			return true
		}
	}
	return false
}
func (s *sentence)FirstSymbolIsNotTerminator(terminators []string) bool{
	return !s.FirstSymbolIsTerminator(terminators)
}





