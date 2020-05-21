package grammarTree

//type llTable struct{
//	Content map[string]map[string][]string
//}
//const path = `C:\Users\hasee\Desktop\Go_Practice\编译器\conf\LL1Table`
//func NewLLTable() *llTable{
//	table := &llTable{make(map[string]map[string][]string)}
//	lines := file.NewFileReader(path).GetFileLines()
//	for _,line := range lines{
//		line = strings.TrimSpace(line)
//		parts:= strings.Split(line," ")
//		lefts := strings.Split(parts[0],"|")
//		table.Add(lefts[0],lefts[1],parts[1])
//	}
//	return table
//}
//
//func (tb *llTable)Add(source string,des string,value string) {
//	if tb.Content[source]==nil{
//		tb.Content[source] = make(map[string][]string)
//	}
//	tb.Content[source][des] = strings.Split(value,"-")
//}
//
//func (tb *llTable)Get(source,des string) []string{
//	return tb.Content[source][des]
//}
