package test

import (
	"fmt"
	"io/ioutil"
	"os"
)

// TODO: 命名有误
type TestDir struct {
	dirPath string
	tables []Testable
	testType TestName
}

func NewTestDir(dirPath string,testType TestName) *TestDir {
	return &TestDir{dirPath:dirPath,testType:testType}
}

func (td *TestDir)GetTestType() TestName{
	return td.testType
}
func (td *TestDir) Test() bool {
	td.testInit()
	for _,table := range td.tables{
		if !table.Test(){
			return false
		}
	}
	return true
}
func (td *TestDir) GetErrMsg() string{
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

func (td *TestDir) testInit() {
	infos,err := ioutil.ReadDir(td.dirPath)
	if err!=nil{
		panic(err)
	}
	td.parse(infos)
}
func (td *TestDir) parse(infos []os.FileInfo) {
	for _,info := range infos{
		td.tables = append(td.tables,
			NewTestTable(
				fmt.Sprintf(`%s\%s`,td.dirPath,info.Name()),
				td.testType,
			),
		)
	}
}
