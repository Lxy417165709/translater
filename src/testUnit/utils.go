package testUnit

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	lineDelimiter = '\n'
	testUnitDelimiter = "||"
	grammarUnitDelimiter = "->"
)

func getTestUnits(filePath string) []*TestUnit {
	values := getUnits(filePath, parseTestUnitLine)
	units := make([]*TestUnit, 0)
	for _, value := range values {
		units = append(units, value.(*TestUnit))
	}
	return units
}
func getGrammarUnits(filePath string) []*GrammarUnit {
	values := getUnits(filePath, parseGrammarUnitLine)
	units := make([]*GrammarUnit, 0)
	for _, value := range values {
		units = append(units, value.(*GrammarUnit))
	}
	return units
}

func getUnits(filePath string, parseFunction interface{}) []interface{} {
	units := make([]interface{}, 0)
	var file *os.File
	var err error
	if file, err = os.Open(filePath); err != nil {
		panic(err)
	}
	lines := getFileLines(file)
	for _, line := range lines {
		funcValue := reflect.ValueOf(parseFunction)
		result := funcValue.Call([]reflect.Value{reflect.ValueOf(line)})[0].Interface()
		units = append(units, result)
	}
	return units
}

func parseGrammarUnitLine(line string) *GrammarUnit {
	parts := strings.Split(strings.TrimSpace(line), grammarUnitDelimiter)
	if len(parts) != 2 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于2", parts))
	}
	identity := strings.TrimSpace(parts[0])[0]
	regexp := strings.TrimSpace(parts[1])
	return NewGrammarUnit(identity, regexp)
}
func parseTestUnitLine(line string) *TestUnit {
	parts := strings.Split(strings.TrimSpace(line), testUnitDelimiter)
	if len(parts) != 3 {
		panic(fmt.Sprintf("分割测试单元：%v 失败，分割后的字段数不等于3", parts))
	}

	regex := strings.TrimSpace(parts[0])
	pattern := strings.TrimSpace(parts[1])
	matchFlag, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		panic(fmt.Sprintf("%v %v", err, parts))
	}
	return NewTestUnit(regex, pattern, intToBool(matchFlag))
}

func getFileLines(file *os.File) []string {
	lines := make([]string, 0)
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

func intToBool(a int) bool {
	return a != 0
}
