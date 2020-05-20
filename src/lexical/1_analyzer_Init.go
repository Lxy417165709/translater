package lexical

import (
	"conf"
	"stateMachine"
)

func Init(lexicalConf *conf.LexicalConf) {
	globalLexicalAnalyzer.lexicalConf = lexicalConf
	globalLexicalAnalyzer.Init()
}
func (la *LexicalAnalyzer) Init() {
	la.initNFAs()
	la.initFinalNFA()
	la.initLexicalDocumentGenerator()
}
func (la *LexicalAnalyzer) initNFAs() {
	for i := 0; i < len(la.lexicalConf.SpecialCharsOfNFAs); i++ {
		specialChar := la.lexicalConf.SpecialCharsOfNFAs[i]
		nfa := stateMachine.NewNFABuilder(string(specialChar)).BuildNFA()
		la.nfas = append(la.nfas, nfa.EliminateBlankStates().MarkSpecialChar())
	}

}
func (la *LexicalAnalyzer) initFinalNFA() {
	la.finalNfa = stateMachine.NewNFA()
	for i := 0; i < len(la.nfas); i++ {
		la.finalNfa.AddParallelNFA(la.nfas[i])
	}
	la.finalNfa.EliminateBlankStates()
}
func (la *LexicalAnalyzer) initLexicalDocumentGenerator() {
	la.lexicalDocumentGenerator = NewlexicalDocumentGenerator(la.lexicalConf)
}

