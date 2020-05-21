package LLONE

import (
	"fmt"
	"strings"
)
// END 是结束符
var terminators = [...]string{
	"LEFT_PAR", "RIGHT_PAR", "IDE", "FDO", "ASO","END","ZS",
}

const blankSymbol = "BLA"
const additionCharBeginChar = byte('a')


func (u *production) Parse(line string) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, "->")
	if len(parts) != 2 {
		panic("分割production发生错误，分割后的长度不为2")
	}
	u.leftNonTerminator = strings.TrimSpace(parts[0])
	for _, sentenceString := range strings.Split(parts[1], "|") {
		sentence := NewEmptySentence()
		sentence.Parse(sentenceString)
		u.sentences = append(u.sentences,sentence)
	}
}
func (u *production) nthSentenceFirstSymbolIsTerminator(index int) bool {
	for i := 0; i < len(terminators); i++ {
		if u.getNthSentenceFirstSymbol(index) == terminators[i] {
			return true
		}
	}
	return false
}
func (u *production) nthSentenceFirstSymbolIsBlank(index int) bool {
	return len(u.getNthSentenceFirstSymbol(index)) == 1 && u.getNthSentenceFirstSymbol(index) == blankSymbol
}
func (u *production) getNthSentenceFirstSymbol(index int) string {
	return u.sentences[index].symbols[0]
}
func (u *production) getLeftRecursionSentence() []*sentence {
	result := make([]*sentence, 0)
	for _, sentence := range u.sentences {
		if u.leftNonTerminator == sentence.symbols[0] {
			result = append(result, sentence)
		}
	}
	return result
}
func (u *production) getFirstNotLeftRecursionSentence() *sentence {
	for _, sentence := range u.sentences {
		if u.leftNonTerminator != sentence.symbols[0] {
			return sentence
		}
	}
	return nil
}
func (u *production) hasLeftRecursionSentence() bool {
	for index := range u.sentences {
		if u.nthSentenceIsLeftRecursion(index) {
			return true
		}
	}
	return false
}
func (u *production) nthSentenceIsLeftRecursion(sentenceIndex int) bool {
	return u.sentences[sentenceIndex].symbols[0] == u.leftNonTerminator
}
func (u *production) ChangeToNonLeftRecursionProductions() []*production {
	result := make([]*production, 0)
	if !u.hasLeftRecursionSentence() {
		result = append(result, u)
		return result
	}

	leftRecursionSentences := u.getLeftRecursionSentence()
	additionChar := additionCharBeginChar
	for _, sentence := range leftRecursionSentences {
		result = append(result, u.formNonLeftRecursionProductions(sentence, additionChar)...)
		additionChar++
	}
	return result
}
func (u *production) formLeftNonTerminator(additionChar byte) string {
	return fmt.Sprintf("%s%s", u.leftNonTerminator, string(additionChar))
}
func (u *production) formNonLeftRecursionProductions(stc *sentence, additionChar byte) []*production {
	result := make([]*production, 0)
	newLeftNonTerminator := u.formLeftNonTerminator(additionChar)
	notLeftRecursionSentence := u.getFirstNotLeftRecursionSentence()

	result = append(result, &production{
		newLeftNonTerminator,
		[]*sentence{
			NewSentence(append(stc.symbols[1:], newLeftNonTerminator)),
			NewBlankSentence(),
		},
	})
	result = append(result, &production{
		u.leftNonTerminator,
		[]*sentence{
			NewSentence(append(notLeftRecursionSentence.symbols,newLeftNonTerminator))	,
		},
	})
	return result
}

func removeBlankSymbol(sentence []string) []string {
	result := make([]string, 0)
	for _, symbol := range sentence {
		if symbol == blankSymbol {
			continue
		}
		result = append(result, symbol)
	}
	return result
}
func hasBlankSymbol(sentence []string) bool {
	for _, symbol := range sentence {
		if symbol == blankSymbol {
			return true
		}
	}
	return false
}
func isTerminator(symbol string) bool {
	for _,terminator := range terminators{
		if symbol == terminator{
			return true
		}
	}
	return false
}

type production struct {
	leftNonTerminator string
	sentences []*sentence
}










