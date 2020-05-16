package regexpsManager

import (
	"file"
	"fmt"
	"strings"
)

//const (
//	wordDelimiter = "|"
//	grammarUnitDelimiter = "->"
//	grammarFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\grammar.md`
//)

//var regexpsManager = NewRegexpsManager(grammarFilePath)

type RegexpsManager struct {
	grammarFilePath string
	grammarUnitDelimiter string
	wordDelimiter string
	charToRegexp map[byte]string
}
func NewRegexpsManager(grammarFilePath string,grammarUnitDelimiter string,wordDelimiter string) *RegexpsManager {
	return &RegexpsManager{
		grammarFilePath,
		grammarUnitDelimiter,
		wordDelimiter,
		make(map[byte]string),

	}
}


func (nm *RegexpsManager) Init() {
	lines := file.NewFileReader(nm.grammarFilePath).GetFileLines()
	for index, line := range lines {
		unit := NewGrammarUnit(0,"",nm.grammarUnitDelimiter)
		unit.Parse(line)
		nm.addSpecialChar(unit.SpecialChar, unit.Regexp)
		fmt.Printf("添加了第 %d 个特殊字符：%s   对应的正则表达式是：%s\n",index,string(unit.SpecialChar),unit.Regexp)
	}
}
func (nm *RegexpsManager)addSpecialChar(specialChar byte, regexp string) {
	nm.charToRegexp[specialChar] = regexp
}
func (nm *RegexpsManager)GetRegexp(specialChar byte) string{
	return nm.charToRegexp[specialChar]
}
func (nm *RegexpsManager)CharIsSpecial(char byte) bool{
	return nm.charToRegexp[char]!=""
}

// LT$ 这种我们不做处理
// LL 这种我们不做处理
// units[i] == 特殊字符   的时候才进行解析
func (nm *RegexpsManager) GetResponseHandledWords(regexp string) []string {
	units := strings.Split(regexp, nm.wordDelimiter)
	result := make([]string, 0)
	for i := 0; i < len(units); i++ {
		units[i] = strings.TrimSpace(units[i])	// 处理..
		if len(units[i]) == 0 {
			continue
		}
		char := units[i][0]
		if len(units[i]) == 1 && nm.CharIsSpecial(char) {
			result = append(result, nm.GetResponseHandledWords(nm.GetRegexp(char))...)
		} else {
			result = append(result, units[i])
		}
	}
	return result
}



