package lexical

import (
	"file"
	"fmt"
	"os"
	"sort"
	"stateMachine"
)

func (la *LexicalAnalyzer) FormLexicalFile() {
	la.formStateMachineFiles()
	la.formKindCodeFile()
	la.formTheMarkdownFileOfTokens()
}
func (la *LexicalAnalyzer) formStateMachineFiles() {
	storeDirPathOfStateMachine := la.lexicalConf.GetStoreDirPathOfStateMachine()
	if err := os.MkdirAll(storeDirPathOfStateMachine, 0666); err != nil {
		panic(err)
	}
	for _, NFA := range la.nfas {
		NFA.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s 状态机.md", storeDirPathOfStateMachine, string(NFA.GetSpecialChar())))
	}
	la.finalNfa.FormTheMermaidGraphOfNFA(fmt.Sprintf("%s/%s", storeDirPathOfStateMachine, la.lexicalConf.FileNameOfStoringFinalNFA))

}
func (la *LexicalAnalyzer) formKindCodeFile() {
	file, err := os.Create(la.lexicalConf.GetStorePathOfKindCodes())
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("索引|单词|类别|种别码\n")
	file.WriteString("--|--|--|--\n")

	tokens := la.finalNfa.GetAllFiltratedTokens()
	sort.Slice(tokens, func(i,j int) bool{
		return tokens[i].GetKindCode()<tokens[j].GetKindCode()
	})

	for index, token := range tokens {
		file.WriteString(fmt.Sprintf("%d|`%s`|`%s`|`%d`\n", index+1, token.GetValue(), token.GetType(), token.GetKindCode()))
	}
}
func (la *LexicalAnalyzer) formTheMarkdownFileOfTokens() {
	tokens := la.GetTokens(file.NewFileReader(la.lexicalConf.SourceFilePath).GetFileBytes())
	if err := la.writeTokensToFile(tokens,la.lexicalConf.GetStorePathOfTokens());err!=nil{
		panic(err)
	}
}
func (la *LexicalAnalyzer) GetTokens(bytes []byte) []*stateMachine.Token {
	la.finalNfa.SetPattern(bytes)
	return la.finalNfa.ParsePattern()
}

func (la *LexicalAnalyzer)FormLexicalDocument() {
	la.lexicalDocumentGenerator.Generate()
}

func (la *LexicalAnalyzer) writeTokensToFile(tokens []*stateMachine.Token,filePath string) error{
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
func (la *LexicalAnalyzer) changeTokensToFileLines(tokens []*stateMachine.Token) []string{
	lines := make([]string,0)
	lines = append(lines,"索引|值|类型|种别码\n")
	lines = append(lines,"--|--|--|--\n")
	for index, token := range tokens {
		lines = append(lines, fmt.Sprintf(
			"%d|`%v`|`%s`|`%d`\n",
			index+1,
			token.GetValue(),
			token.GetType(),
			token.GetKindCode()),
		)
	}
	return lines

}
