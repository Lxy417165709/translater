package conf

type GrammarConf struct {
	SpecialCharTableFilePath                     string `json:"SpecialCharTableFilePath"`
	DelimiterOfPieces	string `json:"DelimiterOfPieces"`
	DelimiterOfWords	string `json:"DelimiterOfWords"`
	MatchMoreThanOnceSymbol      string `json:"MatchMoreThanOnceSymbol"`
	MatchMoreThanZeroTimesSymbol string `json:"MatchMoreThanZeroTimesSymbol"`
}
