package test

import (
	"fmt"
	"io/ioutil"
	"os"
)

type TestableDirectory struct {
	path string
	tables []Testable
	testType TypeOfTest
}

func NewTestableDirectory(path string,testType TypeOfTest) *TestableDirectory {
	return &TestableDirectory{path:path,testType:testType}
}

func (td *TestableDirectory)GetTypeOfTest() TypeOfTest{
	return td.testType
}

func (td *TestableDirectory) Test() bool {
	td.testInit()
	for _,table := range td.tables{
		if !table.Test(){
			return false
		}
	}
	return true
}

func (td *TestableDirectory) GetErrMsg() string{
	for index, table := range td.tables {
		if table.Test() == false {
			return fmt.Sprintf(
				"[第 %d 个测试文件]存在错误\n	%s",
				index+1,
				table.GetErrMsg(),
			)
		}
	}
	panic("测试无错误，这的代码不应该执行")
}

func (td *TestableDirectory) testInit() {
	infos,err := ioutil.ReadDir(td.path)
	if err!=nil{
		panic(err)
	}
	td.parse(infos)
}

func (td *TestableDirectory) parse(infos []os.FileInfo) {
	for _,info := range infos{
		td.tables = append(td.tables,
			NewTestableTable(
				fmt.Sprintf(`%s\%s`,td.path,info.Name()),
				td.testType,
			),
		)
	}
}
