package conf

import (
	"encoding/json"
	"io/ioutil"
)

const configureFilePath = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\conf.json`

var singleConf *conf

func init() {
	singleConf = &conf{}
	singleConf.init()
}

type conf struct {
	GrammarConf    GrammarConf `json:"GrammarConf"`
	LexicalConf    LexicalConf `json:"LexicalConf"`
}

func GetConf() *conf {
	return singleConf
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
