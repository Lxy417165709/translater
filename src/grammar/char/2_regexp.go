package char

import (
	"conf"
	"fmt"
	"strings"
)

type Regexp struct{
	words []string
}
func NewRegexp(RegexpLine string) *Regexp{
	r := &Regexp{}
	r.Parse(RegexpLine)
	return r
}
func(r *Regexp) Parse(content string) {
	content = strings.TrimSpace(content)
	words := strings.Split(content,conf.GetConf().GrammarConf.DelimiterOfWords)
	for _,word := range words{
		word = strings.TrimSpace(word)
		r.words =append(r.words,word)
	}
}
func (r *Regexp)Show() {
	fmt.Println(r.words)
}
func(r *Regexp)GetWord(index int) string{
	return r.words[index]
}
func(r *Regexp)GetWords() []string{
	return r.words
}
