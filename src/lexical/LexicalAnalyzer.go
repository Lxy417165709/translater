package lexical

import (
	"fmt"
	"regexpsManager"
	"stateMachine"
)

const (
	specialCharOfIdentity  = 'I'
	specialCharOfInt       = 'Z'
	specialCharOfKeyWord   = 'W'
	specialCharOfOperator  = 'O'
	specialCharOfDelimiter = 'J'
)

const (
	codeOfIdentity = 10086
	codeOfInt      = 1008611
)
const (
	kindOfIdentity  = "标识符"
	kindOfInt       = "整数"
	kindOfKeyWord   = "关键字"
	kindOfDelimiter = "界符"
	kindOfOperator  = "操作符"
)

type LexicalAnalyzer struct {
	regexpsManager *regexpsManager.RegexpsManager
	dfas           []*stateMachine.NFA
	canParseWords	       []string
	wordToCode     map[string]int
	wordToKind     map[string]string
}

func NewLexicalAnalyzer(regexpsManager *regexpsManager.RegexpsManager) *LexicalAnalyzer {
	return &LexicalAnalyzer{
		regexpsManager: regexpsManager,
	}
}

// 只需初始化一次
func (la *LexicalAnalyzer) Init() {


	identifyDFA := stateMachine.NewNFABuilder(string(specialCharOfIdentity), la.regexpsManager).BuildDFA()
	operatorDFA := stateMachine.NewNFABuilder(string(specialCharOfOperator), la.regexpsManager).BuildDFA()
	intDFA := stateMachine.NewNFABuilder(string(specialCharOfInt), la.regexpsManager).BuildDFA()
	delimiterDFA := stateMachine.NewNFABuilder(string(specialCharOfDelimiter), la.regexpsManager).BuildDFA()
	keyWordDFA := stateMachine.NewNFABuilder(string(specialCharOfKeyWord), la.regexpsManager).BuildDFA()

	// 这里会影响parse，先加入的优先级越大
	la.dfas = append(la.dfas,
		keyWordDFA,
		identifyDFA,
		intDFA,
		operatorDFA,
		delimiterDFA,
	)
	la.FormCanParsedWordAndWordToCode()
	la.FormWordToKind()
}
func (la *LexicalAnalyzer) FormCanParsedWordAndWordToCode() {
	la.wordToCode = make(map[string]int)
	la.canParseWords = make([]string, 0)
	la.canParseWords = append(la.canParseWords, la.GetKeyWords()...)
	la.canParseWords = append(la.canParseWords, la.GetOperators()...)
	la. canParseWords= append(la.canParseWords, la.GetDelimiters()...)
	for index, word := range la.canParseWords {
		la.wordToCode[word] = index + 1
	}
}
func (la *LexicalAnalyzer) FormWordToKind() {
	la.wordToKind = make(map[string]string)
	for _, word := range la.GetKeyWords() {
		la.wordToKind[word] = kindOfKeyWord
	}
	for _, word := range la.GetOperators() {
		la.wordToKind[word] = kindOfOperator
	}
	for _, word := range la.GetDelimiters() {
		la.wordToKind[word] = kindOfDelimiter
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

func (la *LexicalAnalyzer) GetTokens(bytes []byte) []*token {
	tokens := make([]*token, 0)
	wordToSpecialChar := make(map[string]byte)
	for _, dfa := range la.dfas {
		for _, word := range dfa.Get(string(bytes)) {
			if wordToSpecialChar[word] == 0 || wordToSpecialChar[word] == dfa.GetRespondingSpecialChar() {
				wordToSpecialChar[word] = dfa.GetRespondingSpecialChar()
				tokens = append(tokens,
					NewPair(
						la.GetKind(dfa, word),
						la.GetCode(dfa, word),
						word,
					),
				)
			}
		}
	}
	return tokens
}


func (la *LexicalAnalyzer)ShowParsedTokens(bytes []byte) {
	tokens := la.GetTokens(bytes)
	for index,token := range tokens {
		fmt.Printf("[第 %d 个 token] (%v, %s, %d)\n",index+1,token.GetValue(),token.GetKind(),token.GetKindCode())
	}

}
func (la *LexicalAnalyzer) GetKind(handlingDFA *stateMachine.NFA, word string) string {
	if handlingDFA.GetRespondingSpecialChar() == specialCharOfIdentity {
		return kindOfIdentity
	}
	if handlingDFA.GetRespondingSpecialChar() == specialCharOfInt {
		return kindOfInt
	}
	return la.wordToKind[word]
}
func (la *LexicalAnalyzer) GetCode(handlingDFA *stateMachine.NFA, word string) int {
	if handlingDFA.GetRespondingSpecialChar() == specialCharOfIdentity {
		return codeOfIdentity
	}
	if handlingDFA.GetRespondingSpecialChar() == specialCharOfInt {
		return codeOfInt
	}
	return la.wordToCode[word]
}


func (la *LexicalAnalyzer) ShowKindCode() {
	fmt.Println("--------------- 种别码表 ---------------")
	fmt.Println("单词\t\t|类别\t\t|种别码\t\t")
	for _,word := range la.canParseWords{
		fmt.Printf("%s\t\t|%s\t\t|%d\t\t\n",word,la.wordToKind[word],la.wordToCode[word])
	}
	fmt.Printf(" \t\t|%s\t\t|%d\t\t\n",kindOfIdentity,codeOfIdentity)
	fmt.Printf(" \t\t|%s\t\t|%d\t\t\n",kindOfInt,codeOfInt)
	fmt.Println("----------------------------------------")
}
