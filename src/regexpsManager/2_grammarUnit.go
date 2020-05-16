package regexpsManager

import (
	"fmt"
	"strings"
)

type GrammarUnit struct{
	SpecialChar byte
	Regexp   string
	grammarUnitDelimiter string
}
func NewGrammarUnit(specialChar byte,regexp string,grammarUnitDelimiter string) *GrammarUnit{
	return &GrammarUnit{specialChar,regexp,grammarUnitDelimiter}
}


func (unit *GrammarUnit) Parse(line string) {
	parts := strings.Split(strings.TrimSpace(line), unit.grammarUnitDelimiter)
	if len(parts) != 2 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于2", parts))
	}
	identity := strings.TrimSpace(parts[0])[0]
	regexp := strings.TrimSpace(parts[1])
	unit.SpecialChar = identity
	unit.Regexp = regexp
}


