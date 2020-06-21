package char

import (
	"conf"
	"fmt"
	"strconv"
	"strings"
)

type specialCharItem struct {
	specialChar       byte
	_type             string
	kindCodeFlag      int
	regexp            *Regexp
}

func NewSpecialCharItem(specialCharItemLine string) *specialCharItem {
	sc := &specialCharItem{}
	sc.Parse(specialCharItemLine)
	return sc
}

func (sc *specialCharItem) Parse(content string) {
	content = strings.TrimSpace(content)
	pieces := strings.Split(strings.TrimSpace(content), conf.GetConf().GrammarConf.DelimiterOfPieces)
	if len(pieces) != 4 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于4", pieces))
	}
	sc.specialChar = strings.TrimSpace(pieces[0])[0]
	sc._type = strings.TrimSpace(pieces[1])

	var err error
	if sc.kindCodeFlag, err = strconv.Atoi(pieces[2]); err != nil {
		panic(err)
	}

	sc.regexp = NewRegexp(pieces[3])
}

func (sc *specialCharItem) Show() {
	fmt.Printf("%s %s %d\n",string(sc.specialChar),sc._type,sc.kindCodeFlag)
	sc.regexp.Show()
}
