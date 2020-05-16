package grammar

import (
	"file"
	"fmt"
	"stateMachine"
)

func BuildGrammar() {
	lines := file.NewFileReader(grammarFilePath).GetFileLines()
	for index, line := range lines {
		unit := NewGrammarUnit(0,"")
		unit.Parse(line)
		stateMachine.GlobalRegexpsManager.AddSpecialChar(unit.SpecialChar, unit.Regexp)
		fmt.Printf("添加了第 %d 个特殊字符：%s   对应的正则表达式是：%s\n",index,string(unit.SpecialChar),unit.Regexp)
	}
}

