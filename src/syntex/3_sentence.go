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
	s.symbols = strings.Split(line,conf.GetConf().SyntaxConf.DelimiterOfSymbols)
}



