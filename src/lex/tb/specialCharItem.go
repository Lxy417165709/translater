package tb

import (
	"fmt"
	"strings"
)

type SpecialCharItem struct {
	SpecialChar byte
	Regexp      *Regexp
}

func NewSpecialCharItem(content string, delimiterOfPieces string, delimiterOfWords string) (*SpecialCharItem, error) {
	item := &SpecialCharItem{}
	content = strings.TrimSpace(content)
	pieces := strings.Split(content, delimiterOfPieces)
	if len(pieces) != 3 {
		return nil, fmt.Errorf("%s 格式不正确，分割符为 %s", content, delimiterOfPieces)
	}

	TrimAllPiecesSpace(pieces)

	if len(pieces[0]) != 1 {
		return nil, fmt.Errorf("%s 特殊符号(%s)格式不正确", content, pieces[0])
	}
	item.SpecialChar = pieces[0][0]
	regexp, err := NewRegexp(pieces[1], delimiterOfWords)
	if err != nil {
		return nil, err
	}
	item.Regexp = regexp

	return item, nil
}
