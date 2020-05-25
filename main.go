package main

import (
	"conf"
	"file"
	"fmt"
	"syntex"
	"test"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\conf.json`
const sourceFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\2_source.md`

func main() {
	conf.Init(configureFilePath)
	fmt.Println(conf.GetConf().SyntaxConf)
	text := file.NewFileReader(sourceFilePath).GetFileBytes()
	syntaxParser := syntex.NewSyntaxParser()
	syntaxParser.GetSyntaxTree(text)
}

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

func allTest(testFilePaths []string) {
	testTable := test.NewIsMatchOfNFATestTable()
	testTable.CloseTheOutputOfTestInformation()
	for i := 0; i < len(testFilePaths); i++ {
		fmt.Printf("----------------------------------- 第 %d 个测试文件 -----------------------------------\n", i+1)
		testTable.SetTestFile(testFilePaths[i])
		if testTable.RepeatTest(100) == true {
			fmt.Println("--------------------------------------  测试通过  ----------------------------------------")
		} else {
			fmt.Println("--------------------------------------  出现错误  ----------------------------------------")
		}
	}

}
