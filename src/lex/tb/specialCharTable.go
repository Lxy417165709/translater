package tb

import (
	"env"
	"fmt"
	"utils"
)

// 这里可以开启测试，就是看看得到的表内容和文件内容是否一样
type SpecialCharTable struct {
	conf             *env.SpecialCharTableConf
	specialCharItems []*SpecialCharItem

	specialCharToRegexp map[byte]*Regexp
}

func NewDefaultSpecialCharTable() (*SpecialCharTable, error) {
	conf := &env.SpecialCharTableConf{
		SpecialCharOfMarchMoreThanZeroTimes: "*",
		SpecialCharOfMarchMoreThanOnce:      "+",
	}
	return NewSpecialCharTable(conf)
}

func NewSpecialCharTable(conf *env.SpecialCharTableConf) (*SpecialCharTable, error) {
	table := &SpecialCharTable{}
	defer func() {
		table.conf = conf
		table.initSpecialCharToRegexp()
	}()

	if conf.FilePath == "" {
		return table, nil
	}

	lines, err := utils.GetFileLines(conf.FilePath)
	if err != nil {
		return nil, err
	}
	if len(lines) <= 2 {
		return table, nil
	}

	lines = lines[2:] // 前2行是表格标头
	for index, line := range lines {
		item, err := NewSpecialCharItem(line, conf.DelimiterOfPieces, conf.DelimiterOfWords)
		if err != nil {
			return nil, fmt.Errorf("%s 路径，读取第 %d 行时发生错误，%s",
				conf.FilePath,
				index+1,
				err.Error(),
			)
		}
		table.specialCharItems = append(table.specialCharItems, item)
	}
	return table, nil
}

func (sct *SpecialCharTable) initSpecialCharToRegexp() {
	sct.specialCharToRegexp = make(map[byte]*Regexp)
	for _, item := range sct.specialCharItems {
		sct.specialCharToRegexp[item.SpecialChar] = item.Regexp
	}
}

func (sct *SpecialCharTable) GetSpecialCharOfMarchMoreThanOnce() byte {
	return sct.conf.SpecialCharOfMarchMoreThanOnce[0]
}
func (sct *SpecialCharTable) GetSpecialCharOfMarchMoreThanZeroTimes() byte {
	return sct.conf.SpecialCharOfMarchMoreThanZeroTimes[0]
}
func (sct *SpecialCharTable) IsAdditionalChar(char byte) bool {
	return char == sct.GetSpecialCharOfMarchMoreThanOnce() || char == sct.GetSpecialCharOfMarchMoreThanZeroTimes()
}

func (sct *SpecialCharTable) CharIsSpecial(char byte) bool {
	return sct.specialCharToRegexp[char] != nil
}
func (sct *SpecialCharTable) GetRegexp(specialChar byte) *Regexp {
	return sct.specialCharToRegexp[specialChar]
}
