package test

import "fmt"


type Manager struct {
	testDirs []*TestableDirectory
}

func NewManager() *Manager{
	return &Manager{make([]*TestableDirectory,0)}
}
func (m *Manager)Register(testableDirs ...*TestableDirectory) {
	m.testDirs = append(m.testDirs,testableDirs...)
}
func (m *Manager)BeginTest() {
	for _,testDir := range m.testDirs{
		if !testDir.Test() {
			fmt.Printf("[%v] 存在错误。\n	%s",testDir.GetTypeOfTest(),testDir.GetErrMsg())
			return
		}
		fmt.Printf("[%v] 通过。\n",testDir.GetTypeOfTest())
	}
	fmt.Printf("测试全部通过。")
}
