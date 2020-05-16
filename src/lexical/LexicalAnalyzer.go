package lexical

//
//func foo(Type string,words []string,receiver map[string]string) {
//	for _,word := range words{
//		receiver [word]=Type
//	}
//}
//
//var GlobalLexicalAnalyzer *LexicalAnalyzer
//func init() {
//	GlobalLexicalAnalyzer = &LexicalAnalyzer{
//		stateMachine.NewNFABuilder("I").BuildDFA(),
//		stateMachine.NewNFABuilder("Z").BuildDFA(),
//		stateMachine.NewNFABuilder("O").BuildDFA(),
//		stateMachine.NewNFABuilder("J").BuildDFA(),
//	}
//}
//
//
//
//type LexicalAnalyzer struct {
//	IdentifyDFA *stateMachine.NFA
//	IntDFA *stateMachine.NFA
//	OperatorDFA *stateMachine.NFA
//	DelimiterDFA *stateMachine.NFA
//}
//
//func (la *LexicalAnalyzer) Form() {
//
//
//}
//
//
//
//func (la *LexicalAnalyzer) Parse(parsedBytes []byte) map[string][]string {
//	result := make(map[string][]string, 0)
//	wordCount := 0
//	for _, dfa := range la.DFAs {
//		for _, word := range dfa.Get(string(parsedBytes)) {
//			wordType := la.wordToType[word]
//			result[wordType] = append(result[wordType], word)
//			wordCount++
//		}
//	}
//	return result
//}
