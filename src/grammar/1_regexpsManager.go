package grammar

import (
	"conf"
	"file"
	"fmt"
)

const Eps = byte(0)
const beginCode = 1

var singleRegexpsManager = &RegexpsManager{}

func GetRegexpsManager() *RegexpsManager {
	return singleRegexpsManager
}

type RegexpsManager struct {
	grammarConf  *conf.GrammarConf
	grammarUnits []*GrammarUnit
	charToRegexp map[byte]string

	allWords          []string
	fixedWordToToken  map[string]*Token
	variableCharToken map[byte]*Token
}

func (nm *RegexpsManager) GetRegexpDelimiter() string {
	return nm.grammarConf.RegexpDelimiter
}

func Init(grammarConf *conf.GrammarConf) {
	singleRegexpsManager.grammarConf = grammarConf
	singleRegexpsManager.InitGrammarUnits()
	singleRegexpsManager.InitCharToRegexp()
	singleRegexpsManager.InitToken()
}

func (nm *RegexpsManager) InitGrammarUnits() {
	nm.grammarUnits = nm.getAllGrammarUnit()
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







func (nm *RegexpsManager) InitCharToRegexp() {
	nm.charToRegexp = make(map[byte]string)
	for _, unit := range nm.grammarUnits {
		nm.charToRegexp[unit.SpecialChar] = unit.MatchRegexp
	}
}
func (nm *RegexpsManager) InitToken() {
	nm.InitFixedWordToken()
	nm.InitVariableCharToken()
}
func (nm *RegexpsManager) InitFixedWordToken() {
	nm.fixedWordToToken = make(map[string]*Token)

	nowCode := beginCode
	for _, grammarUnit := range nm.grammarUnits {
		if grammarUnit.KindCodeRule == coding {
			for _, fixedWord := range grammarUnit.GetWords() {
				nm.fixedWordToToken[fixedWord] = &Token{
					grammarUnit.SpecialChar,
					nowCode,
					grammarUnit.Type,
					fixedWord,
				}
				nowCode++
			}
		}
	}
}

func (nm *RegexpsManager) InitVariableCharToken() {
	nm.variableCharToken = make(map[byte]*Token)
	for _, grammarUnit := range nm.grammarUnits {
		if grammarUnit.KindCodeRule != coding && grammarUnit.KindCodeRule != notCoding {
			nm.variableCharToken[grammarUnit.SpecialChar] = &Token{
				grammarUnit.SpecialChar,
				grammarUnit.KindCodeRule,
				grammarUnit.Type,
				"",
			}
		}
	}
}

func (nm *RegexpsManager) GetToken(specialChar byte, word string) *Token {
	if nm.isVariableChar(specialChar) {
		token := nm.variableCharToken[specialChar]
		token.SetValue(word)
		return token
	}
	if nm.isFixedWord(word) {
		return nm.fixedWordToToken[word]
	}
	panic("token获取存在错误")
}
func (nm *RegexpsManager) GetAllTokens() []*Token {
	tokens := make([]*Token, 0)
	for _, token := range nm.fixedWordToToken {
		tokens = append(tokens, token)
	}
	for _, token := range nm.variableCharToken {
		tokens = append(tokens, token)
	}
	return tokens
}

func (nm *RegexpsManager) isVariableChar(specialChar byte) bool {
	return nm.variableCharToken[specialChar] != nil
}
func (nm *RegexpsManager) isFixedWord(word string) bool {
	return nm.fixedWordToToken[word] != nil
}

func (nm *RegexpsManager) GetSpecialCharFormRegexp(regexp string) byte {
	if nm.regexpIsSpecial(regexp) {
		return regexp[0]
	}
	return Eps
}

func (nm *RegexpsManager) regexpIsSpecial(regexp string) bool {
	return len(regexp) == 1 && nm.CharIsSpecial(regexp[0])
}

func (nm *RegexpsManager) addSpecialChar(specialChar byte, regexp string) {
	nm.charToRegexp[specialChar] = regexp
}

func (nm *RegexpsManager) GetRegexp(specialChar byte) string {
	return nm.charToRegexp[specialChar]
}
func (nm *RegexpsManager) CharIsSpecial(char byte) bool {
	return nm.charToRegexp[char] != ""
}
