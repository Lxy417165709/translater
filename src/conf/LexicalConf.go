package conf

type LexicalConf struct {
	SpecialCharsOfNFAs string `json:"SpecialCharsOfNFAs"`
	InformationDir		string `json:"InformationDir"`
	SourceFilePath string      `json:"SourceFilePath"`
	KindCodeToTerminatorFilePath string  `json:"KindCodeToTerminatorFilePath"`
	DelimiterOfKindCodeToTerminator string  `json:"DelimiterOfKindCodeToTerminator"`
}

