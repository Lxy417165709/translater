package lexical

import (
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
	kindOfIdentity = "标识符"
	kindOfInt = "整数"
	kindOfKeyWord = "关键字"
	kindOfDelimiter = "界符"
	kindOfOperator = "操作符"
)


type LexicalAnalyzer struct {
	regexpsManager   *regexpsManager.RegexpsManager
	dfas             []*stateMachine.NFA
	dfaToSpecialChar map[*stateMachine.NFA]byte
	wordToCode       map[string]int
	wordToKind		 map[string]string
}

func NewLexicalAnalyzer(regexpsManager *regexpsManager.RegexpsManager) *LexicalAnalyzer {
	return &LexicalAnalyzer{
		regexpsManager: regexpsManager,
	}
}
func (la *LexicalAnalyzer) Init() {

	identifyDFA := stateMachine.NewNFABuilder(string(specialCharOfIdentity), la.regexpsManager).BuildDFA()
	operatorDFA := stateMachine.NewNFABuilder(string(specialCharOfOperator), la.regexpsManager).BuildDFA()
	intDFA := stateMachine.NewNFABuilder(string(specialCharOfInt), la.regexpsManager).BuildDFA()
	delimiterDFA := stateMachine.NewNFABuilder(string(specialCharOfDelimiter), la.regexpsManager).BuildDFA()
	keyWordDFA := stateMachine.NewNFABuilder(string(specialCharOfKeyWord), la.regexpsManager).BuildDFA()
	la.dfaToSpecialChar = make(map[*stateMachine.NFA]byte)
	la.dfaToSpecialChar[identifyDFA] = specialCharOfIdentity
	la.dfaToSpecialChar[operatorDFA] = specialCharOfOperator
	la.dfaToSpecialChar[intDFA ] = specialCharOfInt
	la.dfaToSpecialChar[delimiterDFA] = specialCharOfDelimiter
	la.dfaToSpecialChar[keyWordDFA] = specialCharOfKeyWord

	la.dfas = append(la.dfas,
		identifyDFA,
		intDFA,
		operatorDFA,
		delimiterDFA,
		keyWordDFA,
	)
	la.FormWordToCode()
	la.FormWordToKind()
}
func (la *LexicalAnalyzer) FormWordToCode() {
	la.wordToCode = make(map[string]int)
	words := make([]string, 0)
	words = append(words, la.GetKeyWords()...)
	words = append(words, la.GetOperators()...)
	words = append(words, la.GetDelimiters()...)
	for index, word := range words {
		la.wordToCode[word] = index + 1
	}
}
func (la *LexicalAnalyzer) FormWordToKind() {
	la.wordToKind = make(map[string]string)
	for _,word := range la.GetKeyWords(){
		la.wordToKind[word]=kindOfKeyWord
	}
	for _,word := range la.GetOperators(){
		la.wordToKind[word]=kindOfOperator
	}
	for _,word := range la.GetDelimiters(){
		la.wordToKind[word]=kindOfDelimiter
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

func (la *LexicalAnalyzer) Parse(bytes []byte) []*pair {
	pairs := make([]*pair, 0)
	for _, dfa := range la.dfas {
		for _, word := range dfa.Get(string(bytes)) {
			// TODO: 有问题，关键字被识别为标识符
			pairs = append(pairs,
				NewPair(
					la.GetKind(dfa,word),
					la.GetCode(dfa, word),
					word,
				),
			)
		}
	}
	return pairs
}
func (la *LexicalAnalyzer) GetKind(handlingDFA *stateMachine.NFA, word string) string {
	if la.dfaToSpecialChar[handlingDFA] == specialCharOfIdentity {
		return kindOfIdentity
	}
	if la.dfaToSpecialChar[handlingDFA] == specialCharOfInt {
		return kindOfInt
	}
	return la.wordToKind[word]
}
func (la *LexicalAnalyzer) GetCode(handlingDFA *stateMachine.NFA, word string) int {
	if la.dfaToSpecialChar[handlingDFA] == specialCharOfIdentity {
		return codeOfIdentity
	}
	if la.dfaToSpecialChar[handlingDFA] == specialCharOfInt {
		return codeOfInt
	}
	return la.wordToCode[word]
}
