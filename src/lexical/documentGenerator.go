package lexical

import (
	"bytes"
	"conf"
	"file"
	"fmt"
	"grammar"
	"io/ioutil"
	"os"
	"strings"
)

// TODO: 这个更应该放在全局
const KB = 1024
const stateMachineFileInfoSizeLimit = 2*KB

type lexicalDocumentGenerator struct {
	lexicalConf *conf.LexicalConf
	content     bytes.Buffer
}

func NewlexicalDocumentGenerator(lexicalConf *conf.LexicalConf) *lexicalDocumentGenerator {
	return &lexicalDocumentGenerator{lexicalConf: lexicalConf}
}

func (mdo *lexicalDocumentGenerator) Generate() {
	mdo.constructContent()
	mdo.writeContentToStoreFile()
}

func (mdo *lexicalDocumentGenerator) constructContent() {
	mdo.content.WriteString("@[TOC]\n")
	mdo.content.WriteString("# 我是一个自动生成的MarkDown文件\n")
	mdo.constructContentPartOne()
	mdo.constructContentPartTwo()
	mdo.constructContentPartThree()
}
func (mdo *lexicalDocumentGenerator) writeContentToStoreFile() {
	var storeFile *os.File
	var err error
	if storeFile, err = os.Create(mdo.lexicalConf.DisplayDocumentPath); err != nil {
		panic(err)
	}
	defer storeFile.Close()

	if _, err = storeFile.Write(mdo.content.Bytes()); err != nil {
		panic(err)
	}
	fmt.Println("文档已自动生成！")
}

func (mdo *lexicalDocumentGenerator) constructContentPartOne() {

	mdo.content.WriteString("## 语法\n")
	lines := grammar.GetRegexpsManager().GetReformLinesOfGrammarFile()
	for _,line := range lines{
		mdo.content.WriteString(line + "\n")
	}
}

func (mdo *lexicalDocumentGenerator) constructContentPartTwo() {
	stateMachineStoreFileDir :=  mdo.lexicalConf.GetStoreDirPathOfStateMachine()
	fileInfos, err := ioutil.ReadDir(stateMachineStoreFileDir)
	if err != nil {
		panic(err)
	}
	mdo.content.WriteString("## 自动机\n")
	for _, fileInfo := range fileInfos {
		mdo.content.WriteString("### " + getTheFirstPartOfFileName(fileInfo.Name()) + "\n")
		if stateMachineIsTooLarge(fileInfo) {
			mdo.content.WriteString("状态机过于庞大\n")
			continue
		}
		stateMachineStoreFilePath := fmt.Sprintf("%s/%s", stateMachineStoreFileDir, fileInfo.Name())
		mdo.content.Write(file.NewFileReader(stateMachineStoreFilePath).GetFileBytes())
	}
}
func (mdo *lexicalDocumentGenerator) constructContentPartThree() {
	kindCodeFilePath := mdo.lexicalConf.GetStorePathOfKindCodes()
	tokensFilePath := mdo.lexicalConf.GetStorePathOfTokens()
	mdo.writeFileContent("种别码", kindCodeFilePath, false)
	mdo.writeFileContent("被识别的源代码", mdo.lexicalConf.SourceFilePath, true)
	mdo.writeFileContent("识别出的所有Token", tokensFilePath, false)
}
func (mdo *lexicalDocumentGenerator) writeFileContent(topic string, fileName string, isCode bool) {
	mdo.content.WriteString(fmt.Sprintf("## %s\n", topic))
	if isCode {
		mdo.content.WriteString("```go\n")
	}
	mdo.content.Write(file.NewFileReader(fileName).GetFileBytes())
	if isCode {
		mdo.content.WriteString("```\n")
	}
}


func stateMachineIsTooLarge(stateMachineFileInfo os.FileInfo) bool{
	return stateMachineFileInfo.Size()>stateMachineFileInfoSizeLimit
}
func getTheFirstPartOfFileName(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

