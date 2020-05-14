package testUnit

import (
	"fmt"
	"lexicalTest"
)

func ShowTestResult(filePath string) {
	testUnits := getTestUnits(filePath)
	for index, unit := range testUnits {
		if unit.test() == false {
			fmt.Printf("没通过第 %d 个测试单元: %v \n", index+1, *unit)
		} else {
			fmt.Printf("通过第 %d 个测试单元: %v \n", index+1, *unit)
		}
	}
}
func BuildGrammar(filePath string) {
	grammarUnits := getGrammarUnits(filePath)
	for index, unit := range grammarUnits {
		lexicalTest.GlobalNFAManager.Add(unit.identity,unit.regexp)
		fmt.Printf("添加了第 %d 个特殊字符：%s   对应的正则表达式是：%s\n",index,string(unit.identity),unit.regexp)
	}
}

