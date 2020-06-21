package test

import (
	"file"
	"fmt"
)

type TestableTable struct {
	filePath  string
	testItems []Testable
	testType TypeOfTest
}

func NewTestableTable(filePath string, testType TypeOfTest) *TestableTable {
	return &TestableTable{filePath: filePath, testType: testType}
}

func (tm *TestableTable) Test() bool {
	tm.testInit()
	for _, item := range tm.testItems {
		if item.Test() == false {
			return false
		}
	}
	return true
}

func (tm *TestableTable) GetErrMsg() string {
	for index, item := range tm.testItems {
		if item.Test() == false {
			return fmt.Sprintf(
				"[第 %d 个测试样例]存在错误\n	%s",
				index+1,
				item.GetErrMsg(),
			)
		}
	}
	panic("测试无错误，这的代码不应该执行")
}

func (tm *TestableTable) testInit() {
	itemContents := file.NewFileReader(tm.filePath).GetFileLines()
	tm.parse(itemContents)
}

func (tm *TestableTable) parse(contents []string) {
	if !globalFactory.NewFunctionIsExist(tm.testType) {
		panic(fmt.Sprintf("不存在 %s 的测试类型",tm.testType))
	}
	for _, content := range contents {
		newFunction := globalFactory.GetCreateFunction(tm.testType)
		item := newFunction(content)
		tm.testItems = append(tm.testItems, item)
	}
}
