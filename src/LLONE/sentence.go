package LLONE

import "strings"

type sentence struct{
	symbols []string
}

func NewEmptySentence() *sentence{
	return &sentence{}
}
func NewSentence(symbols []string) *sentence{
	return &sentence{symbols}
}
func NewBlankSentence() *sentence{
	return &sentence{[]string{blankSymbol}}
}

func (s *sentence)Parse(line string) {
	line = strings.TrimSpace(line)
	s.symbols = strings.Split(line," ")
}

func(s *sentence)isBlank() bool{
	return len(s.symbols)==1 && s.symbols[0]==blankSymbol
}

func (s *sentence)firstSymbolIsTerminator()bool {
	return isTerminator(s.symbols[0])
}
func (s *sentence)firstSymbolIsNotTerminator()bool {
	return !isTerminator(s.symbols[0])
}
