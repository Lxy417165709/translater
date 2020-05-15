package main

import (
	"fmt"
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
var grammarFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\grammar.md`
var programFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\doc\source2.md`
func main() {

	testUnit.BuildGrammar(grammarFilePath)

	allTest(20)
	fmt.Println("测试通过！")
	//dfa := lexicalTest.NewNFABuilder("N").BuildDFA()
	//dfa.Show()
	//fmt.Println(dfa.Get("0.00 0.00 0000 00000.004 00.5"))

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

func foo() {
	//for i:=0;i<20;i++{
	//	w := lexicalTest.NewNFABuilder("d$d$").BuildNFA()
	//	fmt.Println("nfa:")
	//	w.Show()
	//	if w.IsMatch("")==false{
	//		panic(fmt.Sprintf("[nfa] 第%d 次，发生错误！",i+1))
	//	}
	//	w.Merge()
	//	fmt.Println("nfa merge:")
	//	w.Show()
	//	if w.IsMatch("")==false{
	//		panic(fmt.Sprintf("[nfa merge]第%d 次，发生错误！",i+1))
	//	}
	//	w.ChangeToDFA()
	//	fmt.Println("dfa:")
	//	w.Show()
	//	if w.IsMatch("")==false{
	//		panic(fmt.Sprintf("[dfa] 第%d 次，发生错误！",i+1))
	//	}
	//	//fmt.Println(w.IsMatch("int"))
	//}

}
