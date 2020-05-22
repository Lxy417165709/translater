package main

import (
	"conf"
	"grammar"
	"machine"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\conf.json`
const tempPath =`C:\Users\hasee\Desktop\Go_Practice\编译器\doc\gen\temp.md`
func main() {
	conf.Init(configureFilePath)
	//testTable := test.NewIsMatchOfNFATestTable(
	//	&conf.GetConf().IsMatchOfNFATestConf,&conf.GetConf().GrammarConf,
	//)
	//testTable.OpenTheOutputOfTestInformation()
	//testTable.RepeatTest(1)

	nfaBuilder := machine.NewNFABuilder(grammar.NewSpecialCharTable(&conf.GetConf().GrammarConf))
	nfa := nfaBuilder.BuildNFAByWord("2")
	if err := nfa.StoreMermaidGraphOfThisNFA(tempPath);err!=nil{
		panic(err)
	}

}




// result := nfaBuilder.BuildNFAByWord("2$").EliminateBlankStates().IsMatch("2")
// 这个测试用例有问题




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
}
// 测试也要进行配置！
//func allTest() {
//	testManager := testLay.TestManager{}
//	for i := 0; i < len(testFilePaths); i++ {
//		fmt.Printf("----------------------------------- 第 %d 个测试文件 -----------------------------------\n", i+1)
//		//testManager.CloseTheOutputOfTestInformation()
//		testManager.OpenTheOutputOfTestInformation()
//		testManager.SetFileReader(file.NewFileReader(testFilePaths[i]))
//		testManager.SetTestObject(testLay.NewNFATestUnit())
//		if testManager.RepeatTest(1) == true{
//			fmt.Println("--------------------------------------  测试通过  ----------------------------------------")
//		}else{
//			fmt.Println("--------------------------------------  出现错误  ----------------------------------------")
//		}
//	}
//}
