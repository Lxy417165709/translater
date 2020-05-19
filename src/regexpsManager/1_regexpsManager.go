package regexpsManager

import (
	"conf"
	"file"
	"strings"
)



const (
	RepeatPlusSymbol = '@'
	RepeatZeroSymbol = '$'
	RegexSplitString = "|"
	Eps = byte(0)
)

var singleRegexpsManager  = &RegexpsManager{}


type RegexpsManager struct {
	grammarConf *conf.GrammarConf
	charToRegexp map[byte]string
	fixedWordToToken	map[string]*Token
	variableCharToken	map[byte]*Token
}


func GetRegexpsManager() *RegexpsManager{
	return singleRegexpsManager
}
func Init(grammarConf *conf.GrammarConf) {
	singleRegexpsManager.grammarConf = grammarConf
	singleRegexpsManager.InitCharToRegexp()
	singleRegexpsManager.InitTokenOfFixedWord()
	singleRegexpsManager.InitTokenOfVariableCharToken()
}

func (nm *RegexpsManager) InitCharToRegexp() {
	nm.charToRegexp = make(map[byte]string)
	lines := file.NewFileReader(nm.grammarConf.FilePath).GetFileLines()
	for _, line := range lines {
		unit := NewGrammarUnit(0,"",nm.grammarConf.UnitDelimiter)
		unit.Parse(line)
		nm.addSpecialChar(unit.SpecialChar, unit.Regexp)
	}
}

func (nm *RegexpsManager) InitTokenOfVariableCharToken() {
	nm.variableCharToken = make(map[byte]*Token)
	for _,char := range []byte(nm.grammarConf.SpecialCharOfVariableWord) {
		nm.variableCharToken[char] = &Token{
			char,
			int(char)*100,
			"",
		}
	}
}
func (nm *RegexpsManager) InitTokenOfFixedWord() {
	nm.fixedWordToToken = make(map[string]*Token)
	fixedWords := make([]string,0)
	fixedWordToSpecialChar := make(map[string]byte)
	for _,specialChar := range []byte(nm.grammarConf.SpecialCharOfFixedWord){
		for _,fixedWord := range nm.getResponseHandledWords(specialChar){
			fixedWords = append(fixedWords,fixedWord)
			fixedWordToSpecialChar[fixedWord] = specialChar
		}
	}




	for code,fixedWord := range fixedWords{
		nm.fixedWordToToken[fixedWord] = &Token{
			fixedWordToSpecialChar[fixedWord],
			code,
			fixedWord,
		}
	}
}




func (nm *RegexpsManager)GetSpecialCharFormRegexp(regexp string) byte{
	if nm.regexpIsSpecial(regexp) {
		return regexp[0]
	}
	return Eps
}

func (nm *RegexpsManager)regexpIsSpecial(regexp string) bool{
	return len(regexp) == 1 && nm.CharIsSpecial(regexp[0])
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
func (nm *RegexpsManager) getResponseHandledWords(specialChar byte) []string {
	units := strings.Split(nm.GetRegexp(specialChar),nm.grammarConf.WordDelimiter)
	result := make([]string, 0)
	for i := 0; i < len(units); i++ {
		result = append(result, strings.TrimSpace(units[i]))
	}
	return result
}


func (nm *RegexpsManager) GetToken(specialChar byte,word string) *Token{
	if nm.isVariableChar(specialChar) {
		return nm.variableCharToken[specialChar]
	}
	if nm.isFixedChar(specialChar) {
		return nm.fixedWordToToken[word]
	}
	panic("token获取存在错误")
}
func (nm *RegexpsManager) GetAllTokens() []*Token{
	tokens := make([]*Token,0)
	for _,token := range nm.fixedWordToToken{
		tokens = append(tokens, token)
	}
	for _,token := range nm.variableCharToken{
		tokens = append(tokens, token)
	}
	return tokens
}




func (nm *RegexpsManager) isVariableChar(specialChar byte) bool{
	return charIsLiving(specialChar,[]byte(nm.grammarConf.SpecialCharOfVariableWord))
}
func (nm *RegexpsManager) isFixedChar(specialChar byte) bool{
	return charIsLiving(specialChar,[]byte(nm.grammarConf.SpecialCharOfFixedWord))
}
func charIsLiving(targetChar byte,charSet []byte) bool{
	for _,char := range charSet{
		if targetChar==char{
			return true
		}
	}
	return false
}
