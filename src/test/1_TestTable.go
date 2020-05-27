package test

import (
	"file"
	"fmt"
)

// TODO: 重构，这个名字命得不好，成员变量可以进行提取
// TODO: 非常烂的代码
type TestTable struct {
	filePath  string
	testItems []Testable
	testType TestName
}

func NewTestTable(filePath string, testType TestName) *TestTable {
	return &TestTable{filePath: filePath, testType: testType}
}

func (tm *TestTable) Test() bool {
	tm.testInit()
	for _, isMatchOfNFATestItem := range tm.testItems {
		if isMatchOfNFATestItem.Test() == false {
			return false
		}
	}
	return true
}

func (tm *TestTable) GetErrMsg() string {
	for index, isMatchOfNFATestItem := range tm.testItems {
		if isMatchOfNFATestItem.Test() == false {
			return fmt.Sprintf(
				"[第 %d 个测试样例]存在错误\n	%s",
				index+1,
				isMatchOfNFATestItem.GetErrMsg(),
			)
		}
	}
	panic("测试无错误，这的代码不应该执行")
}

func (tm *TestTable) testInit() {
	itemContents := file.NewFileReader(tm.filePath).GetFileLines()
	tm.parse(itemContents)
}

func (tm *TestTable) parse(contents []string) {
	if tm.testType == NFATest {
		for _, content := range contents {
			item := NewNFATestItem(content)
			tm.testItems = append(tm.testItems, item)
		}
	}
	if tm.testType == SyntaxTest {
		for _, content := range contents {
			item := NewSyntaxAnalyzerTestItem(content)
			tm.testItems = append(tm.testItems, item)
		}
	}
}
