package conf

type GrammarConf struct {
	FilePath                     string `json:"FilePath"`
	PartDelimiter                string `json:"PartDelimiter"`
	RegexpDelimiter              string `json:"RegexpDelimiter"`
	MatchMoreThanOnceSymbol      string `json:"MatchMoreThanOnceSymbol"`
	MatchMoreThanZeroTimesSymbol string `json:"MatchMoreThanZeroTimesSymbol"`
}
