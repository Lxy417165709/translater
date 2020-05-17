package main

import (
	"file"
	"fmt"
	"lexical"
	"regexpsManager"
	"testLay"
)

var testFilePaths = [...]string{
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
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest10.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest11.md`,
}

var programFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\source.md`
var nfaFilePath =  `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfa_Visualization_data\nfa.md`
const (
	wordDelimiter        = "|"
	grammarUnitDelimiter = "->"
	grammarFilePath      = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\grammar.md`
)

func main() {
	regexpsManager := regexpsManager.NewRegexpsManager(
		grammarFilePath,
		grammarUnitDelimiter,
		wordDelimiter,
	)
	regexpsManager.Init()

	lex := lexical.NewLexicalAnalyzer(regexpsManager)
	lex.Init()
	lex.ShowParsedTokens(file.NewFileReader(programFilePath).GetFileBytes())

	//stateMachine.NewNFABuilder("U",regexpsManager).BuildNotBlankStateNFA().OutputNFA(nfaFilePath)
	//stateMachine.NewNFABuilder("O",regexpsManager).BuildNotBlankStateNFA().OutputNFA(nfaFilePath )
	//stateMachine.NewNFABuilder("J",regexpsManager).BuildNotBlankStateNFA().OutputNFA(nfaFilePath )


}

func allTest() {
	regexpsManager := regexpsManager.NewRegexpsManager(
		grammarFilePath,
		grammarUnitDelimiter,
		wordDelimiter,
	)
	regexpsManager.Init()
	testManager := testLay.TestManager{}
	for i := 0; i < len(testFilePaths); i++ {
		fmt.Printf("----------------------------------- 第 %d 个测试文件 -----------------------------------\n", i+1)
		testManager.CloseTheOutputOfTestInformation()
		testManager.SetFileReader(file.NewFileReader(testFilePaths[i]))
		testManager.SetTestObject(testLay.NewNFATestUnit(regexpsManager))
		if testManager.RepeatTest(100) == true{
			fmt.Println("--------------------------------------  测试通过  ----------------------------------------")
		}else{
			fmt.Println("--------------------------------------  出现错误  ----------------------------------------")
		}
	}

}
