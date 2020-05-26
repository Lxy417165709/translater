package main

import (
	"conf"
	"file"
	"syntex"
	"test"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\conf.json`
const sourceFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\2_source.md`

func main() {
	conf.Init(configureFilePath)
	text := file.NewFileReader(sourceFilePath).GetFileBytes()
	syntaxParser := syntex.NewSyntaxParser()
	syntaxParser.GetSyntaxTree(text)
}


func allTest() {
	var testFilePaths = []string{
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest1.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest2.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest3.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest4.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest5.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest6.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest7.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest8.md`,
		`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest9.md`,
	}
	testManager := test.NewIsMatchOfNFATestTable()
	testManager.OpenTheOutputOfTestInformation()
	testManager.RepeatTestOfMultipleFiles(testFilePaths,1)
}


