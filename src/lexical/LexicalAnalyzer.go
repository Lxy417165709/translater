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
	specialCharOfDecimal = 'X'
)

const (
	codeOfIdentity = 10086
	codeOfInt      = 1008611
	codeOfDecimal = 12345
)
const (
	kindOfIdentity  = "标识符"
	kindOfInt       = "整数"
	kindOfKeyWord   = "关键字"
	kindOfDelimiter = "界符"
	kindOfOperator  = "操作符"
	KindOfDecimal = "小数"
)

func (la *LexicalAnalyzer) FormSpecialCharToKind() {
	la.specialCharToKind = make(map[byte]string)
	la.specialCharToKind[specialCharOfInt]=kindOfInt
	la.specialCharToKind[specialCharOfIdentity]=kindOfIdentity
	la.specialCharToKind[specialCharOfOperator]=kindOfOperator
	la.specialCharToKind[specialCharOfKeyWord]=kindOfKeyWord
	la.specialCharToKind[specialCharOfDelimiter]=kindOfDelimiter
	la.specialCharToKind[specialCharOfDecimal] = KindOfDecimal
}





type LexicalAnalyzer struct {
	regexpsManager *regexpsManager.RegexpsManager
	variableNFAs   []*stateMachine.NFA
	dfas           []*stateMachine.NFA
	finalNfa *stateMachine.NFA
	canParseWords	       []string
	wordToCode     map[string]int
	wordToKind     map[string]string
	specialCharToKind map[byte] string
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
	la. canParseWords= append(la.canParseWords, la.GetDelimiters()...)
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


func (la *LexicalAnalyzer)ShowParsedTokens(bytes []byte) {
	tokens := la.GetTokens(bytes)
	for index,token := range tokens {
		fmt.Printf("[第 %d 个 token] (%v,\t %s,\t %d)\n",index+1,token.GetValue(),la.GetSpecialCharToKind(token.GetSpecialChar()),token.GetKindCode())
	}
}


func (la *LexicalAnalyzer) ShowKindCode() {
	fmt.Println("--------------- 种别码表 ---------------")
	fmt.Println("单词\t\t|类别\t\t|种别码\t\t")
	for _,word := range la.canParseWords{
		fmt.Printf("%s\t\t|%s\t\t|%d\t\t\n",word,la.wordToKind[word],la.wordToCode[word])
	}
	fmt.Printf(" \t\t|%s\t\t|%d\t\t\n",kindOfIdentity,codeOfIdentity)
	fmt.Printf(" \t\t|%s\t\t|%d\t\t\n",kindOfInt,codeOfInt)
	fmt.Printf(" \t\t|%s\t\t|%d\t\t\n",KindOfDecimal,codeOfDecimal)
	fmt.Println("----------------------------------------")
}


// 只需初始化一次
func (la *LexicalAnalyzer) Init() {

	la.dfas = []*stateMachine.NFA{
		stateMachine.NewNFABuilder(string(specialCharOfDelimiter), la.regexpsManager).BuildNotBlankStateNFA().MarkDown(),
		stateMachine.NewNFABuilder(string(specialCharOfKeyWord), la.regexpsManager).BuildNotBlankStateNFA().MarkDown(),
		stateMachine.NewNFABuilder(string(specialCharOfOperator), la.regexpsManager).BuildNotBlankStateNFA().MarkDown(),
	}

	identityNfa := stateMachine.NewNFABuilder(string(specialCharOfIdentity), la.regexpsManager).BuildNotBlankStateNFA().MarkDown()
	for _,pair := range identityNfa.GetWordEndPair(){
		pair.EndStates.SetCode(codeOfIdentity)
	}
	intNfa := stateMachine.NewNFABuilder(string(specialCharOfInt), la.regexpsManager).BuildNotBlankStateNFA().MarkDown()
	for _,pair := range intNfa.GetWordEndPair(){
		pair.EndStates.SetCode(codeOfInt)
	}
	decimalNfa := stateMachine.NewNFABuilder(string(specialCharOfDecimal), la.regexpsManager).BuildNotBlankStateNFA().MarkDown()
	for _,pair := range decimalNfa.GetWordEndPair(){
		pair.EndStates.SetCode(codeOfDecimal)
	}
	la.variableNFAs = []*stateMachine.NFA{
		identityNfa,
		intNfa,
		decimalNfa,
	}


	la.finalNfa = stateMachine.NewEmptyNFA(la.regexpsManager)
	for i:=0;i<len(la.dfas);i++{
		la.finalNfa.GetStartState().Link(la.dfas[i].GetStartState())
	}
	for i:=0;i<len(la.variableNFAs);i++{
		la.finalNfa.GetStartState().Link(la.variableNFAs[i].GetStartState())
	}
	la.finalNfa.EliminateBlankStates()
	la.FormSpecialCharToKind()
	la.MarkCodeAndFormWordToCodeAndCanParseCode()
	la.ShowKindCode()

	la.finalNfa.OutputNFA(`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfa_Visualization_data\nfa.md`)
}



func (la *LexicalAnalyzer) GetSpecialCharToKind(SpecialChar byte) string{
	return la.specialCharToKind[SpecialChar]
}



func (la *LexicalAnalyzer) MarkCodeAndFormWordToCodeAndCanParseCode() {
	la.wordToCode= make(map[string]int)
	la.wordToKind = make(map[string]string)
	code := 0
	for _,nfa := range la.dfas{
		for _,pair := range nfa.GetWordEndPair(){
			la.wordToCode[pair.Word]=code
			la.wordToKind[pair.Word] = la.GetSpecialCharToKind(nfa.GetRespondingSpecialChar())
			la.canParseWords = append(la.canParseWords,pair.Word)
		}
	}
}

func (la *LexicalAnalyzer)GetAllWordEndPair() []*stateMachine.WordEndPair{
	wordEndPairs := make([]*stateMachine.WordEndPair,0)
	for _,nfa := range la.dfas{
		wordEndPairs = append(wordEndPairs,nfa.GetWordEndPair()...)
	}


	return wordEndPairs
}


