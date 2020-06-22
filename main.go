package main

import (
	"bufio"
	"conf"
	"fmt"
	"os"
	"syntex"
	"test"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\translater\conf\conf.json`
const testConfigureFilePath = `C:\Users\hasee\Desktop\Go_Practice\translater\conf\conf.json`
const sourceFilePath = `C:\Users\hasee\Desktop\Go_Practice\translater\conf\2_source.md`

func main() {
	conf.Init(configureFilePath)
	p :=syntex.NewSyntaxParser()

	defer func () {
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()

	var code string
	for {
		fmt.Print("> ")
		var codePart string
		Scanf(&codePart)
		if len(codePart)!=0{
			code = code+codePart
			continue
		}


		p.GetSyntaxTree([]byte(code))
		p.ParseSyntaxTree()
		code=""
	}
}


func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}


func allTest() {
	conf.Init(testConfigureFilePath)
	testObj1 := test.NewTestableDirectory(conf.GetConf().IsMatchOfNFATestConf.TestFilePath,test.NFATest)
	//testObj2 := test.NewTestableDirectory(conf.GetConf().SyntaxAnalyzerTestConf.TestFilePath,test.SyntaxTest)
	manager := test.NewManager()
	manager.Register(testObj1)
	manager.BeginTest()
}


