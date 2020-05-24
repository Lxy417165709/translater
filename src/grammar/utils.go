package grammar

import "fmt"
const (
	beginCode = 1
	coding = 0
	notCoding=-1
	codeStage = 100
)
func AddBackticks(source string) string{
	return fmt.Sprintf("`%s`",source)
}
