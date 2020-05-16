package file

import (
	"bufio"
	"io"
	"os"
)

const (
	lineDelimiter = '\n'
//	testUnitDelimiter = "||"
//	grammarUnitDelimiter = "->"
)


type fileReader struct{
	filePath string
}
func NewFileReader(filePath string) *fileReader{
	return &fileReader{filePath}
}



//func (fr *fileReader)GetUnits(object Parsable) []interface{} {
//	units := make([]interface{}, 0)
//	lines := fr.getFileLines()
//	for _, line := range lines {
//		object.Parse(line)
//		units = append(units, object)
//	}
//	return units
//}


func (fr *fileReader)GetFileLines() []string {
	lines := make([]string, 0)
	file,err := os.Open(fr.filePath)
	if err!=nil{
		panic(err)
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString(lineDelimiter)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		lines = append(lines, line)
	}
	return lines
}

