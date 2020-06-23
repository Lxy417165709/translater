package tb

import (
	"fmt"
	"strings"
)

type Production struct {
	nonTerminator string
	sentences     []*Sentence
}

// 分隔符也可以进行配置
// 错误判断可以独立出来的
func NewProduction(content string, delimiterOfPieces, delimiterOfSentences, delimiterOfSymbols string) (*Production, error) {
	production := &Production{}
	content = strings.TrimSpace(content)

	parts := strings.Split(content, delimiterOfPieces)
	if len(parts) != 2 {
		return nil, fmt.Errorf("产生式(%s)分割失败", content)
	}
	production.nonTerminator = strings.TrimSpace(parts[0])
	if len(production.nonTerminator) == 0 {
		return nil, fmt.Errorf("产生式(%s) 的左部非终结符为空", content)
	}

	for index, sentenceString := range strings.Split(parts[1], delimiterOfSentences) {
		sentence, err := NewSentence(sentenceString, delimiterOfSymbols)
		if err != nil {
			return nil, fmt.Errorf("产生式(%s) 第 %d 个语句(%s) 解析失败，%s", content, index, sentenceString, err.Error())
		}
		production.sentences = append(production.sentences, sentence)
	}
	return production, nil
}
