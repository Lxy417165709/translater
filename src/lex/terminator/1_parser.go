package terminator

import (
	"conf"
	"file"
	"fmt"
	"lex/token"
	"strconv"
	"strings"
)

type Parser struct {
	kindCodeToTerminator map[int] string
}

func  NewParser() *Parser{
	spp := &Parser{}
	spp.initKindCodeToTerminator()
	return spp
}

func (spp *Parser) initKindCodeToTerminator() {
	spp.kindCodeToTerminator = make(map[int]string)
	lines := file.NewFileReader(conf.GetConf().LexicalConf.KindCodeToTerminatorFilePath).GetFileLines()
	spp.formKindCodeToTerminator(lines)
}
func (spp *Parser) formKindCodeToTerminator(lines []string) {
	for _,line := range lines{
		line = strings.TrimSpace(line)
		parts := strings.Split(line,conf.GetConf().LexicalConf.DelimiterOfKindCodeToTerminator)
		if len(parts)!=2{
			panic("分割错误，分割后的字段数量不等于2")
		}
		kindCode,err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err!=nil{
			panic(err)
		}
		spp.kindCodeToTerminator[kindCode]= strings.TrimSpace(parts[1])
	}
}

func (spp *Parser) GetAllTerminators()[]string{
	terminators := make([]string,0)
	for _,terminator := range spp.kindCodeToTerminator{
		terminators = append(terminators,terminator)
	}
	return terminators
}

func (spp *Parser) ChangeTokenToTerminatorPair(token *token.Token) *Pair{
	symbol := spp.kindCodeToTerminator[token.GetKindCode()]
	if symbol==""{
		panic(fmt.Sprintf("种别码: %d 没有对应symbol",token.GetKindCode()))
	}

	return &Pair{
		spp.kindCodeToTerminator[token.GetKindCode()],
		token.GetRealValue(),
	}
}



