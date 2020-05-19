package conf

import "fmt"

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


func (lc *LexicalConf) GetStorePathOfKindCodes() string {
	return fmt.Sprintf("%s/%s", lc.InformationDir,lc.FileNameOfStoringKindCodes)
}
func (lc *LexicalConf) GetStorePathOfTokens() string {
	return fmt.Sprintf("%s/%s", lc.InformationDir, lc.FileNameOfStoringTokens)
}
func (lc *LexicalConf) GetStoreDirPathOfStateMachine() string {
	return fmt.Sprintf("%s/%s", lc.InformationDir, lc.StateMachineDirName)
}

