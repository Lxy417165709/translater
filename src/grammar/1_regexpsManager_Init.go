package grammar

import (
	"conf"
	"file"
)

const Eps = byte(0)
const beginCode = 1

func Init(grammarConf *conf.GrammarConf) {
	singleRegexpsManager.grammarConf = grammarConf
	singleRegexpsManager.InitGrammarUnits()
	singleRegexpsManager.InitCharToRegexp()
	singleRegexpsManager.InitFixedWordCodeAndVariableCharCode()
	singleRegexpsManager.InitSpecialCharToType()
}
func (nm *RegexpsManager) InitGrammarUnits() {
	nm.grammarUnits = nm.getAllGrammarUnit()
}
func (nm *RegexpsManager) InitCharToRegexp() {
	nm.charToRegexp = make(map[byte]string)
	for _, unit := range nm.grammarUnits {
		nm.charToRegexp[unit.SpecialChar] = unit.MatchRegexp
	}
}

func (nm *RegexpsManager) InitFixedWordCodeAndVariableCharCode() {
	nm.fixedWordToCode = make(map[string]int)
	nm.variableCharToCode = make(map[byte]int)

	nowCode := beginCode
	for _, grammarUnit := range nm.grammarUnits {
		switch grammarUnit.KindCodeRule{
		case coding:
			for _, fixedWord := range grammarUnit.GetWords() {
				nm.fixedWordToCode[fixedWord]=nowCode
				nowCode++
			}
		case notCoding:
		default:
			nm.variableCharToCode[grammarUnit.SpecialChar]=grammarUnit.KindCodeRule
		}
	}
}

func (nm *RegexpsManager) InitSpecialCharToType() {
	nm.specialCharToType = make(map[byte]string)
	for _, grammarUnit := range nm.grammarUnits {
		nm.specialCharToType[grammarUnit.SpecialChar] = grammarUnit.Type
	}
}




func (nm *RegexpsManager) getAllGrammarUnit() []*GrammarUnit {
	// 前2行是表格格式
	lines := file.NewFileReader(nm.grammarConf.FilePath).GetFileLines()[2:]
	result := make([]*GrammarUnit, 0)
	for i := 0; i < len(lines); i++ {
		unit := NewGrammarUnit(nm.grammarConf.PartDelimiter, nm.grammarConf.RegexpDelimiter)
		unit.Parse(lines[i])
		result = append(result, unit)
	}
	return result
}
