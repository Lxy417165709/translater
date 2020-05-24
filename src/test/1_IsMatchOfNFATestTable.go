package test

import (
	"conf"
	"file"
	"fmt"
	"machine"
)

type IsMatchOfNFATestTable struct {
	cf *conf.Conf
	isMatchOfNFATestConf *conf.IsMatchOfNFATestTableConf
	grammarConf *conf.GrammarConf
	isMatchOfNFATestItems []*isMatchOfNFATestItem
	isNeedToShowTheOutputOfTestInformation bool
}

func NewIsMatchOfNFATestTable(cf *conf.Conf) *IsMatchOfNFATestTable{
	return &IsMatchOfNFATestTable{
		cf:cf,
	}
}

func (tm *IsMatchOfNFATestTable)SetTestFile(filePath string) {
	tm.isMatchOfNFATestItems = make([]*isMatchOfNFATestItem,0)
	itemContents := file.NewFileReader(filePath).GetFileLines()
	tm.Parse(itemContents)
}

func (tm *IsMatchOfNFATestTable)Parse(contents []string) {

	nfaBuilder := machine.NewNFABuilder(tm.cf)
	for _,content := range contents{
		item := NewIsMatchOfNFATestItem(
			nfaBuilder ,
			tm.isMatchOfNFATestConf.DelimiterOfPieces,
			tm.grammarConf.DelimiterOfWords,
			content,
		)
		tm.isMatchOfNFATestItems = append(tm.isMatchOfNFATestItems,item)
	}
}


// 可以封装为对象返回
func (tm *IsMatchOfNFATestTable) RepeatTest(repeatTimes int) bool {
	for i := 0; i < repeatTimes; i++ {
		for index,isMatchOfNFATestItem := range tm.isMatchOfNFATestItems{
			if isMatchOfNFATestItem.Test() == false  {
				if tm.isNeedToShowTheOutputOfTestInformation{
					fmt.Printf("[第 %d 次测试, 第 %d 个测试单元] 发生错误，没有通过测试的对象是：%v\n",i+1,index+1,isMatchOfNFATestItem)
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

func (tm *IsMatchOfNFATestTable) CloseTheOutputOfTestInformation() {
	tm.isNeedToShowTheOutputOfTestInformation = false
}
func (tm *IsMatchOfNFATestTable) OpenTheOutputOfTestInformation() {
	tm.isNeedToShowTheOutputOfTestInformation = true
}

