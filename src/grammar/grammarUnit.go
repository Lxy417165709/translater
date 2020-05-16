package grammar

import (
	"fmt"
	"strings"
)
const (
	grammarUnitDelimiter = "->"
	grammarFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\grammar.md`
)
type GrammarUnit struct{
	SpecialChar byte
	Regexp   string
}
func NewGrammarUnit(specialChar byte,regexp string) *GrammarUnit{
	return &GrammarUnit{specialChar,regexp}
}


func (unit *GrammarUnit) Parse(line string) {
	parts := strings.Split(strings.TrimSpace(line), grammarUnitDelimiter)
	if len(parts) != 2 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于2", parts))
	}
	identity := strings.TrimSpace(parts[0])[0]
	regexp := strings.TrimSpace(parts[1])
	unit.SpecialChar = identity
	unit.Regexp = regexp
}
