package machine

import "os"

func (nfa *NFA) IsMatch(pattern string) bool {
	return nfa.startState.IsMatchFromHere(pattern)
}

func (nfa *NFA) GetAllWordPairs() []*WordPair {
	return nfa.startState.GetAllWordPairsFromHere("", make(map[*State]bool))
}

func (nfa *NFA) EliminateBlankStates() *NFA {
	hasVisited := make(map[*State]bool)
	nfa.startState.EliminateNextBlankStatesFromHere(hasVisited)
	return nfa
}

func (nfa *NFA) MarkSpecialChar(specialChar byte) {
	nfa.startState.MarkSpecialCharFromHere(specialChar, make(map[*State]bool))
}

func (nfa *NFA) StoreMermaidGraphOfThisNFA(filePath string) error {
	var file *os.File
	var err error
	if file, err = os.Create(filePath); err != nil {
		return err
	}
	defer file.Close()
	for _, line := range nfa.getMermaidLines() {
		if _, err = file.WriteString(line); err != nil {
			return err
		}
	}
	return err
}
func (nfa *NFA) getMermaidLines() []string {
	lines := make([]string, 0)
	lines = append(lines, "```mermaid\ngraph LR\n")
	lines = append(lines, nfa.getMetaMermaidData()...)
	lines = append(lines, "```\n")
	return lines
}
func (nfa *NFA) getMetaMermaidData() []string {
	return nfa.startState.GetLineOfLinkInformationFromHere(0, make(map[*State]int), make(map[*State]bool))
}
