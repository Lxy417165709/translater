package conf

type LexicalConf struct {
	SpecialCharsOfNFAs string `json:"SpecialCharsOfNFAs"`
	InformationDir		string `json:"InformationDir"`
	SourceFilePath string      `json:"SourceFilePath"`
	StateMachineDirName string `json:"StateMachineDirName"`
	FileNameOfStoringKindCodes string `json:"FileNameOfStoringKindCodes"`
	FileNameOfStoringTokens string `json:"FileNameOfStoringTokens"`
	FileNameOfStoringFinalNFA string `json:"FileNameOfStoringFinalNFA"`
	DisplayDocumentPath string `json:"DisplayDocumentPath"`
}
