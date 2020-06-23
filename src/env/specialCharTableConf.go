package env



// 词法分析文法配置（可以进行改名）
type SpecialCharTableConf struct {
	FilePath                     string `json:"FilePath"`
	DelimiterOfPieces	string `json:"DelimiterOfPieces"`
	DelimiterOfWords	string `json:"DelimiterOfWords"`
	SpecialCharOfMarchMoreThanOnce      string `json:"SpecialCharOfMarchMoreThanOnce"`
	SpecialCharOfMarchMoreThanZeroTimes string `json:"SpecialCharOfMarchMoreThanZeroTimes"`
}
