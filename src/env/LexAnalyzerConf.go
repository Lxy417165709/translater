package env

type LexAnalyzerConf struct {
	SymbolTableConf *SymbolTableConf `json:"SymbolTableConf"`
	NFABuilderConf *NFABuilderConf `json:"NFABuilderConf"`
}

