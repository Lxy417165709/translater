package lexical

import (
	file2 "file"
	"fmt"
	"os"
	"regexpsManager"
	"stateMachine"
)



const (
	specialCharOfIdentity  = 'I'
	specialCharOfInt       = 'Z'
	specialCharOfKeyWord   = 'W'
	specialCharOfOperator  = 'O'
	specialCharOfDelimiter = 'J'
	specialCharOfDecimal   = 'X'
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
	variableNFAs      []*stateMachine.NFA
	fixedNFAs         []*stateMachine.NFA
	finalNfa          *stateMachine.NFA
	wordToCode        map[string]int
	wordToKind        map[string]string
	specialCharToKind map[byte]string
}

func NewLexicalAnalyzer() *LexicalAnalyzer {
	return &LexicalAnalyzer{}
}

func (la *LexicalAnalyzer) GetTokens(bytes []byte) []*regexpsManager.Token {
	tokens := la.finalNfa.GetTokens(string(bytes))
	return tokens
}

func (la *LexicalAnalyzer) FromTheMarkdownFileOfTokens(parsedFilePath string, storeFilePath string) {
	tokens := la.GetTokens(file2.NewFileReader(parsedFilePath).GetFileBytes())
	file, err := os.Create(storeFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("索引|值|类型|种别码\n")
	file.WriteString("--|--|--|--\n")
	for index, token := range tokens {
		file.WriteString(fmt.Sprintf("%d|`%v`|`%s`|`%d`\n", index+1, token.GetValue(), la.GetSpecialCharToKind(token.GetSpecialChar()), token.GetKindCode()))
	}
}

func (la *LexicalAnalyzer) FormKindCodeFile(storeFilePath string) {
	file, err := os.Create(storeFilePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("索引|单词|类别|种别码\n")
	file.WriteString("--|--|--|--\n")
	for index, token := range regexpsManager.GetRegexpsManager().GetAllTokens() {
		file.WriteString(fmt.Sprintf("%d|`%s`|`%s`|`%d`\n", index+1, token.GetValue(), la.specialCharToKind[token.GetSpecialChar()], token.GetKindCode()))
	}
}

func (la *LexicalAnalyzer) InitVariableNFA() {
	la.variableNFAs = []*stateMachine.NFA{
		stateMachine.NewNFABuilder(string(specialCharOfIdentity)).BuildNotBlankStateNFA().MarkToken(),
		stateMachine.NewNFABuilder(string(specialCharOfInt)).BuildNotBlankStateNFA().MarkToken(),
		stateMachine.NewNFABuilder(string(specialCharOfDecimal)).BuildNotBlankStateNFA().MarkToken(),
	}
}
func (la *LexicalAnalyzer) InitFixedNFA() {
	la.fixedNFAs = []*stateMachine.NFA{
		stateMachine.NewNFABuilder(string(specialCharOfKeyWord)).BuildNotBlankStateNFA().MarkToken(),
		stateMachine.NewNFABuilder(string(specialCharOfOperator)).BuildNotBlankStateNFA().MarkToken(),
		stateMachine.NewNFABuilder(string(specialCharOfDelimiter)).BuildNotBlankStateNFA().MarkToken(),
	}
}

func (la *LexicalAnalyzer) InitFinalNFA() {
	la.finalNfa = stateMachine.NewNFA(regexpsManager.Eps)
	for i := 0; i < len(la.fixedNFAs); i++ {
		la.finalNfa.AddParallelNFA(la.fixedNFAs[i])
	}
	for i := 0; i < len(la.variableNFAs); i++ {
		la.finalNfa.AddParallelNFA(la.variableNFAs[i])
	}
	la.finalNfa.EliminateBlankStates()
}

// 只需初始化一次
func (la *LexicalAnalyzer) Init() {
	la.InitSpecialCharToKind()
	la.InitVariableNFA()
	la.InitFixedNFA()
	la.InitFinalNFA()
}



const stateMachineDirName = "stateMachine"
const kindCodeFileName = "kindCode.md"
const tokensFileName = "tokens.md"
const finalNFAFileName = "最终状态机.md"




func (la *LexicalAnalyzer) FormLexicalFile(storeFileDir string) {
	la.FormStateMachineFiles(storeFileDir)
	la.FormKindCodeFile(fmt.Sprintf("%s/%s", storeFileDir, kindCodeFileName))
}
func (la *LexicalAnalyzer) FormStateMachineFiles(storeFileDir string) {
	stateMachineStoreFileDir := fmt.Sprintf("%s/%s", storeFileDir, stateMachineDirName)
	if err := os.MkdirAll(stateMachineStoreFileDir, 0666); err != nil {
		panic(err)
	}
	for _, variableNFA := range la.variableNFAs {
		variableNFA.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s状态机.md", stateMachineStoreFileDir, la.GetSpecialCharToKind(variableNFA.GetSpecialChar())))
	}
	for _, fixedNFA := range la.fixedNFAs {
		fixedNFA.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s状态机.md", stateMachineStoreFileDir, la.GetSpecialCharToKind(fixedNFA.GetSpecialChar())))
	}
	la.finalNfa.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s", stateMachineStoreFileDir, finalNFAFileName))

}

func (la *LexicalAnalyzer) GetSpecialCharToKind(SpecialChar byte) string {
	return la.specialCharToKind[SpecialChar]
}
