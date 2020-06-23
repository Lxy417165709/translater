package main

import (
	"bufio"
	"env"
	"final"
	"fmt"
	"os"
)

const confPath = `C:\Users\hasee\Desktop\Go_Practice\compiler\doc\conf.json`

func main() {
	env.Init(confPath)
	compiler, err := final.NewCompiler(env.GetConf().CompilerConf)
	if err != nil {
		panic(err)
	}

	var code string
	for {
		fmt.Print("> ")
		var codePart string
		Scanf(&codePart)
		if len(codePart) != 0 {
			code = code + codePart
			continue
		}

		if err := compiler.ExecCode(code); err != nil {
			fmt.Printf("(error)%s\n", err.Error())
		}
		code = ""
	}
}

func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
