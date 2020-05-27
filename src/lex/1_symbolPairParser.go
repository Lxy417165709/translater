package lex

import (
	"fmt"
	"lex/token"
)

type TerminatorPairParser struct {
	kindCodeToTerminators map[int] string
}

func  NewSymbolPairParser() *TerminatorPairParser{
	spp := &TerminatorPairParser{}
	spp.initKindCodeToSymbol()
	return spp
}

// 这个也可以进行配置
func (spp *TerminatorPairParser)initKindCodeToSymbol() {
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
	spp.kindCodeToTerminators[10003]="XS"
}


func (spp *TerminatorPairParser) changeTokenToTerminatorPair(token *token.Token) *TerminatorPair{
	symbol := spp.kindCodeToTerminators[token.GetKindCode()]
	if symbol==""{
		panic(fmt.Sprintf("种别码: %d 没有对应symbol",token.GetKindCode()))
	}

	return &TerminatorPair{
		spp.kindCodeToTerminators[token.GetKindCode()],
		token.GetRealValue(),
	}
}



