package LLONE

import (
	"fmt"
	"strings"
)

type llUnit struct{
	symbol string
	expresses [][]string
}
func (u *llUnit)getLeftRecursionExpress() [][]string{
	result := make([][]string,0)
	for _,express := range u.expresses {
		if u.symbol==express[0]{
			result = append(result,express)
		}
	}
	return result
}
func (u *llUnit)getFirstNotLeftRecursionExpress() []string{
	for _,express := range u.expresses {
		if u.symbol!=express[0]{
			return express
		}
	}
	return nil
}


func (u *llUnit)Parse(line string) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line,"->")
	if len(parts)!=2{
		panic("分割llUnit发生错误，分割后的长度不为2")
	}
	u.symbol = strings.TrimSpace(parts[0])
	for index,element := range strings.Split(parts[1],"|"){
		u.expresses = append(u.expresses,[]string{})
		expressParts := strings.Split(element," ")
		for _,expressPart := range expressParts{
			expressPart = strings.TrimSpace(expressPart)

			u.expresses[index] = append(u.expresses[index],expressPart)
		}

	}
}

func (u *llUnit)nExpressFirstIsEndDelimiter(index int) bool{
	for i:=0;i<len(endDelimiters);i++{
		if u.getNExpressFirst(index)==endDelimiters[i]{
			return true
		}
	}
	return false
}
func (u *llUnit)nExpressFirstIsBlank(index int) bool{
	return len(u.getNExpressFirst(index))==1 && u.getNExpressFirst(index)==blankDelimiter
}



func (u *llUnit)getNExpressFirst(index int)string{
	return u.expresses[index][0]
}


func (u *llUnit) isLeftRecursion() bool{
	for _,express := range u.expresses{
		if u.expressIsLeftRecursion(express){
			return true
		}
	}
	return false
}

const additionCharBeginChar = byte( 'a')
func (u *llUnit) GetDelimiterLeftRecursionUnits() []*llUnit{
	result := make([]*llUnit,0)
	if !u.isLeftRecursion() {
		result = append(result,u)
		return result
	}

	leftRecursionExpresses := u.getLeftRecursionExpress()
	additionChar := additionCharBeginChar
	for _,express := range leftRecursionExpresses{
		result = append(result,u.FormNotLeftRecursionUnits(express,additionChar)...)
		additionChar++
	}
	return result
}

func (u *llUnit) FormNotLeftRecursionUnits(express []string,additionChar byte) []*llUnit{
	result := make([]*llUnit,0)
	newExpressSymbol := u.getExpressNewSymbol(additionChar)
	notLeftRecursionExpress := u.getFirstNotLeftRecursionExpress()
	result = append(result, &llUnit{
		newExpressSymbol,
		[][]string{append(express[1:],newExpressSymbol),{blankDelimiter}},
	})
	result = append(result, &llUnit{
		u.symbol,
		[][]string{append(notLeftRecursionExpress,newExpressSymbol)},
	})
	return  result
}
func (u *llUnit) getExpressNewSymbol(additionChar byte) string{
	return fmt.Sprintf("%s%s",u.symbol,string(additionChar))
}




var endDelimiters = []string{
	"LEFT_PAR","RIGHT_PAR","IDE","FDO","ASO","ASS","DEL",
}
var blankDelimiter = "BLA"
func isEndDelimiters(str string) bool{
	for i:=0;i<len(endDelimiters);i++{
		if str==endDelimiters[i]{
			return true
		}
	}
	return false
}

func hasBlankDelimiter(parts []string) bool{
	for _,part := range parts{
		if part==blankDelimiter{
			return true
		}
	}
	return false
}
func delBlankDelimiter(parts []string)[]string{
	result := make([]string,0)
	for _,part := range parts{
		if part==blankDelimiter{
			continue
		}
		result = append(result,part)
	}
	return result
}

func (u *llUnit)expressIsLeftRecursion(express []string) bool{
	return express[0]==u.symbol
}
