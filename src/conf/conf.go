package conf

import (
	"encoding/json"
	"io/ioutil"
)


var singleConf *Conf

func Init(ConfFilePath string) {
	singleConf = &Conf{}
	singleConf.init(ConfFilePath)
}

type Conf struct {
	IsMatchOfNFATestConf IsMatchOfNFATestTableConf `json:"IsMatchOfNFATestConf"`
	GrammarConf    GrammarConf `json:"GrammarConf"`
	LexicalConf    LexicalConf `json:"LexicalConf"`
	SyntaxConf SyntaxConf `json:"SyntaxConf"`
}

func GetConf() *Conf {
	return singleConf
}

func (cn *Conf) init(ConfFilePath string) {
	var err error
	var jsonBytes []byte
	if jsonBytes, err = ioutil.ReadFile(ConfFilePath); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(jsonBytes, cn); err != nil {
		panic(err)
	}
}
