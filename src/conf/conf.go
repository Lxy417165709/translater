package conf

import (
	"encoding/json"
	"io/ioutil"
)


var singleConf *conf

func Init(confFilePath string) {
	singleConf = &conf{}
	singleConf.init(confFilePath)
}

type conf struct {
	IsMatchOfNFATestConf IsMatchOfNFATestTableConf `json:"IsMatchOfNFATestConf"`
	GrammarConf    GrammarConf `json:"GrammarConf"`
	LexicalConf    LexicalConf `json:"LexicalConf"`
}

func GetConf() *conf {
	return singleConf
}

func (cn *conf) init(confFilePath string) {
	var err error
	var jsonBytes []byte
	if jsonBytes, err = ioutil.ReadFile(confFilePath); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(jsonBytes, cn); err != nil {
		panic(err)
	}
}
