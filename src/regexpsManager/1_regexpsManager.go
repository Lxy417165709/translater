package regexpsManager

import (
	"file"
	"strings"
)


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


func (nm *RegexpsManager)Init() {
	lines := file.NewFileReader(nm.grammarFilePath).GetFileLines()
	for _, line := range lines {
		unit := NewGrammarUnit(0,"",nm.grammarUnitDelimiter)
		unit.Parse(line)
		nm.addSpecialChar(unit.SpecialChar, unit.Regexp)
		//fmt.Printf("添加了第 %d 个特殊字符：%s   对应的正则表达式是：%s\n",index,string(unit.SpecialChar),unit.Regexp)
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

//  从 W -> begin|end       获得   [begin end]
func (nm *RegexpsManager) GetResponseHandledWords(specialChar byte) []string {
	units := strings.Split(nm.GetRegexp(specialChar), nm.wordDelimiter)
	result := make([]string, 0)
	for i := 0; i < len(units); i++ {
		result = append(result, strings.TrimSpace(units[i]))
	}
	return result
}
