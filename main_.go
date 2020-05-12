package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"lexical"
)

var reservedWords = [...]string{
	"if", "then", "else", "while", "begin", "do", "end","int",
}
var delimiters = [...]byte{
	';', ',',')' ,'(','{','}',
}
var operators = [...]string{
	"+", "-", "*", "/", "<", ">", "=",
	"+=", "-=", "*=", "/=", "<=", ">=", "==", "<<=", ">>=",
	"&", "|", "!", "^",
	"&=", "|=", "!=", "^=",
	"&&", "||",
	"++", "--", "<<", ">>",
	":=","**",
}
var blankChars = [...]byte{
	'\t',' ','\n','\r',
}

var lexicalAnalyzer = lexical.NewLexicalAnalyzer(blankChars[:],delimiters[:],reservedWords[:],operators[:])

func main() {
	sourceFileBytes := getSourceFileBytes()
	segments := lexicalAnalyzer.ParseSourceFileBytes(sourceFileBytes)
	for index,segment := range segments{
		fmt.Printf("%d:  %v\n",index,segment)
	}
}



func getSourceFileBytes() []byte{
	var sourceFile *os.File
	var err error
	sourceFilePath := `C:\Users\hasee\Desktop\Go_Practice\编译器\source.md`
	if sourceFile, err = os.Open(sourceFilePath); err != nil {
		panic(err)
	}
	var sourceFileBytes []byte
	if sourceFileBytes, err = ioutil.ReadAll(sourceFile); err != nil {
		panic(err)
	}
	return sourceFileBytes
}


func getFileBytes(filePath string) []byte{
	var file  *os.File
	var err error
	if file, err = os.Open(filePath); err != nil {
		panic(err)
	}
	var bytes []byte
	if bytes, err = ioutil.ReadAll(file); err != nil {
		panic(err)
	}


	return bytes
}
