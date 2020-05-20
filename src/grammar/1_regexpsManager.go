package grammar

import (
	"conf"
	"fmt"
)

type RegexpsManager struct {
	grammarConf  *conf.GrammarConf
	grammarUnits []*GrammarUnit
	charToRegexp map[byte]string

	fixedWordToCode  map[string]int
	variableCharToCode map[byte]int
	specialCharToType map[byte]string
}

var singleRegexpsManager = &RegexpsManager{}

func GetRegexpsManager() *RegexpsManager {
	return singleRegexpsManager
}

func (nm *RegexpsManager) GetRegexpDelimiter() string {
	return nm.grammarConf.RegexpDelimiter
}
func (nm *RegexpsManager) GetReformLinesOfGrammarFile() []string{
	result := make([]string,0)
	result = append(result,"索引|特殊符号|类型|种别码编码规则|匹配")
	result = append(result,"--|--|--|--|--")
	grammarUnits := nm.getAllGrammarUnit()
	for index,grammarUnit := range grammarUnits {
		result = append(result,fmt.Sprintf("%d%s%s",index+1,nm.grammarConf.PartDelimiter,grammarUnit.reformToLine()))

	}
	return result
}
func (nm *RegexpsManager) GetSpecialCharFormRegexp(regexp string) byte {
	if nm.regexpIsSpecial(regexp) {
		return regexp[0]
	}
	return Eps
}
func (nm *RegexpsManager) GetRegexp(specialChar byte) string {
	return nm.charToRegexp[specialChar]
}
func (nm *RegexpsManager) CharIsSpecial(char byte) bool {
	return nm.charToRegexp[char] != ""
}

func (nm *RegexpsManager) regexpIsSpecial(regexp string) bool {
	return len(regexp) == 1 && nm.CharIsSpecial(regexp[0])
}



func (nm *RegexpsManager) GetType(specialChar byte) string{
	return nm.specialCharToType[specialChar]
}
func (nm *RegexpsManager) GetCode(word string,specialChar byte) int{
	if nm.isFixedWord(word) {
		return nm.fixedWordToCode[word]
	}
	if nm.IsVariableChar(specialChar) {
		return nm.variableCharToCode[specialChar]
	}
	panic("获取 code 出现错误")
}
func (nm *RegexpsManager) IsVariableChar(specialChar byte) bool {
	return nm.variableCharToCode[specialChar] != 0
}
func (nm *RegexpsManager) isFixedWord(word string) bool {
	return nm.fixedWordToCode[word] != 0
}
