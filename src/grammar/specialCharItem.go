package grammar

import (
	"fmt"
	"strconv"
	"strings"
)

type specialCharItem struct {
	specialChar       byte
	_type             string
	kindCodeFlag      int
	regexp            *Regexp
	delimiterOfPieces string
	delimiterOfWords  string
}

func NewSpecialCharItem(specialCharItemLine, delimiterOfPieces, delimiterOfWords string) *specialCharItem {
	sc := &specialCharItem{
		delimiterOfPieces: delimiterOfPieces,
		delimiterOfWords:  delimiterOfWords,
	}
	sc.Parse(specialCharItemLine)
	return sc
}

func (sc *specialCharItem) Parse(content string) {
	content = strings.TrimSpace(content)

	pieces := strings.Split(strings.TrimSpace(content), sc.delimiterOfPieces)
	if len(pieces) != 4 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于4", pieces))
	}
	sc.specialChar = strings.TrimSpace(pieces[0])[0]
	sc._type = strings.TrimSpace(pieces[1])

	var err error
	if sc.kindCodeFlag, err = strconv.Atoi(pieces[2]); err != nil {
		panic(err)
	}

	sc.regexp = NewRegexp(pieces[3], sc.delimiterOfWords)
}

func (sc *specialCharItem) Show() {
	fmt.Print(string(sc.specialChar), " ", sc._type, " ", sc.kindCodeFlag, " ")
	sc.regexp.Show()
}


//func  (sc *specialCharItem) GetWords() []string {
//	words := strings.Split(sc.MatchRegexp, sc.RegexpDelimiter)
//	for i := 0; i < len(words); i++ {
//		words[i] = strings.TrimSpace(words[i])
//	}
//	return words
//}

//func  (sc *specialCharItem) reformToLine() string {
//	partOne := AddBackticks(string(sc.SpecialChar))
//	partTwo := AddBackticks(string(sc.Type))
//	partThree := AddBackticks(strconv.Itoa(sc.KindCodeRule))
//	partFour := bytes.Buffer{}
//
//	words := sc.GetWords()
//	for index, word := range words {
//		if index == len(words)-1 {
//			partFour.WriteString(AddBackticks(word))
//		} else {
//			partFour.WriteString(AddBackticks(word) + sc.RegexpDelimiter)
//		}
//	}
//	return fmt.Sprintf(
//		"%s%s%s%s%s%s%s",
//		partOne,
//		sc.PartDelimiter,
//		partTwo,
//		sc.PartDelimiter,
//		partThree,
//		sc.PartDelimiter,
//		partFour.String(),
//	)
//}
