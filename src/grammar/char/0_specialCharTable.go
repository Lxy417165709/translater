package char

import (
	"conf"
)

const (
	beginCode = 1
	coding = 0
	notCoding= -1
	codeStage = 100
)


type SpecialCharTable struct {
	specialCharItems  []*specialCharItem
	specialCharToRegexp map[byte]*Regexp
	specialCharToType map[byte]string
	fixedWordToCode map[string]int
	variableCharToCode map[byte]int
}

func NewSpecialCharTable() *SpecialCharTable {
	sct := &SpecialCharTable{}
	sct.initSpecialCharItems(conf.GetConf().GrammarConf.SpecialCharTableFilePath)
	sct.initSpecialCharToRegexp()
	sct.initSpecialCharToType()
	sct.initFixedWordCodeAndVariableCharToCode()
	return sct
}








