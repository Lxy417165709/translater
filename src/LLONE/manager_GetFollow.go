package LLONE



//
// 感觉可以用模板方法模式
func (one *LLOne) GetFollow() {
	one.Follow = make(map[string][]string)
	one.delimitersBuffer = make(map[string][]string)
	one.Follow["EXP"] = append(one.Follow["EXP"], "END")	// 添加终止符

	for  {
		for one.initHandleUnitPosition(); one.handlingUnitIsNotOver(); one.moveToNextUnit() {
			for one.initHandleUnitExpressPosition(); one.handlingUnitExpressIsNotOver(); one.moveToNextUnitExpress() {
				one.handleGettingFollow()
			}
		}
		if !one.syncFollowDelimiterBuffer() {
			break
		}
	}

	// TODO: BLA [END ASO RIGHT_PAR] 空白的应该加入吗？
	//for k,v := range one.Follow{
	//	fmt.Println(k,v)
	//}
}

// 有冗余
func (one *LLOne) syncFollowDelimiterBuffer() bool {
	isAdd := false
	for symbol,expressParts:= range one.delimitersBuffer{
		for _,expressPart := range expressParts{
			if !one.endDelimiterIsLivingInFollow(symbol, expressPart) {
				one.Follow[symbol] = append(one.Follow[symbol], expressPart)
				isAdd = true
			}
		}
	}
	one.flushDelimitersBuffer()
	return isAdd
}
// 有冗余
func (one *LLOne) endDelimiterIsLivingInFollow(symbol string, endDelimiter string) bool {
	for _, delimiter := range one.Follow[symbol] {
		if endDelimiter == delimiter {
			return true
		}
	}
	return false
}


func (one *LLOne) handleGettingFollow() {
	handlingUnit := one.handledUnits[one.handlingUnitPosition]
	handlingExpress := handlingUnit.expresses[one.handlingUnitExpressPosition]

	for i:=0;i<len(handlingExpress);i++{
		if isEndDelimiters(handlingExpress[i]){
			continue
		}
		nextPosition := i+1
		switch {

		// 位于最后
		case i==len(handlingExpress)-1:
			one.delimitersBuffer[handlingExpress[i]] = append(one.delimitersBuffer[handlingExpress[i]],one.Follow[handlingUnit.symbol]...)

		case  i==len(handlingExpress)-2:
			// 位于倒数第二，而且倒数第一存在空符
			if hasBlankDelimiter(one.First[handlingExpress[nextPosition]]){
				one.delimitersBuffer[handlingExpress[i]] = append(one.delimitersBuffer[handlingExpress[i]],one.Follow[handlingUnit.symbol]...)
			}

			// 冗余
			if isEndDelimiters(handlingExpress[nextPosition]) {
				one.delimitersBuffer[handlingExpress[i]] = append(
					one.delimitersBuffer[handlingExpress[i]],
					handlingExpress[nextPosition],
				)
			}else{
				one.delimitersBuffer[handlingExpress[i]] = append(
					one.delimitersBuffer[handlingExpress[i]],
					delBlankDelimiter(one.First[handlingExpress[nextPosition]])...
				)
			}
		default:
			// 冗余
			if isEndDelimiters(handlingExpress[nextPosition]) {
				one.delimitersBuffer[handlingExpress[i]] = append(
					one.delimitersBuffer[handlingExpress[i]],
					handlingExpress[nextPosition],
				)
			}else{
				one.delimitersBuffer[handlingExpress[i]] = append(
					one.delimitersBuffer[handlingExpress[i]],
					delBlankDelimiter(one.First[handlingExpress[nextPosition]])...
				)
			}
		}
	}
}
