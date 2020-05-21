package LLONE

import "fmt"

func (one *LLOne) GetFirst() {
	one.First = make(map[string][]string)
	one.delimitersBuffer = make(map[string][]string)
	for  {
		for one.initHandleUnitPosition(); one.handlingUnitIsNotOver(); one.moveToNextUnit() {
			for one.initHandleUnitExpressPosition(); one.handlingUnitExpressIsNotOver(); one.moveToNextUnitExpress() {
				one.handleGettingFirst()
			}
		}
		if !one.syncFirstDelimiterBuffer(){
			break
		}
	}
	for k,v := range one.First{
		fmt.Println(k,v)
	}
}
func (one *LLOne) handleGettingFirst() {
	handlingUnit := one.handledUnits[one.handlingUnitPosition]
	handlingExpress := handlingUnit.expresses[one.handlingUnitExpressPosition]
	if expressIsBlank(handlingExpress) {
		one.handleGettingFirstOfExpressIsBlank()
	} else {
		one.handleGettingFirstOfExpressIsNotBlank()
	}
}

func (one *LLOne) handleGettingFirstOfExpressIsBlank() {
	handlingUnit := one.handledUnits[one.handlingUnitPosition]
	one.delimitersBuffer[handlingUnit.symbol] = append(one.delimitersBuffer[handlingUnit.symbol], blankDelimiter)
}
func (one *LLOne) handleGettingFirstOfExpressIsNotBlank() {
	handlingUnit := one.handledUnits[one.handlingUnitPosition]
	handlingExpress := handlingUnit.expresses[one.handlingUnitExpressPosition]
	for pIndex, expressPart := range handlingExpress {
		firstSetOfExpressPart := one.First[expressPart]
		switch {
		case isEndDelimiters(expressPart):

			one.delimitersBuffer[handlingUnit.symbol] = append(one.delimitersBuffer[handlingUnit.symbol], expressPart)
			return
		case !hasBlankDelimiter(firstSetOfExpressPart):
			one.delimitersBuffer[handlingUnit.symbol] = append(one.delimitersBuffer[handlingUnit.symbol], delBlankDelimiter(firstSetOfExpressPart)...)
			return
		case hasBlankDelimiter(firstSetOfExpressPart):
			if pIndex == len(handlingExpress)-1 {
				one.delimitersBuffer[handlingUnit.symbol] = append(one.delimitersBuffer[handlingUnit.symbol], firstSetOfExpressPart...)
			} else {
				one.delimitersBuffer[handlingUnit.symbol] = append(one.delimitersBuffer[handlingUnit.symbol], delBlankDelimiter(firstSetOfExpressPart)...)
			}
		default:
			panic("存在没有考虑的情况")
		}
	}
}
func (one *LLOne) syncFirstDelimiterBuffer() bool {
	isAdd := false
	for symbol,expressParts:= range one.delimitersBuffer{
		for _,expressPart := range expressParts{
			if !one.endDelimiterIsLivingInFirst(symbol, expressPart) {
				one.First[symbol] = append(one.First[symbol], expressPart)
				isAdd = true
			}
		}
	}
	one.flushDelimitersBuffer()
	return isAdd
}





func (one *LLOne) handlingUnitIsNotOver() bool {
	return one.handlingUnitPosition < len(one.handledUnits)
}
func (one *LLOne) handlingUnitExpressIsNotOver() bool {
	handlingUnit := one.handledUnits[one.handlingUnitPosition]
	return one.handlingUnitExpressPosition < len(handlingUnit.expresses)
}
func (one *LLOne) initHandleUnitPosition() {
	one.handlingUnitPosition = 0
}
func (one *LLOne) initHandleUnitExpressPosition() {
	one.handlingUnitExpressPosition = 0
}
func (one *LLOne) moveToNextUnit() {
	one.handlingUnitPosition++
}
func (one *LLOne) moveToNextUnitExpress() {
	one.handlingUnitExpressPosition++
}
func (one *LLOne) endDelimiterIsLivingInFirst(symbol string, endDelimiter string) bool {
	for _, delimiter := range one.First[symbol] {
		if endDelimiter == delimiter {
			return true
		}
	}
	return false
}

func expressIsBlank(express []string) bool {
	return len(express) == 1 && express[0] == blankDelimiter
}
