package file

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	basePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\aboutLexicalAnalyzer\`
	sonNFAPath = basePath + `sonNFA\`
	filePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\aboutLexicalAnalyzer\test.md`
)

type MarkDownObject struct{
	content bytes.Buffer
}
func NewMarkDownObject() *MarkDownObject{
	return &MarkDownObject{}
}
func (mdo *MarkDownObject) Write() {
	mdo.content.WriteString("# 我是一个自动生成的MarkDown文件\n")
	mdo.WritePartOne()
	mdo.WritePartTwo()
	mdo.WritePartThree()
	targetFile,err := os.Create(filePath)
	defer targetFile.Close()
	if err!=nil{
		panic(err)
	}
	targetFile.Write(mdo.content.Bytes())

}
func (mdo *MarkDownObject)WritePartOne() {
	mdo.Foo("语法","1_grammar.md",true)
}


func (mdo *MarkDownObject)WritePartTwo() {

	fileInfos, err := ioutil.ReadDir(sonNFAPath)
	if err!=nil {
		panic(err)
	}
	mdo.content.WriteString("## 自动机\n")
	for _,fileInfo := range fileInfos{
		mdo.content.WriteString("### " + fileInfo.Name() + "\n")
		if fileInfo.Size()>2000{
			mdo.content.WriteString("状态机过于庞大\n")
		}else{
			mdo.content.Write(NewFileReader(sonNFAPath + fileInfo.Name()).GetFileBytes())
		}
	}
}

func (mdo *MarkDownObject) WritePartThree() {
	mdo.Foo("种别码","code.md",false)
	mdo.Foo("被识别的源代码","2_source.md",true)
	mdo.Foo("识别出的所有Token","tokens.md",false)
}

func (mdo *MarkDownObject)Foo(topic string,fileName string,isCode bool) {
	mdo.content.WriteString(fmt.Sprintf("## %s\n",topic))
	if isCode {
		mdo.content.WriteString("```go\n")
	}
	mdo.content.Write(NewFileReader(basePath + fileName).GetFileBytes())
	if isCode {
		mdo.content.WriteString("```\n")
	}
}
