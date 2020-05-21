package LLONE

import (
	"file"
)

type LLOne struct {
	filePath string

	llUnits      []*llUnit
	handledUnits []*llUnit

	First  map[string][]string
	Follow map[string][]string
	Select map[string][]string

	handlingUnitPosition        int
	handlingUnitExpressPosition int
	delimitersBuffer            map[string][]string
}

func NewLLOne(filePath string) *LLOne {
	one := &LLOne{filePath: filePath}
	one.initLLUnits()
	return one
}

func (one *LLOne) initLLUnits() {
	lines := file.NewFileReader(one.filePath).GetFileLines()
	for _, line := range lines {
		unit := &llUnit{}
		unit.Parse(line)
		one.llUnits = append(one.llUnits, unit)
	}
}
