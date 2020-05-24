package grammar

func (sct *SpecialCharTable) CharIsSpecial(char byte) bool {
	return sct.specialCharToRegexp[char] != nil
}
func (sct *SpecialCharTable) GetRegexp(specialChar byte) *Regexp {
	return sct.specialCharToRegexp[specialChar]
}
func (sct *SpecialCharTable) GetType(specialChar byte) string {
	return sct.specialCharToType[specialChar]
}
func (sct *SpecialCharTable) GetCode(specialChar byte, word string) int {
	if sct.wordIsFixed(word) {
		return sct.fixedWordToCode[word]
	}
	if sct.CharIsVariable(specialChar) {
		return sct.variableCharToCode[specialChar]
	}
	panic("获取 code 出现错误")
}
func (sct *SpecialCharTable) CharIsVariable(specialChar byte) bool {
	return sct.variableCharToCode[specialChar] != 0
}
func (sct *SpecialCharTable) wordIsFixed(word string) bool {
	return sct.fixedWordToCode[word] != 0
}
