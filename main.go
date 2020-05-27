package main

import (
	"conf"
	"test"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\conf.json`
const sourceFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\2_source.md`

func main() {
	conf.Init(configureFilePath)
	//syntaxParser := syntex.NewSyntaxParser()
	//syntaxParser.GetSyntaxTree([]byte("jkjk+1274+oowie127789+898.457+wekksad/(577)+koaowo12898-jcweq"))
	allTest()
}


func allTest() {
	// TODO: 这个可以单独进行配置
	testObj1 := test.NewTestDir(conf.GetConf().IsMatchOfNFATestConf.TestFilePath,test.NFATest)
	testObj2 := test.NewTestDir(conf.GetConf().SyntaxAnalyzerTestConf.TestFilePath,test.SyntaxTest)
	manager := test.NewManager()
	manager.Register(testObj1,testObj2)
	manager.BeginTest()
}


