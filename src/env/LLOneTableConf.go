package env


type LLOneTableConf struct{
	FilePath string `json:"FilePath"`
	DelimiterOfPieces string  `json:"DelimiterOfPieces"`
	DelimiterOfSentences string `json:"DelimiterOfSentences"`
	DelimiterOfSymbols string `json:"DelimiterOfSymbols"`
	BlankSymbol string `json:"BlankSymbol"`
	StartSymbol string `json:"StartSymbol"`
	EndSymbol string `json:"EndSymbol"`
}
