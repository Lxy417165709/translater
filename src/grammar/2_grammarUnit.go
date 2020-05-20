package grammar

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	notCoding = -1
	coding    = 0
)

type GrammarUnit struct {
	SpecialChar     byte
	Type            string
	KindCodeRule    int
	MatchRegexp     string
	PartDelimiter   string
	RegexpDelimiter string
}

func (unit *GrammarUnit) Parse(line string) {
	var err error
	parts := strings.Split(strings.TrimSpace(line), unit.PartDelimiter)
	if len(parts) != 4 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于4", parts))
	}
	unit.SpecialChar = strings.TrimSpace(parts[0])[0]
	unit.Type = strings.TrimSpace(parts[1])
	unit.KindCodeRule, err = strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	unit.MatchRegexp = strings.TrimSpace(parts[3])
}

func NewGrammarUnit(PartDelimiter string, RegexpDelimiter string) *GrammarUnit {
	return &GrammarUnit{PartDelimiter: PartDelimiter, RegexpDelimiter: RegexpDelimiter}
}

func (unit *GrammarUnit) reformToLine() string {
	partOne := AddBackticks(string(unit.SpecialChar))
	partTwo := AddBackticks(string(unit.Type))
	partThree := AddBackticks(strconv.Itoa(unit.KindCodeRule))
	partFour := bytes.Buffer{}

	words := unit.GetWords()
	for index, word := range words {
		if index == len(words)-1 {
			partFour.WriteString(AddBackticks(word))
		} else {
			partFour.WriteString(AddBackticks(word) + unit.RegexpDelimiter)
		}
	}
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		partOne,
		unit.PartDelimiter,
		partTwo,
		unit.PartDelimiter,
		partThree,
		unit.PartDelimiter,
		partFour.String(),
	)
}
func (unit *GrammarUnit) GetWords() []string {
	words := strings.Split(unit.MatchRegexp, unit.RegexpDelimiter)
	for i := 0; i < len(words); i++ {
		words[i] = strings.TrimSpace(words[i])
	}
	return words
}
