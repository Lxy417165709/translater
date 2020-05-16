package stateMachine

var GlobalRegexpsManager = NewRegexpsManager()
type RegexpsManager struct {
	charToRegexp map[byte]string
}
func NewRegexpsManager() *RegexpsManager {
	return &RegexpsManager{make(map[byte]string)}
}

func (nm *RegexpsManager)AddSpecialChar(specialChar byte, regexp string) {
	nm.charToRegexp[specialChar] = regexp
}
func (nm *RegexpsManager)GetRegexp(specialChar byte) string{
	return nm.charToRegexp[specialChar]
}
func (nm *RegexpsManager)CharIsSpecial(char byte) bool{
	return nm.charToRegexp[char]!=""
}


