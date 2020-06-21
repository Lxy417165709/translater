package test

import (
	"conf"
	"fmt"
	"grammar/char"
	"grammar/machine"
	"strconv"
	"strings"
)


type NFATestItem struct {
	regexpContent string
	pattern string
	isMatch bool
	nfaBuilder *machine.NFABuilder
}

func NewNFATestItem(content string) Testable{
	item := &NFATestItem{
		nfaBuilder:machine.NewNFABuilder(),
	}
	item.parse(content)
	return item
}

func (imn *NFATestItem) Test() bool{
	regexp := char.NewRegexp(imn.regexpContent)
	nfa := imn.nfaBuilder.BuildNFAByRegexp(regexp).EliminateBlankStates()
	return nfa.IsMatch(imn.pattern) == imn.isMatch
}


func (imn *NFATestItem) GetErrMsg() string{
	return fmt.Sprintf(
		"样例: [模式串: %s 样例串: %s] 出错，期望结果: %v 测试结果为: %v\n",
		imn.regexpContent,
		imn.pattern,
		imn.isMatch,
		!imn.isMatch,
	)
}



func (imn *NFATestItem) parse(line string) {
	parts := strings.Split(strings.TrimSpace(line), conf.GetConf().IsMatchOfNFATestConf.DelimiterOfPieces)
	if len(parts) != 3 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于3", parts))
	}
	imn.regexpContent = strings.TrimSpace(parts[0])
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

