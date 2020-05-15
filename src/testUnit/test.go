package testUnit

import (
	"fmt"
	"lexicalTest"
)

func ShowTestResult(filePath string) {
	testUnits := getTestUnits(filePath)

	//fmt.Println("------------------------------------ NFA -------------------------------------")
	testType := "NFA"
	for index, unit := range testUnits {
		//fmt.Printf("[%s] 第 %d 个测试单元: %v \n", testType,index+1,*unit)
		if !unit.nfaTest() {
			//fmt.Printf("[%s] 没通过第 %d 个测试单元: %v \n", testType,index+1, *unit)
			panic(fmt.Sprintf("[%s] 没通过第 %d 个测试单元: %v \n", testType,index+1,*unit))

		} else {
			//fmt.Printf("通过第 %d 个测试单元: %v \n", index+1, *unit)
		}
	}
	//fmt.Println("------------------------------------ DFA -------------------------------------")
	testType = "DFA"
	for index, unit := range testUnits {
		//fmt.Printf("[%s] 第 %d 个测试单元: %v \n", testType,index+1,*unit)
		if !unit.dfaTest() {
			//fmt.Printf("[%s] 没通过第 %d 个测试单元: %v \n", testType,index+1, *unit)
			panic(fmt.Sprintf("[%s] 没通过第 %d 个测试单元: %v \n", testType,index+1,*unit))
		} else {
			//fmt.Printf("通过第 %d 个测试单元: %v \n", index+1, *unit)
		}
	}
	//fmt.Println("--------------------------------- 测试完成 -------------------------------------")
}




func BuildGrammar(filePath string) {
	grammarUnits := getGrammarUnits(filePath)
	for index, unit := range grammarUnits {
		lexicalTest.GlobalRegexpsManager.AddSpecialChar(unit.specialChar, unit.regexp)
		fmt.Printf("添加了第 %d 个特殊字符：%s   对应的正则表达式是：%s\n",index,string(unit.specialChar),unit.regexp)
	}
}

