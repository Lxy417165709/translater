package grammar

import "fmt"

func AddBackticks(source string) string{
	return fmt.Sprintf("`%s`",source)
}
