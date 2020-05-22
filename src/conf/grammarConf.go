package conf

type GrammarConf struct {
	SpecialCharTableFilePath                     string `json:"SpecialCharTableFilePath"`
	DelimiterOfPieces	string `json:"DelimiterOfPieces"`
	DelimiterOfWords	string `json:"DelimiterOfWords"`
	MatchMoreThanOnceSymbol      string `json:"MatchMoreThanOnceSymbol"`
	MatchMoreThanZeroTimesSymbol string `json:"MatchMoreThanZeroTimesSymbol"`

	RegexpDelimiter string	 // 准备删除的
	PartDelimiter string	 // 准备删除的
	FilePath string	 // 准备删除的
}
