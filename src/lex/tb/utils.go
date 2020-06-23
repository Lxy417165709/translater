package tb

import "strings"

func TrimAllPiecesSpace(pieces []string){
	for i:=0;i<len(pieces);i++{
		pieces[i] = strings.TrimSpace(pieces[i])
	}
}
