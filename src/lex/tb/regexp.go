package tb

import (
	"fmt"
	"strings"
)

type Regexp struct {
	Words []string
}

func NewRegexp(content string, delimiterOfWords string) (*Regexp, error) {
	regexp := &Regexp{}
	content = strings.TrimSpace(content)
	for index, unTrimWord := range strings.Split(content, delimiterOfWords) {
		word := strings.TrimSpace(unTrimWord)
		if len(word) == 0 {
			return nil, fmt.Errorf("%s 第 %d 个单词(%s)为空", content, index, unTrimWord)
		}
		regexp.Words = append(regexp.Words, word)
	}
	return regexp, nil
}
