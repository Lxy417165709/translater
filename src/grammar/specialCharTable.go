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
}

func NewSpecialCharTable(conf *conf.GrammarConf) *SpecialCharTable {
	sct := &SpecialCharTable{
		delimiterOfPieces: conf.DelimiterOfPieces,
		delimiterOfWords:  conf.DelimiterOfWords,
	}
	sct.initSpecialCharItems(conf.SpecialCharTableFilePath)
	sct.initSpecialCharToRegexp()
	return sct
}
func (sct *SpecialCharTable) Parse(itemContents []string) {
	for _, itemContent := range itemContents {
		sct.specialCharItems = append(
			sct.specialCharItems,
			NewSpecialCharItem(strings.TrimSpace(itemContent), sct.delimiterOfPieces, sct.delimiterOfWords),
		)
	}
}
func (sct *SpecialCharTable) Show() {
	for _, item := range sct.specialCharItems {
		item.Show()
	}
}
func (sct *SpecialCharTable)CharIsSpecial(char byte) bool{
	return sct.specialCharToRegexp[char]!=nil
}
func (sct *SpecialCharTable)GetRegexp(specialChar byte) *Regexp{
	return sct.specialCharToRegexp[specialChar]
}

func (sct *SpecialCharTable) initSpecialCharItems(filePath string) {
	lines := file.NewFileReader(filePath).GetFileLines()
	itemContents := lines[2:]
	sct.Parse(itemContents)
}
func (sct *SpecialCharTable)initSpecialCharToRegexp() {
	sct.specialCharToRegexp = make(map[byte]*Regexp)
	for _,item := range sct.specialCharItems{
		sct.specialCharToRegexp[item.specialChar] = item.regexp
	}
}



