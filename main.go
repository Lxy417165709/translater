package main

import (
	"file"
	"fmt"
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

var programFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\source2.md`

const (
	wordDelimiter        = "|"
	grammarUnitDelimiter = "->"
	grammarFilePath      = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\grammar.md`
)

func main() {
	//regexpsManager := regexpsManager.NewRegexpsManager(
	//	grammarFilePath,
	//	grammarUnitDelimiter,
	//	wordDelimiter,
	//)
	//regexpsManager.Init()
	//lexicalAnalyzer := lexical.NewLexicalAnalyzer(regexpsManager)
	//lexicalAnalyzer.Init()
	//pairs := lexicalAnalyzer.Parse(file.NewFileReader(programFilePath).GetFileBytes())
	//for index,pair := range pairs{
	//	fmt.Printf("[第 %d 个] (%s, %d, %v)\n",index+1,pair.GetKind(),pair.GetKindCode(),pair.GetValue())
	//}

	allTest()
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
