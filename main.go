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
	allTest()
	/*
		D -> 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
		L -> a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r|s |t |u|v|w|x|y|z
		T -> D | L
		W -> if | while | for | int | else
		O -> > | < | >= | <= | = | ==
		Z -> D+
		I -> LT*
	*/

	// TODO：加标记

	//testUnit.BuildGrammar(grammarFilePath)
	//g := lexicalTest.NewNFABuilder("G").BuildDFA()
	//g.MarkDown('G')
	//d := lexicalTest.NewNFABuilder("D").BuildDFA()
	//d.MarkDown('D')
	//g.AddParallelNFA(d)
	////g.Merge()
	//g.Show()
	////g.ChangeToDFA()
	////g.Show()
	//if !g.IsDFA(){
	//	fmt.Println("dfa不是dfa")
	//}

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

func allTest() {
	for index, testFilePath := range testFilePaths {
		fmt.Printf("----------------------------- 第 %d 个测试文件-----------------------------\n", index+1)
		testUnit.ShowTestResult(testFilePath)
		fmt.Println()
	}
}
