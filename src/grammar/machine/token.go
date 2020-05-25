package machine

import (
	"fmt"
	"strconv"
)

type Token struct {
	specialChar byte
	kindCode int
	_type string
	value string
}

func (t *Token)toLine(lineIndex int) string{
	return fmt.Sprintf(
		"%d|`%s`|`%s`|`%d`\n",
		lineIndex+1,
		t.value,
		t._type,
		t.kindCode,
	)
}
func (t *Token) GetRealValue() interface{} {
	switch t._type {
	case "整数":
		realValue,err :=strconv.Atoi(t.value)
		if err!=nil{
			panic(err)
		}
		return realValue
	}
	return t.value
}
func (t *Token) GetKindCode() int{
	return t.kindCode
}
