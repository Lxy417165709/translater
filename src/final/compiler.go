package final

import (
	"env"
	"lex"
	"syn"
)

type Compiler struct {
	lexAnalyzer *lex.Analyzer
	synAnalyzer *syn.Analyzer
}

func NewCompiler(conf *env.CompilerConf) (*Compiler, error) {
	lexAnalyzer, err := lex.NewAnalyzer(conf.LexAnalyzerConf)
	if err != nil {
		return nil, err
	}

	synAnalyzer, err := syn.NewAnalyzer(conf.SynAnalyzerConf)
	if err != nil {
		return nil, err
	}
	return &Compiler{
		lexAnalyzer: lexAnalyzer,
		synAnalyzer: synAnalyzer,
	}, nil

}

func (cp *Compiler) ExecCode(code string) error {
	tokens, err := cp.lexAnalyzer.GetTokens([]byte(code))
	if err != nil {
		return err
	}
	err = cp.synAnalyzer.GetSyntaxTree(tokens)
	if err != nil {
		return err
	}
	return cp.synAnalyzer.ParseSyntaxTree()
}
