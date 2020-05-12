package lexical



type LexicalAnalyzer struct {
	stateMachine *stateMachine
}

func NewLexicalAnalyzer(blankChars, delimiters []byte,reservedWords, operators []string) *LexicalAnalyzer{
	stateMachine := NewStateMachine(blankChars,delimiters,reservedWords,operators)
	return &LexicalAnalyzer{
		stateMachine,
	}
}


func (la *LexicalAnalyzer) ParseSourceFileBytes(sourceFileBytes []byte) []*pair {
	la.stateMachine.SetHandleBytes(sourceFileBytes)
	return la.stateMachine.Handle()
}


