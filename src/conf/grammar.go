package conf

type GrammarConf struct {
	FilePath      string `json:"FilePath"`
	WordDelimiter string `json:"WordDelimiter"`
	UnitDelimiter string `json:"UnitDelimiter"`
	SpecialCharOfFixedWord string `json:"SpecialCharOfFixedWord"`
	SpecialCharOfVariableWord string `json:"SpecialCharOfVariableWord"`
}
