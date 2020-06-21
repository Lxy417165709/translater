package main

import (
	"conf"
	"fmt"
	"syntex"
	"test"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\translater\conf\conf.json`
const testConfigureFilePath = `C:\Users\hasee\Desktop\Go_Practice\translater\conf\conf.json`
const sourceFilePath = `C:\Users\hasee\Desktop\Go_Practice\translater\conf\2_source.md`

func main() {
	conf.Init(configureFilePath)
	p :=syntex.NewSyntaxParser()
	p.GetSyntaxTree([]byte("8*1-5+10/5*5"))
	p.ShowSyntaxTree()
	fmt.Println(p.GetSyntaxTreeResult())
}


func allTest() {
	conf.Init(testConfigureFilePath)
	testObj1 := test.NewTestableDirectory(conf.GetConf().IsMatchOfNFATestConf.TestFilePath,test.NFATest)
	//testObj2 := test.NewTestableDirectory(conf.GetConf().SyntaxAnalyzerTestConf.TestFilePath,test.SyntaxTest)
	manager := test.NewManager()
	manager.Register(testObj1)
	manager.BeginTest()
}


