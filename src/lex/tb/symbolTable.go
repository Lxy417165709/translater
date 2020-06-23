package tb

import (
	"env"
	"fmt"
	"utils"
)

// 名字还可以更好，不过暂时这样先
type SymbolTable struct {
	conf             *env.SymbolTableConf
	symbolTableItems []*SymbolTableItem

	specialCharToTerminatorReference map[byte]string
}

// 读表过程其实有点模板的味道了
func NewSymbolTable(conf *env.SymbolTableConf) (*SymbolTable, error) {
	table := &SymbolTable{}
	lines, err := utils.GetFileLines(conf.FilePath)
	if err != nil {
		return nil, err
	}
	if len(lines) <= 2 {
		return table, nil
	}

	lines = lines[2:] // 前2行是表格标头
	for index, line := range lines {
		item, err := NewSymbolTableItem(line, conf.DelimiterOfPieces)
		if err != nil {
			return nil, fmt.Errorf("%s 路径，读取第 %d 行时发生错误，%s",
				conf.FilePath,
				index+1,
				err.Error(),
			)
		}
		table.symbolTableItems = append(table.symbolTableItems, item)
	}
	//for i:=0;i<len(table.symbolTableItems);i++{
	//	fmt.Println(*table.symbolTableItems[i])
	//}
	table.initSpecialCharToTerminatorReference()
	return table, nil
}

func (st *SymbolTable)initSpecialCharToTerminatorReference() {
	st.specialCharToTerminatorReference = make(map[byte]string)
	for _,item := range st.symbolTableItems{
		st.specialCharToTerminatorReference[item.SpecialChar] = item.TerminatorReference
	}
}

func (st *SymbolTable)GetSpecialCharsOfCreatingNFA() []byte{
	specialCharOfCreatingNFA := make([]byte,0)
	for _,item := range st.symbolTableItems{
		specialCharOfCreatingNFA = append(specialCharOfCreatingNFA, item.SpecialChar)
	}
	return specialCharOfCreatingNFA
}

func (st *SymbolTable)GetTerminator(specialChar byte, value string) string {
	terminatorReference := st.specialCharToTerminatorReference[specialChar]
	if terminatorReference == ""{
		return value
	}
	return terminatorReference
}
