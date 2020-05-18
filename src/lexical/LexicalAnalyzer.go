package lexical

import (
	"fmt"
	"os"
	"regexpsManager"
	"stateMachine"
)
const nfaFilePath =  `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\aboutLexicalAnalyzer\nfa.md`
const codeFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\aboutLexicalAnalyzer\code.md`

const (
	specialCharOfIdentity  = 'I'
	specialCharOfInt       = 'Z'
	specialCharOfKeyWord   = 'W'
	specialCharOfOperator  = 'O'
	specialCharOfDelimiter = 'J'
	specialCharOfDecimal   = 'X'
)

const (
	codeOfIdentity = 10086
	codeOfInt      = 1008611
	codeOfDecimal  = 12345
)
const (
	kindOfIdentity  = "标识符"
	kindOfInt       = "整数"
	kindOfKeyWord   = "关键字"
	kindOfDelimiter = "界符"
	kindOfOperator  = "操作符"
	KindOfDecimal   = "小数"
)

func (la *LexicalAnalyzer) InitSpecialCharToKind() {
	la.specialCharToKind = make(map[byte]string)
	la.specialCharToKind[specialCharOfInt] = kindOfInt
	la.specialCharToKind[specialCharOfIdentity] = kindOfIdentity
	la.specialCharToKind[specialCharOfOperator] = kindOfOperator
	la.specialCharToKind[specialCharOfKeyWord] = kindOfKeyWord
	la.specialCharToKind[specialCharOfDelimiter] = kindOfDelimiter
	la.specialCharToKind[specialCharOfDecimal] = KindOfDecimal
}

type LexicalAnalyzer struct {
	regexpsManager    *regexpsManager.RegexpsManager
	variableNFAs      []*stateMachine.NFA
	fixedNFAs         []*stateMachine.NFA
	finalNfa          *stateMachine.NFA
	canParseWords     []string
	wordToCode        map[string]int
	wordToKind        map[string]string
	specialCharToKind map[byte]string
	wordEndPairs      []*stateMachine.WordEndPair
}

func NewLexicalAnalyzer(regexpsManager *regexpsManager.RegexpsManager) *LexicalAnalyzer {
	return &LexicalAnalyzer{
		regexpsManager: regexpsManager,
	}
}

func (la *LexicalAnalyzer) FormCanParsedWordAndWordToCode() {
	la.wordToCode = make(map[string]int)
	la.canParseWords = make([]string, 0)
	la.canParseWords = append(la.canParseWords, la.GetKeyWords()...)
	la.canParseWords = append(la.canParseWords, la.GetOperators()...)
	la.canParseWords = append(la.canParseWords, la.GetDelimiters()...)
	for index, word := range la.canParseWords {
		la.wordToCode[word] = index + 1
	}
}

func (la *LexicalAnalyzer) GetKeyWords() []string {
	return la.regexpsManager.GetResponseHandledWords(specialCharOfKeyWord)
}
func (la *LexicalAnalyzer) GetOperators() []string {
	return la.regexpsManager.GetResponseHandledWords(specialCharOfOperator)
}
func (la *LexicalAnalyzer) GetDelimiters() [] string {
	return la.regexpsManager.GetResponseHandledWords(specialCharOfDelimiter)
}

func (la *LexicalAnalyzer) GetTokens(bytes []byte) []*stateMachine.Token {
	tokens := la.finalNfa.GetTokenByNFA(string(bytes))
	return tokens
}

func (la *LexicalAnalyzer) FromTheMarkdownFileOfTokens(data []byte,filePath string) {
	tokens := la.GetTokens(data)
	file,err := os.Create(filePath)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	file.WriteString("索引|值|类型|种别码\n")
	file.WriteString("--|--|--|--\n")
	for index, token := range tokens {
		file.WriteString(fmt.Sprintf("%d|`%v`|`%s`|`%d`\n", index+1, token.GetValue(), la.GetSpecialCharToKind(token.GetSpecialChar()), token.GetKindCode()))
	}
}

func (la *LexicalAnalyzer) FormKindCodeFile(filePath string) {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("索引|单词|类别|种别码\n")
	file.WriteString("--|--|--|--\n")
	for index, word := range la.canParseWords {
		file.WriteString(fmt.Sprintf("%d|`%s`|`%s`|`%d`\n", index+1,word, la.wordToKind[word], la.wordToCode[word]))
	}
	file.WriteString(fmt.Sprintf("%d| |`%s`|`%d`\n",len(la.canParseWords)+1, kindOfIdentity, codeOfIdentity))
	file.WriteString(fmt.Sprintf("%d| |`%s`|`%d`\n",len(la.canParseWords)+2, kindOfInt, codeOfInt))
	file.WriteString(fmt.Sprintf("%d| |`%s`|`%d`\n",len(la.canParseWords)+3, KindOfDecimal, codeOfDecimal))
}


func (la *LexicalAnalyzer) InitVariableNFA() {
	identityNfa := stateMachine.NewNFABuilder(string(specialCharOfIdentity), la.regexpsManager).BuildNotBlankStateNFA().MarkDown()
	for _, pair := range identityNfa.GetWordEndPair() {
		pair.EndStates.SetCode(codeOfIdentity)
	}
	intNfa := stateMachine.NewNFABuilder(string(specialCharOfInt), la.regexpsManager).BuildNotBlankStateNFA().MarkDown()
	for _, pair := range intNfa.GetWordEndPair() {
		pair.EndStates.SetCode(codeOfInt)
	}
	decimalNfa := stateMachine.NewNFABuilder(string(specialCharOfDecimal), la.regexpsManager).BuildNotBlankStateNFA().MarkDown()
	for _, pair := range decimalNfa.GetWordEndPair() {
		pair.EndStates.SetCode(codeOfDecimal)
	}
	la.variableNFAs = []*stateMachine.NFA{
		identityNfa,
		intNfa,
		decimalNfa,
	}
}
func (la *LexicalAnalyzer) InitFixedNFA() {
	la.fixedNFAs = []*stateMachine.NFA{
		stateMachine.NewNFABuilder(string(specialCharOfKeyWord), la.regexpsManager).BuildNotBlankStateNFA().MarkDown(),
		stateMachine.NewNFABuilder(string(specialCharOfOperator), la.regexpsManager).BuildNotBlankStateNFA().MarkDown(),
		stateMachine.NewNFABuilder(string(specialCharOfDelimiter), la.regexpsManager).BuildNotBlankStateNFA().MarkDown(),
	}
}
func (la *LexicalAnalyzer) InitFinalNFA() {
	la.finalNfa = stateMachine.NewEmptyNFA(la.regexpsManager)
	for i := 0; i < len(la.fixedNFAs); i++ {
		la.finalNfa.GetStartState().Link(la.fixedNFAs[i].GetStartState())
	}
	for i := 0; i < len(la.variableNFAs); i++ {
		la.finalNfa.GetStartState().Link(la.variableNFAs[i].GetStartState())
	}
	la.finalNfa.EliminateBlankStates()
}

// 只需初始化一次
func (la *LexicalAnalyzer) Init() {
	la.InitVariableNFA()
	la.InitFixedNFA()
	la.InitFinalNFA()

	la.InitSpecialCharToKind()
	la.InitOtherInformationAndMarkState()

	la.FormKindCodeFile(codeFilePath)
	la.finalNfa.FormTheMermaidGraphOfNFA(nfaFilePath)
}

func (la *LexicalAnalyzer) GetSpecialCharToKind(SpecialChar byte) string {
	return la.specialCharToKind[SpecialChar]
}

func (la *LexicalAnalyzer) InitOtherInformationAndMarkState() {
	la.wordToCode = make(map[string]int)
	la.wordToKind = make(map[string]string)
	la.canParseWords = make([]string, 0)

	for _, nfa := range la.fixedNFAs {
		for _, pair := range nfa.GetWordEndPair() {
			code := len(la.canParseWords)+1
			la.wordToCode[pair.Word] = code
			la.wordToKind[pair.Word] = la.GetSpecialCharToKind(nfa.GetRespondingSpecialChar())
			la.canParseWords = append(la.canParseWords, pair.Word)
			pair.EndStates.SetCode(code)
		}
	}
}
