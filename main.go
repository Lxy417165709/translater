package main

import (
	"fmt"
	"testUnit"
)
var testFilePaths = []string{
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest1.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest2.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest3.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest4.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest5.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest6.md`,
}
var grammarFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\grammar.md`

func main() {
	testUnit.BuildGrammar(grammarFilePath)
	//testUnit.ShowTestResult(testUnitFilePath)
	allTest()
}

func allTest() {
	for index,testFilePath := range testFilePaths {
		fmt.Printf("----------------------------- 第 %d 个测试文件-----------------------------\n",index + 1)
		testUnit.ShowTestResult(testFilePath)
		fmt.Println()
	}
}



