package machine


import "os"

func (nfa *NFA) StoreMermaidGraphOfThisNFA(filePath string) error {
	var file *os.File
	var err error
	if file, err = os.Create(filePath);err!=nil{
		return err
	}
	defer file.Close()
	for _,line := range nfa.getMermaidLines() {
		if _,err = file.WriteString(line);err!=nil{
			return err
		}
	}
	return err
}
func (nfa *NFA) getMermaidLines() []string{
	lines := make([]string,0)
	lines = append(lines, "```mermaid\ngraph LR\n")
	lines = append(lines,nfa.getMetaMermaidData()...)
	lines = append(lines,"```\n")
	return lines
}
func (nfa *NFA) getMetaMermaidData() []string{
	return nfa.startState.GetLineOfLinkInformationFromHere(0, make(map[*state]int), make(map[*state]bool))
}
