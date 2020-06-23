package env



type CompilerConf struct {
	SynAnalyzerConf *SynAnalyzerConf `json:"SynAnalyzerConf"`
	LexAnalyzerConf *LexAnalyzerConf `json:"LexAnalyzerConf"`
}
