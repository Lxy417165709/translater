package lexical

import (
	"bytes"
	"conf"
	"file"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	KB = 1024
)

type MarkDownObject struct {
	storeFileName string
	content bytes.Buffer
}

func NewMarkDownObject(storeFileName string) *MarkDownObject {
	return &MarkDownObject{storeFileName: storeFileName}
}

func (mdo *MarkDownObject) Generate() {
	mdo.constructContent()
	mdo.writeContentToStoreFile()
}

func (mdo *MarkDownObject) constructContent() {
	mdo.content.WriteString("@[TOC]\n")
	mdo.content.WriteString("# 我是一个自动生成的MarkDown文件\n")
	mdo.constructContentPartOne()
	mdo.constructContentPartTwo()
	mdo.constructContentPartThree()
}
func (mdo *MarkDownObject) writeContentToStoreFile() {
	storeFile, err := os.Create(mdo.storeFileName)
	defer storeFile.Close()
	if err != nil {
		panic(err)
	}
	storeFile.Write(mdo.content.Bytes())
	fmt.Println("文档已自动生成！")
}

func (mdo *MarkDownObject) constructContentPartOne() {
	mdo.writeFileContent("语法", conf.GetConf().GrammarConf.FilePath, true)
}
func (mdo *MarkDownObject) constructContentPartTwo() {
	stateMachineStoreFileDir := fmt.Sprintf("%s/%s", conf.GetConf().LexicalInformationDir, stateMachineDirName)
	fileInfos, err := ioutil.ReadDir(stateMachineStoreFileDir)
	if err != nil {
		panic(err)
	}
	mdo.content.WriteString("## 自动机\n")
	for _, fileInfo := range fileInfos {
		mdo.content.WriteString("### " + getTheFirstPartOfFileName(fileInfo.Name()) + "\n")
		if fileInfo.Size() > 2*KB {
			mdo.content.WriteString("状态机过于庞大\n")
			continue
		}
		stateMachineStoreFilePath := fmt.Sprintf("%s/%s", stateMachineStoreFileDir, fileInfo.Name())
		mdo.content.Write(file.NewFileReader(stateMachineStoreFilePath).GetFileBytes())
	}
}
func (mdo *MarkDownObject) constructContentPartThree() {
	kindCodeFilePath :=  fmt.Sprintf("%s/%s",conf.GetConf().LexicalInformationDir,kindCodeFileName)
	tokensFilePath := fmt.Sprintf("%s/%s",conf.GetConf().LexicalInformationDir,tokensFileName)
	mdo.writeFileContent("种别码", kindCodeFilePath, false)
	mdo.writeFileContent("被识别的源代码", conf.GetConf().SourceFilePath, true)
	mdo.writeFileContent("识别出的所有Token",tokensFilePath, false)
}
func (mdo *MarkDownObject) writeFileContent(topic string, fileName string, isCode bool) {
	mdo.content.WriteString(fmt.Sprintf("## %s\n", topic))
	if isCode {
		mdo.content.WriteString("```go\n")
	}
	mdo.content.Write(file.NewFileReader(fileName).GetFileBytes())
	if isCode {
		mdo.content.WriteString("```\n")
	}
}

func getTheFirstPartOfFileName(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}
