package testLay

import (
	"file"
	"fmt"
	"interface"
)

type TestManager struct {
	fileReader *file.FileReader
	testObject _interface.TestableObject
	isNeedToShowTheOutputOfTestInformation bool
}
func (tm *TestManager) CloseTheOutputOfTestInformation() {
	tm.isNeedToShowTheOutputOfTestInformation = false
}
func (tm *TestManager) OpenTheOutputOfTestInformation() {
	tm.isNeedToShowTheOutputOfTestInformation = true
}

func (tm *TestManager) SetFileReader(fileReader *file.FileReader) {
	tm.fileReader = fileReader
}

func (tm *TestManager) SetTestObject(testObject _interface.TestableObject) {
	tm.testObject = testObject
}

// 可以封装为对象返回
func (tm *TestManager) RepeatTest(repeatTimes int) bool {
	for i := 0; i < repeatTimes; i++ {
		lines := tm.fileReader.GetFileLines()
		for index, line := range lines {
			tm.testObject.Parse(line)
			if tm.testObject.Test() == false {
				if tm.isNeedToShowTheOutputOfTestInformation{
					fmt.Printf("[第 %d 次测试, 第 %d 个测试单元] 发生错误，没有通过测试的对象是：%v\n",i+1,index+1,tm.testObject)
				}
				return false
			}
		}
		if tm.isNeedToShowTheOutputOfTestInformation{
			fmt.Printf("[第 %d 次测试] 正常\n",i+1)
		}
	}
	return true
}
