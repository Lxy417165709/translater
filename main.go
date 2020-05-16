package main

import (
	"fmt"
	"grammar"
	"io/ioutil"
	"os"
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
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest7.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest8.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest9.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest10.md`,
	`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\nfaTestFile\nfaGraphTest11.md`,
}

var programFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\source2.md`
func main() {

	grammar.BuildGrammar()
	//stateMachine.NewNFABuilder("a|b").BuildDFA().Show()

	allTest(100)
	fmt.Println("测试通过！")
	//fmt.Println(lexical.GlobalLexicalAnalyzer.Parse(getProgramData()))
}

func getProgramData() []byte{
	file,err := os.Open(programFilePath)
	if err!=nil{
		panic(err)
	}
	bytes,err := ioutil.ReadAll(file)
	if err!=nil{
		panic(err)
	}
	return bytes
	//result := g.Get(string( getProgramData()))
	//for i:=0;i<len(result);i++{
	//	fmt.Println(i,result[i])
	//}
}

func allTest(testTimes int) {
	for i:=0;i<testTimes;i++{
		for _, testFilePath := range testFilePaths {
			//fmt.Printf("----------------------------- 第 %d 个测试文件-----------------------------\n", index+1)
			testUnit.ShowTestResult(testFilePath)
			//fmt.Println()
		}
	}
}

