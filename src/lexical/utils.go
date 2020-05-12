package lexical

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}
func bytesToNumber(bytes []byte) int {
	result := 0
	for i := 0; i < len(bytes); i++ {
		result = result*10 + int(bytes[i]-'0')
	}
	return result
}
