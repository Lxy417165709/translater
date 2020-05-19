package conf

type GrammarConf struct {
	FilePath      string `json:"FilePath"`
	PartDelimiter string `json:"PartDelimiter"`
	RegexpDelimiter string `json:"RegexpDelimiter"`
}
