package main

import (
	"LLONE"
	"file"
	"fmt"
	"testLay"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\conf.json`
const llOneFilePath=`C:\Users\hasee\Desktop\Go_Practice\编译器\conf\LL1`
func main() {
	stf := LLONE.NewStateTableFormer(llOneFilePath)
	stf.GetFirst()
	fmt.Println()
	stf.GetFollow()
	fmt.Println()
	stf.GetSelect()
	fmt.Println()
	stf.GetStateTable()
	fmt.Println()
	fmt.Println(stf.ShowStateTable())
}






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
func allTest() {
	testManager := testLay.TestManager{}
	for i := 0; i < len(testFilePaths); i++ {
		fmt.Printf("----------------------------------- 第 %d 个测试文件 -----------------------------------\n", i+1)
		//testManager.CloseTheOutputOfTestInformation()
		testManager.OpenTheOutputOfTestInformation()
		testManager.SetFileReader(file.NewFileReader(testFilePaths[i]))
		testManager.SetTestObject(testLay.NewNFATestUnit())
		if testManager.RepeatTest(1) == true{
			fmt.Println("--------------------------------------  测试通过  ----------------------------------------")
		}else{
			fmt.Println("--------------------------------------  出现错误  ----------------------------------------")
		}
	}
}
