package lexical

import (
	"fmt"
	"grammar"
	"stateMachine"
)

func init() {

	grammar.BuildGrammar()
	Z := stateMachine.NewNFABuilder("Z").BuildDFA()
	I := stateMachine.NewNFABuilder("I").BuildDFA()
	W := stateMachine.NewNFABuilder("W").BuildDFA()
	O := stateMachine.NewNFABuilder("O").BuildDFA()
	J := stateMachine.NewNFABuilder("J").BuildDFA()

	dfaToType := make(map[*stateMachine.NFA]string)
	dfaToType[Z] = "整数"
	dfaToType[I] = "标识符"
	dfaToType[W] = "关键字"
	dfaToType[O] = "操作符"
	dfaToType[J] = "界符"

	dfas := make([]*stateMachine.NFA, 0)
	dfas = append(dfas, W, I, Z, O, J)
	GlobalLexicalAnalyzer = NewLexicalAnalyzer(dfas, dfaToType)
	//fmt.Println(GlobalLexicalAnalyzer)
}

var GlobalLexicalAnalyzer *LexicalAnalyzer

type LexicalAnalyzer struct {
	DFAs      []*stateMachine.NFA
	dfaToType map[*stateMachine.NFA]string
}

func NewLexicalAnalyzer(nfas []*stateMachine.NFA, dfaToType map[*stateMachine.NFA]string) *LexicalAnalyzer {
	return &LexicalAnalyzer{nfas, dfaToType}
}

func (la *LexicalAnalyzer) Parse(parsedBytes []byte) map[string][]string {
	result := make(map[string][]string, 0)
	wordType := make(map[string]string)
	wordCount := 0
	for _, dfa := range la.DFAs {
		for _, word := range dfa.Get(string(parsedBytes)) {
			if wordType[word] == "" {
				wordType[word] = la.dfaToType[dfa]
			}
			result[wordType[word]] = append(result[wordType[word]], word)
			wordCount++
		}
	}
	fmt.Println(wordCount)
	return result
}
