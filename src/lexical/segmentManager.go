package lexical

type charType int

const (
	blank charType = iota
	digit
	letter
	delimiter
	operatorChar
	invalid
)


type segmentManager struct{
	blankChars    []byte
	delimiters    []byte
	reservedWords []string
	operators     []string
	segmentToCode map[interface{}]int
	operatorCharJudgeHelper map[byte]bool
}
const invalidCode = -1

func NewSegmentManager(blankChars, delimiters []byte, reservedWords, operators []string)*segmentManager{
	segmentToCode := make(map[interface{}]int)
	code := 0
	for _, element := range reservedWords {
		segmentToCode[element] = code
		code++
	}
	for _, element := range delimiters {
		segmentToCode[element] = code
		code++
	}
	for _, element := range operators {
		segmentToCode[element] = code
		code++
	}

	operatorCharJudgeHelper := make(map[byte]bool)
	for _, operator := range operators {
		for _, char := range operator {
			operatorCharJudgeHelper[byte(char)] = true
		}
	}




	return &segmentManager{
		blankChars,
		delimiters,
		reservedWords,
		operators,
		segmentToCode,
		operatorCharJudgeHelper,
	}
}

func (sm *segmentManager) getCode(bytes []byte) int{
	switch {
	case sm.bytesIsReservedWord(bytes),sm.bytesIsOperator(bytes):
		return sm.segmentToCode[string(bytes)]
	case sm.bytesIsDelimiter(bytes):
		return sm.segmentToCode[bytes[0]]
	case sm.bytesIsNumber(bytes):
		return sm.getConstSpecialCode()
	case sm.bytesIsIdentify(bytes):
		return sm.getIdentifySpecialCode()
	}

	return invalidCode
}
func (sm *segmentManager) getIdentifySpecialCode() int {
	return len(sm.segmentToCode)
}
func (sm *segmentManager) getConstSpecialCode() int {
	return len(sm.segmentToCode) + 1
}

func (sm *segmentManager) bytesIsReservedWord(bytes []byte) bool{
	for _,reservedWord := range sm.reservedWords {
		if string(bytes) == reservedWord{
			return true
		}
	}
	return false
}
func (sm *segmentManager) bytesIsDelimiter(bytes []byte) bool{
	return len(bytes)==1 && sm.isDelimiter(bytes[0])
}
func (sm *segmentManager) bytesIsOperator(bytes []byte) bool{
	for _, operator := range sm.operators {
		if string(bytes) == operator {
			return true
		}
	}
	return false
}
func (sm *segmentManager) bytesIsNumber(bytes []byte)bool{
	for _,bt := range bytes{
		if isDigit(bt)==false{
			return false
		}
	}
	return true
}
func (sm *segmentManager) bytesIsIdentify(bytes []byte) bool{
	// 这应该涉及到标识符的状态机
	// 在这简化了，直接判断不是其它的 就是标识符
	if sm.bytesIsReservedWord(bytes){
		return false
	}
	if sm.bytesIsDelimiter(bytes){
		return false
	}
	if sm.bytesIsNumber(bytes) {
		return false
	}
	if sm.bytesIsOperator(bytes) {
		return false
	}

	return true
}


func (sm *segmentManager) getCharType(ch byte) charType {
	switch {
	case sm.isBlank(ch):
		return blank
	case isDigit(ch):
		return digit
	case isLetter(ch):
		return letter
	case sm.isDelimiter(ch):
		return delimiter
	case sm.isOperatorPart(ch):
		return operatorChar
	}
	return invalid
}
func (sm *segmentManager) isBlank(ch byte) bool {
	for _, blankChar := range sm.blankChars {
		if ch == blankChar {
			return true
		}
	}
	return false
}
func (sm *segmentManager) isDelimiter(ch byte) bool {
	for _, delimiter := range sm.delimiters {
		if ch == delimiter {
			return true
		}
	}
	return false
}
func (sm *segmentManager) isOperatorPart(ch byte) bool {
	return sm.operatorCharJudgeHelper[ch]
}
