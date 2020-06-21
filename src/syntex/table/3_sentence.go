package table

import (
	"conf"
	"strings"
)

type Sentence struct{
	symbols []string
}

func NewSentence(symbols []string) *Sentence{
	return &Sentence{
		symbols:symbols,
	}
}
func (s *Sentence)Parse(line string) {
	line = strings.TrimSpace(line)
	symbols := strings.Split(line,conf.GetConf().SyntaxConf.DelimiterOfSymbols)
	for _,symbol := range symbols{
		s.symbols = append(s.symbols,strings.TrimSpace(symbol))
	}
}

// TODO: 这个也依赖了全局
func  (s *Sentence)IsBlank() bool{
	return len(s.symbols)==1 && s.symbols[0]==conf.GetConf().SyntaxConf.BlankSymbol
}

func (s *Sentence)FirstSymbolIsTerminator(terminators []string) bool{
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
func (s *Sentence)FirstSymbolIsNotTerminator(terminators []string) bool{
	return !s.FirstSymbolIsTerminator(terminators)
}



func (s *Sentence)GetSymbols() []string{
	return s.symbols
}

