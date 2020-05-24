package lex

import (
	"fmt"
	"machine"
)

type SymbolPairParser struct {
	kindCodeToTerminators map[int] string
}

func  NewSymbolPairParser() *SymbolPairParser{
	spp := &SymbolPairParser{}
	spp.initKindCodeToSymbol()
	return spp
}
// 这个也可以进行配置
func (spp *SymbolPairParser)initKindCodeToSymbol() {
	spp.kindCodeToTerminators = make(map[int]string)
	spp.kindCodeToTerminators[101]="ADD"
	spp.kindCodeToTerminators[102]="SUB"
	spp.kindCodeToTerminators[103]="FAC"
	spp.kindCodeToTerminators[104]="DIV"
	spp.kindCodeToTerminators[105]="EQR"
	spp.kindCodeToTerminators[202]="LEFT_PAR"
	spp.kindCodeToTerminators[203]="RIGHT_PAR"
	spp.kindCodeToTerminators[10001]="IDE"
	spp.kindCodeToTerminators[10002]="ZS"
}



func (spp *SymbolPairParser) changeTokenToTerminatorPair(token *machine.Token) *TerminatorPair{
	symbol := spp.kindCodeToTerminators[token.GetKindCode()]
	if symbol==""{
		panic(fmt.Sprintf("种别码: %d 没有对应symbol",token.GetKindCode()))
	}

	return &TerminatorPair{
		spp.kindCodeToTerminators[token.GetKindCode()],
		token.GetRealValue(),
	}
}


type TerminatorPair struct {
	terminator string
	value interface{}
}

func (sp *TerminatorPair) GetSymbol() string{
	return sp.terminator
}
func NewNotValuePair(terminator string) *TerminatorPair {
	return &TerminatorPair{
		terminator:terminator,
	}
}

