package machine

import (
	"fmt"
	"os"
	"sort"
)

func (tp *TokenParser) FormTheMarkdownFileOfTokens(filePath string)error {
	var file *os.File
	var err error
	if file, err = os.Create(filePath); err != nil{
		return err
	}
	defer file.Close()
	lines := tp.changeTokensToFileLines()
	for _,line  := range lines {
		if _,err = file.WriteString(line);err!=nil{
			return err
		}
	}
	return nil
}
func (tp *TokenParser) ShowTheMarkdownFileOfTokens() {
	lines := tp.changeTokensToFileLines()
	for i:=0;i<len(lines);i++{
		fmt.Print(lines[i])
	}

}
func (tp *TokenParser) changeTokensToFileLines() []string{
	lines := make([]string,0)
	lines = append(lines,"索引|值|类型|种别码\n")
	lines = append(lines,"--|--|--|--\n")
	for index, token := range tp.tokens {
		lines = append(lines, token.toLine(index))
	}
	return lines
}


func (tp *TokenParser) ShowTheMarkdownFileOfAllTokens() {
	lines := tp.getKindCodeLines()
	for i:=0;i<len(lines);i++{
		fmt.Print(lines[i])
	}
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
		lines = append(lines,token.toLine(index))
	}
	return lines
}

func (tp *TokenParser) wordPairToToken(wordPair *wordPair) *Token{
	token := &Token{
		wordPair.specialChar,
		tp.specialCharTable.GetCode(wordPair.specialChar,wordPair.word),
		tp.specialCharTable.GetType(wordPair.specialChar),
		wordPair.word,
	}
	return token
}
func (tp *TokenParser) removeDuplicateWordPairs(wordPairs []*wordPair) []*wordPair{
	result := make([]*wordPair,0)
	variableCharHasRecord := make(map[byte]bool)
	for _,wordPair := range wordPairs{
		if tp.specialCharTable.CharIsVariable(wordPair.specialChar) {
			if variableCharHasRecord[wordPair.specialChar]{
				continue
			}
			variableCharHasRecord[wordPair.specialChar]=true
			wordPair.word = ""
		}
		result = append(result,wordPair)
	}
	return result

}
func (tp *TokenParser)sortWordPairs(wordPairs []*wordPair) {
	sort.Slice(wordPairs, func(i, j int) bool {
		iCode := tp.specialCharTable.GetCode(wordPairs[i].specialChar,wordPairs[i].word)
		jCode := tp.specialCharTable.GetCode(wordPairs[j].specialChar,wordPairs[j].word)
		return iCode < jCode
	})
}






