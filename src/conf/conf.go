package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\conf.json`
var globalConf *conf

func init() {
	globalConf = &conf{}
	globalConf.init()
	fmt.Println(globalConf)
}
type conf struct {
	GrammarConf GrammarConf `json:"GrammarConf"`
	SourceFilePath  string `json:"SourceFilePath"`
	LexicalAnalyzerDisplayDocumentPath string `json:"LexicalAnalyzerDisplayDocumentPath"`
	LexicalInformationDir string `json:"LexicalInformationDir"`
}

func GetConf() *conf {
	return globalConf
}

func (cn *conf) init() {
	var err error
	var jsonBytes []byte
	if jsonBytes, err = ioutil.ReadFile(configureFilePath); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(jsonBytes, cn); err != nil {
		panic(err)
	}
}
