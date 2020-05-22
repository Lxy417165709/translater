package test

import (
	"fmt"
	"machine"
	"strconv"
	"strings"
)


type isMatchOfNFATestItem struct {
	word string
	pattern string
	isMatch bool
	nfaBuilder *machine.NFABuilder
	delimiterOfPieces string
}

func NewIsMatchOfNFATestItem(nfaBuilder *machine.NFABuilder,delimiterOfPieces string,content string) *isMatchOfNFATestItem{
	item := &isMatchOfNFATestItem{nfaBuilder:nfaBuilder,delimiterOfPieces:delimiterOfPieces}
	item.Parse(content)
	return item
}

func (imn *isMatchOfNFATestItem) Test() bool{
	nfa := imn.nfaBuilder.BuildNFAByWord(imn.word).EliminateBlankStates()
	return nfa.IsMatch(imn.pattern) == imn.isMatch
}
func (imn *isMatchOfNFATestItem) Parse(line string) {
	parts := strings.Split(strings.TrimSpace(line), imn.delimiterOfPieces)
	if len(parts) != 3 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于3", parts))
	}
	imn.word = strings.TrimSpace(parts[0])
	imn.pattern = strings.TrimSpace(parts[1])
	matchFlag, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		panic(err)
	}
	imn.isMatch = intToBool(matchFlag)
}
func intToBool(a int) bool {
	return a != 0
}

