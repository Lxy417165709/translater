package tb

import (
	"fmt"
	"strings"
)

type Sentence struct{
	Symbols []string
}

func NewSentence(content string,delimiterOfSymbols string) (*Sentence,error) {
	sentence := &Sentence{}
	content = strings.TrimSpace(content)
	if len(content)==0{
		return nil,fmt.Errorf("语句解析内容为空")
	}
	symbols := strings.Split(content,delimiterOfSymbols)
	for _,symbol := range symbols{
		symbol = strings.TrimSpace(symbol)
		sentence.Symbols = append(sentence.Symbols,strings.TrimSpace(symbol))
	}
	return sentence,nil
}
