package lexicalTest



func charToNFA(char byte) *NFA {
	if !GlobalRegexpsManager.CharIsSpecial(char) {
		return NewNFA(char)
	}

	regexp := GlobalRegexpsManager.GetRegexp(char)
	nfa := NewNFABuilder(regexp).BuildNFA()
	nfa.startState.markFlag = char
	return nfa // 这里创建的是DFA
}
