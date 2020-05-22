package grammar

import (
	"fmt"
	"strings"
)

type Regexp struct{
	words []string
	delimiterOfWords string
}
func NewRegexp(RegexpLine string,delimiterOfWords string) *Regexp{
	r := &Regexp{delimiterOfWords:delimiterOfWords}
	r.Parse(RegexpLine)
	return r
}
func(r *Regexp) Parse(content string) {
	content = strings.TrimSpace(content)
	words := strings.Split(content,r.delimiterOfWords)
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
