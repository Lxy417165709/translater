package conf


type SyntaxConf struct {
	SyntaxFilePath string `json:"SyntaxFilePath"`
	EndSymbol string `json:"EndSymbol"`
	BlankSymbol string `json:"BlankSymbol"`
	StartSymbol string `json:"StartSymbol"`
	AdditionCharBeginChar string `json:"AdditionCharBeginChar"`
	DelimiterOfSyntaxPieces string `json:"DelimiterOfSyntaxPieces"`
	DelimiterOfSentences string `json:"DelimiterOfSentences"`
	DelimiterOfSymbols string `json:"DelimiterOfSymbols"`
}
