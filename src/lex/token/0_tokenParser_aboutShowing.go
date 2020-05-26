package token

import (
	"fmt"
	"grammar/machine"
	"os"
	"sort"
)

func (tp *TokenParser) FormTheMarkdownFileOfTokens(text []byte,filePath string)error {
	var file *os.File
	var err error
	if file, err = os.Create(filePath); err != nil{
		return err
	}
	defer file.Close()
	lines := tp.changeTokensToFileLines(text)
	for _,line  := range lines {
		if _,err = file.WriteString(line);err!=nil{
			return err
		}
	}
	return nil
}
func (tp *TokenParser) ShowTheMarkdownFileOfTokens(text []byte) {
	lines := tp.changeTokensToFileLines(text)
	for i:=0;i<len(lines);i++{
		fmt.Print(lines[i])
	}

}
func (tp *TokenParser) ShowTheMarkdownFileOfAllTokens() {
	lines := tp.getKindCodeLines()
	for i:=0;i<len(lines);i++{
		fmt.Print(lines[i])
	}
}

func (tp *TokenParser) changeTokensToFileLines(text []byte) []string{
	lines := make([]string,0)
	lines = append(lines,"索引|值|类型|种别码\n")
	lines = append(lines,"--|--|--|--\n")
	for index, token := range tp.GetTokens(text) {
		lines = append(lines, token.ToLine(index))
	}
	return lines
}
func (tp *TokenParser) getKindCodeLines() []string{
	wordPairs := tp.finalNFA.GetAllWordPairs()
	wordPairs = tp.removeDuplicateWordPairs(wordPairs)
	tp.sortWordPairs(wordPairs)

	lines := make([]string,0)
	lines = append(lines,"索引|单词|类别|种别码\n")
	lines = append(lines,"--|--|--|--\n")
	for index, wordPair := range wordPairs {
		token := tp.wordPairToToken(wordPair)
		lines = append(lines,token.ToLine(index))
	}
	return lines
}
func (tp *TokenParser) removeDuplicateWordPairs(wordPairs []*machine.WordPair) []*machine.WordPair{
	result := make([]*machine.WordPair,0)
	variableCharHasRecord := make(map[byte]bool)
	for _,wordPair := range wordPairs{
		if tp.specialCharTable.CharIsVariable(wordPair.GetSpecialChar()) {
			if variableCharHasRecord[wordPair.GetSpecialChar()]{
				continue
			}
			variableCharHasRecord[wordPair.GetSpecialChar()]=true
			wordPair.SetWord("")
		}
		result = append(result,wordPair)
	}
	return result

}
func (tp *TokenParser) sortWordPairs(wordPairs []*machine.WordPair) {
	sort.Slice(wordPairs, func(i, j int) bool {
		iCode := tp.specialCharTable.GetCode(wordPairs[i].GetSpecialChar(),wordPairs[i].GetWord())
		jCode := tp.specialCharTable.GetCode(wordPairs[j].GetSpecialChar(),wordPairs[j].GetWord())
		return iCode < jCode
	})
}





