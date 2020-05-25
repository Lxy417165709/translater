package syntex

import (
	"conf"
	"fmt"
	"strings"
)

type production struct {
	leftNonTerminator string
	sentences []*sentence
}

func NewProduction(leftNonTerminator string,sentences []*sentence)*production {
	return &production{
		leftNonTerminator:leftNonTerminator,
		sentences:sentences,
	}
}


func (u *production) Parse(line string) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, "->")
	if len(parts) != 2 {
		panic("分割production发生错误，分割后的长度不为2")
	}
	u.leftNonTerminator = strings.TrimSpace(parts[0])
	for _, sentenceString := range strings.Split(parts[1], conf.GetConf().SyntaxConf.DelimiterOfSentences) {
		sentence := NewSentence(nil)
		sentence.Parse(sentenceString)
		u.sentences = append(u.sentences,sentence)
	}
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
	additionChar := conf.GetConf().SyntaxConf.AdditionCharBeginChar[0]
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
	production1 := NewProduction(newLeftNonTerminator,[]*sentence{
		NewSentence(append(stc.symbols[1:], newLeftNonTerminator)),
		NewSentence([]string{conf.GetConf().SyntaxConf.BlankSymbol}),
	})
	production2 := NewProduction(u.leftNonTerminator,[]*sentence{
		NewSentence(append(notLeftRecursionSentence.symbols,newLeftNonTerminator)),
	})
	result = append(result,production1,production2)
	return result
}













