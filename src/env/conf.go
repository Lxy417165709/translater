package env

import (
	"encoding/json"
	"io/ioutil"
)

func Init(ConfFilePath string) {
	singleConf = &Conf{}
	singleConf.init(ConfFilePath)
}
func GetConf() *Conf {
	return singleConf
}
var singleConf *Conf
type Conf struct {
	CompilerConf *CompilerConf `json:"CompilerConf"`
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
