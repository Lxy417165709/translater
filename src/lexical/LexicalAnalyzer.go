package lexical

import (
	"conf"
	file2 "file"
	"fmt"
	"grammar"
	"os"
	"stateMachine"
)


func Init(lexicalConf *conf.LexicalConf) {
	globalLexicalAnalyzer.lexicalConf = lexicalConf
	globalLexicalAnalyzer.Init()
}
var globalLexicalAnalyzer = &LexicalAnalyzer{}

type LexicalAnalyzer struct {
	lexicalConf *conf.LexicalConf
	nfas        []*stateMachine.NFA
	finalNfa    *stateMachine.NFA
	lexicalDocumentGenerator *lexicalDocumentGenerator
}
func (la *LexicalAnalyzer) initLexicalDocumentGenerator() {
	la.lexicalDocumentGenerator = NewlexicalDocumentGenerator(la.lexicalConf)
}
func GetLexicalAnalyzer() *LexicalAnalyzer{
	return globalLexicalAnalyzer
}
func (la *LexicalAnalyzer) GetTokens(bytes []byte) []*grammar.Token {
	return la.finalNfa.GetTokens(string(bytes))
}

func (la *LexicalAnalyzer) formTheMarkdownFileOfTokens() {
	tokens := la.GetTokens(file2.NewFileReader(la.lexicalConf.SourceFilePath).GetFileBytes())
	if err := la.writeTokensToFile(tokens,la.getStorePathOfTokens());err!=nil{
		panic(err)
	}
}

func (la *LexicalAnalyzer) writeTokensToFile(tokens []*grammar.Token,filePath string) error{
	var file *os.File
	var err error
	if file, err = os.Create(filePath); err != nil{
		return err
	}
	defer file.Close()
	lines := la.changeTokensToFileLines(tokens)
	for _,line  := range lines {
		if _,err = file.WriteString(line);err!=nil{
			return err
		}
	}
	return nil
}
func (la *LexicalAnalyzer) changeTokensToFileLines(tokens []*grammar.Token) []string{
	lines := make([]string,0)
	lines = append(lines,"索引|值|类型|种别码\n")
	lines = append(lines,"--|--|--|--\n")
	for index, token := range tokens {
		lines = append(lines, fmt.Sprintf(
			"%d|`%v`|`%s`|`%d`\n",
			index+1,
			token.GetValue(),
			grammar.GetRegexpsManager().GetType(token.GetSpecialChar()),
			token.GetKindCode()),
		)
	}
	return lines

}


func (la *LexicalAnalyzer) formKindCodeFile() {
	file, err := os.Create(la.getStorePathOfKindCodes())
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("索引|单词|类别|种别码\n")
	file.WriteString("--|--|--|--\n")
	for index, token := range grammar.GetRegexpsManager().GetAllTokens() {
		file.WriteString(fmt.Sprintf("%d|`%s`|`%s`|`%d`\n", index+1, token.GetValue(), grammar.GetRegexpsManager().GetType(token.GetSpecialChar()), token.GetKindCode()))
	}
}

// 只需初始化一次
func (la *LexicalAnalyzer) Init() {
	la.initNFAs()
	la.initFinalNFA()
	la.initLexicalDocumentGenerator()
}
func (la *LexicalAnalyzer) initNFAs() {
	for i := 0; i < len(la.lexicalConf.SpecialCharsOfNFAs); i++ {
		specialChar := la.lexicalConf.SpecialCharsOfNFAs[i]
		nfa := stateMachine.NewNFABuilder(string(specialChar)).BuildNFA()
		la.nfas = append(la.nfas, nfa.EliminateBlankStates().MarkSpecialChar())
	}

}
func (la *LexicalAnalyzer) initFinalNFA() {
	la.finalNfa = stateMachine.NewNFA()
	for i := 0; i < len(la.nfas); i++ {
		la.finalNfa.AddParallelNFA(la.nfas[i])
	}
	la.finalNfa.EliminateBlankStates()
}

func (la *LexicalAnalyzer) getStorePathOfKindCodes() string {
	return fmt.Sprintf("%s/%s", la.lexicalConf.InformationDir, la.lexicalConf.FileNameOfStoringKindCodes)
}
func (la *LexicalAnalyzer) getStorePathOfTokens() string {
	return fmt.Sprintf("%s/%s", la.lexicalConf.InformationDir, la.lexicalConf.FileNameOfStoringTokens)
}
func (la *LexicalAnalyzer) getStoreDirPathOfStateMachine() string {
	return fmt.Sprintf("%s/%s", la.lexicalConf.InformationDir, la.lexicalConf.StateMachineDirName)
}


func (la *LexicalAnalyzer)FormLexicalDocument() {
	la.lexicalDocumentGenerator.Generate()
}
func (la *LexicalAnalyzer) FormLexicalFile() {
	la.formStateMachineFiles()
	la.formKindCodeFile()
	la.formTheMarkdownFileOfTokens()
}


func (la *LexicalAnalyzer) formStateMachineFiles() {
	storeDirPathOfStateMachine := la.getStoreDirPathOfStateMachine()
	if err := os.MkdirAll(storeDirPathOfStateMachine, 0666); err != nil {
		panic(err)
	}
	for _, NFA := range la.nfas {
		NFA.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s 状态机.md", storeDirPathOfStateMachine, string(NFA.GetSpecialChar())))
	}
	la.finalNfa.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s", storeDirPathOfStateMachine, la.lexicalConf.FileNameOfStoringFinalNFA))

}
