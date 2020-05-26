package test

import (
	"file"
	"fmt"
)

const (
	infoTemplateOfBeginningTest   = "----------------------------------- 第 %d 个测试文件 -----------------------------------\n"
	infoTemplateOfFileTestSuccess = "--------------------------------------  测试通过  ----------------------------------------\n"
	infoTemplateOfBFileTestFail   = "--------------------------------------  出现错误  ----------------------------------------\n"
	infoTemplateOfItemTestSuccess = "[第 %d 次测试] 正常\n"
	infoTemplateOfItemTestFail    = "[第 %d 次测试, 第 %d 个测试单元] 发生错误，没有通过测试的对象是：%v\n"
)

// TODO: 重构，这个名字命得不好，成员变量可以进行提取
// TODO: 非常烂的代码
type IsMatchOfNFATestTable struct {
	isMatchOfNFATestItems                  []*isMatchOfNFATestItem
	isNeedToShowTheOutputOfTestInformation bool
}

func NewIsMatchOfNFATestTable() *IsMatchOfNFATestTable {
	return &IsMatchOfNFATestTable{}
}
func (tm *IsMatchOfNFATestTable) CloseTheOutputOfTestInformation() {
	tm.isNeedToShowTheOutputOfTestInformation = false
}
func (tm *IsMatchOfNFATestTable) OpenTheOutputOfTestInformation() {
	tm.isNeedToShowTheOutputOfTestInformation = true
}

func (tm *IsMatchOfNFATestTable) RepeatTestOfMultipleFiles(testFilePaths []string, repeatTimes int) bool {

	for i := 0; i < len(testFilePaths); i++ {
		if tm.isNeedToShowTheOutputOfTestInformation {
			fmt.Println(tm.getOutPutOfBeginTesting(i))
		}

		testResult := true
		outputInformation := tm.getOutPutOfTestingFileSuccess()
		if tm.RepeatTestOfSingleFile(testFilePaths[i], repeatTimes) == false {
			outputInformation = tm.getOutPutOfTestingFileFail()
			testResult = false
		}


		if tm.isNeedToShowTheOutputOfTestInformation {
			fmt.Print(outputInformation)
		}
		if testResult == false {
			return false
		}
	}
	return true
}



func (tm *IsMatchOfNFATestTable) RepeatTestOfSingleFile(filePath string, repeatTimes int) bool {
	tm.testInit(filePath)
	for i := 0; i < repeatTimes; i++ {
		testResult := true
		outputInformation := tm.getOutputOfTestingItemSuccess(i+1)
		for index, isMatchOfNFATestItem := range tm.isMatchOfNFATestItems {
			if isMatchOfNFATestItem.Test() == false {
				outputInformation = tm.getOutputOfTestingItemFail(i+1,index+1,isMatchOfNFATestItem)
				testResult = false
				break
			}
		}
		if tm.isNeedToShowTheOutputOfTestInformation {
			fmt.Print(outputInformation)
		}
		if testResult == false {
			return false
		}
	}
	return true
}

func (tm *IsMatchOfNFATestTable)getOutPutOfBeginTesting(indexOfTestFile int) string{
	return fmt.Sprintf(infoTemplateOfBeginningTest, indexOfTestFile+1)
}
func (tm *IsMatchOfNFATestTable)getOutPutOfTestingFileSuccess() string{
	return infoTemplateOfFileTestSuccess
}
func (tm *IsMatchOfNFATestTable)getOutPutOfTestingFileFail() string{
	return infoTemplateOfBFileTestFail
}
func (tm *IsMatchOfNFATestTable)getOutputOfTestingItemSuccess(nowTimes int)  string{
	return fmt.Sprintf(infoTemplateOfItemTestSuccess, nowTimes)
}
func (tm *IsMatchOfNFATestTable)getOutputOfTestingItemFail(nowTimes int,indexOfTestingFile int,testingItem *isMatchOfNFATestItem)  string{
	return fmt.Sprintf(
		infoTemplateOfItemTestFail,
		nowTimes,
		indexOfTestingFile+1,
		testingItem,
	)
}




func (tm *IsMatchOfNFATestTable) testInit(filePath string) {
	tm.isMatchOfNFATestItems = make([]*isMatchOfNFATestItem, 0)
	itemContents := file.NewFileReader(filePath).GetFileLines()
	tm.parse(itemContents)
}
func (tm *IsMatchOfNFATestTable) parse(contents []string) {
	for _, content := range contents {
		item := NewIsMatchOfNFATestItem(content)
		tm.isMatchOfNFATestItems = append(tm.isMatchOfNFATestItems, item)
	}
}
