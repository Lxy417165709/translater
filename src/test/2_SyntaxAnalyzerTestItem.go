package test

import (
	"conf"
	"fmt"
	"strings"
	"syntex"
)


type SyntaxAnalyzerTestItem struct {
	express string
	isValid bool
	syntaxParser *syntex.SyntaxParser
}

func NewSyntaxAnalyzerTestItem(content string) *SyntaxAnalyzerTestItem{
	item := &SyntaxAnalyzerTestItem{
		syntaxParser:syntex.NewSyntaxParser(),
	}
	item.parse(content)
	return item
}

// 返回是否测试成功
func (sa *SyntaxAnalyzerTestItem) Test() bool{
	return sa.syntaxParser.IsValid([]byte(sa.express)) == sa.isValid
}
func (sa *SyntaxAnalyzerTestItem) GetErrMsg() string{
	return fmt.Sprintf(
		"样例: [表达式: %s] 出错，期望结果: %v 测试结果为: %v\n",
		sa.express,
		sa.isValid,
		!sa.isValid,
	)
}

// TODO: 这用到了全局配置
func (sa *SyntaxAnalyzerTestItem)parse(line string) {
	parts := strings.Split(strings.TrimSpace(line), conf.GetConf().SyntaxAnalyzerTestConf.DelimiterOfPieces)
	if len(parts) != 2 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于2", parts))
	}
	sa.express = strings.TrimSpace(parts[0])
	sa.isValid = stringToBool(strings.TrimSpace(parts[1]))
}
func stringToBool(a string) bool {
	return a == "1"
}

