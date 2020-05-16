package testUnit

import (
	"file"
	"fmt"
)

func ShowTestResult(filePath string) {
	lines := file.NewFileReader(filePath).GetFileLines()
	//fmt.Println("------------------------------------ NFA -------------------------------------")
	testType := "NFA"
	for index, line := range lines {
		unit := NewTestUnit("","",false)
		unit.Parse(line)
		//fmt.Println(line,unit)
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
	for index, line := range lines {
		unit := NewTestUnit("","",false)
		unit.Parse(line)
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






