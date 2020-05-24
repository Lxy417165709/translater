package grammar

import (
	"conf"
	"file"
	"strings"
)

type SpecialCharTable struct {
	specialCharItems  []*specialCharItem
	delimiterOfPieces string
	delimiterOfWords  string

	specialCharToRegexp map[byte]*Regexp
	specialCharToType map[byte]string
	fixedWordToCode map[string]int
	variableCharToCode map[byte]int
}

func NewSpecialCharTable(cf *conf.Conf) *SpecialCharTable {
	sct := &SpecialCharTable{
		delimiterOfPieces: cf.GrammarConf.DelimiterOfPieces,
		delimiterOfWords:  cf.GrammarConf.DelimiterOfWords,
	}
	sct.initSpecialCharItems(cf.GrammarConf.SpecialCharTableFilePath)
	sct.initSpecialCharToRegexp()
	sct.initSpecialCharToType()
	sct.initFixedWordCodeAndVariableCharToCode()
	return sct
}


func (sct *SpecialCharTable) initSpecialCharItems(filePath string) {
	lines := file.NewFileReader(filePath).GetFileLines()
	itemContents := lines[2:]
	sct.parse(itemContents)
}

func (sct *SpecialCharTable)initSpecialCharToRegexp() {
	sct.specialCharToRegexp = make(map[byte]*Regexp)
	for _,item := range sct.specialCharItems{
		sct.specialCharToRegexp[item.specialChar] = item.regexp
	}
}
func (sct *SpecialCharTable) initSpecialCharToType() {
	sct.specialCharToType = make(map[byte]string)
	for _, item := range sct.specialCharItems {
		sct.specialCharToType[item.specialChar] = item._type
	}
}
func (sct *SpecialCharTable) initFixedWordCodeAndVariableCharToCode() {
	sct.fixedWordToCode = make(map[string]int)
	sct.variableCharToCode = make(map[byte]int)
	codeBase := beginCode
	for _, item := range sct.specialCharItems {
		switch item.kindCodeFlag {
		case coding:
			for index, fixedWord := range item.regexp.GetWords() {
				sct.fixedWordToCode[fixedWord]=codeBase+index
			}
			codeBase += codeStage
		case notCoding:
		default:
			sct.variableCharToCode[item.specialChar]=item.kindCodeFlag
		}
	}
}

func (sct *SpecialCharTable) parse(itemContents []string) {
	for _, itemContent := range itemContents {
		sct.specialCharItems = append(
			sct.specialCharItems,
			NewSpecialCharItem(strings.TrimSpace(itemContent), sct.delimiterOfPieces, sct.delimiterOfWords),
		)
	}
}



