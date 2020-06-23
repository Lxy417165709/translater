package tb

import (
	"fmt"
	"strings"
)

type SymbolTableItem struct {
	SpecialChar byte
	TerminatorReference string
}

func NewSymbolTableItem(content string,delimiterOfPieces string) (*SymbolTableItem,error){
	item := &SymbolTableItem{}
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
	item.TerminatorReference = pieces[1]
	return item, nil
}


