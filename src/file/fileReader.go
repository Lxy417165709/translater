package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

const (
	lineDelimiter = '\n'
)

type FileReader struct {
	filePath string
}

func NewFileReader(filePath string) *FileReader {
	return &FileReader{filePath}
}

func (fr *FileReader) GetFileBytes() []byte {
	file, err := os.Open(fr.filePath)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (fr *FileReader) GetFileLines() []string {
	lines := make([]string, 0)
	file, err := os.Open(fr.filePath)
	if err != nil {
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
