package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const lineDelimiter = '\n'

func GetFileLines(filePath string) ([]string, error) {
	var err error
	var file *os.File
	lines := make([]string, 0)

	if file, err = os.Open(filePath); err != nil {
		return nil, fmt.Errorf("%s 打开失败。", filePath)
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for lineCount := 0; ; lineCount++ {
		var line string
		var err error

		if line, err = buf.ReadString(lineDelimiter); err != nil {
			if err ==io.EOF {
				break
			}
			return nil, fmt.Errorf("%s，第 %d 行读取时发生错误。", filePath, lineCount+1)
		}

		lines = append(lines, line)
	}
	return lines, nil
}
