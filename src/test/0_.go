package test

import "fmt"

type TestName string
const (
	NFATest TestName = "NFA测试"
	SyntaxTest  TestName= "语法分析测试"
)


type Manager struct {
	testDirs []TestableDir
}

func NewManager() *Manager{
	return &Manager{make([]TestableDir,0)}
}
func (m *Manager)Register(testableDirs ...TestableDir) {
	m.testDirs = append(m.testDirs,testableDirs...)
}
func (m *Manager)BeginTest() {
	for _,testDir := range m.testDirs{
		if !testDir.Test() {
			fmt.Printf("[%v] 存在错误。\n	%s",testDir.GetTestType(),testDir.GetErrMsg())
			return
		}
		fmt.Printf("[%v] 通过。\n",testDir.GetTestType())
	}
	fmt.Printf("测试全部通过。")
}
