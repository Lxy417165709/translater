package lexical

import (
	"conf"
	file2 "file"
	"fmt"
	"os"
	"regexpsManager"
	"stateMachine"
)


type LexicalAnalyzer struct {
	lexicalConf *conf.LexicalConf
	nfas []*stateMachine.NFA
	finalNfa          *stateMachine.NFA
}

func NewLexicalAnalyzer(lexicalConf *conf.LexicalConf) *LexicalAnalyzer {
	return &LexicalAnalyzer{lexicalConf:lexicalConf}
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
		file.WriteString(fmt.Sprintf("%d|`%v`|`%s`|`%d`\n", index+1, token.GetValue(), regexpsManager.GetRegexpsManager().GetType(token.GetSpecialChar()), token.GetKindCode()))
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
		file.WriteString(fmt.Sprintf("%d|`%s`|`%s`|`%d`\n", index+1, token.GetValue(), regexpsManager.GetRegexpsManager().GetType(token.GetSpecialChar()), token.GetKindCode()))
	}
}



// 只需初始化一次
func (la *LexicalAnalyzer) Init() {
	la.initNFAs()
	la.initFinalNFA()
}
func (la *LexicalAnalyzer) initNFAs() {
	for i:=0;i<len(la.lexicalConf.SpecialCharsOfNFAs);i++{
		specialChar := la.lexicalConf.SpecialCharsOfNFAs[i]
		nfa := stateMachine.NewNFABuilder(string(specialChar)).BuildNFA()
		nfa.EliminateBlankStates()
		nfa.MarkToken()
		la.nfas = append(la.nfas,nfa)
	}

}
func (la *LexicalAnalyzer) initFinalNFA() {
	la.finalNfa = stateMachine.NewNFA(regexpsManager.Eps)
	for i:=0;i<len(la.nfas);i++{
		la.finalNfa.AddParallelNFA(la.nfas[i])
	}
	la.finalNfa.EliminateBlankStates()
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
	for _, NFA := range la.nfas {
		NFA.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s 状态机.md", stateMachineStoreFileDir,string( NFA.GetSpecialChar())))
	}
	la.finalNfa.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s", stateMachineStoreFileDir, finalNFAFileName))

}

