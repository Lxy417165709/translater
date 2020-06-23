package syn

import (
	"dto"
	"env"
	"syn/tb"
)

type Analyzer struct {
	conf *env.SynAnalyzerConf
	stateTable *tb.StateTable

	syntaxTreeRoot  *TreeNode
	treeNodeStack   []*TreeNode
	tokens          []*dto.Token
	readingPosition int
}

func NewAnalyzer(conf *env.SynAnalyzerConf) (*Analyzer, error) {
	synAnalyzer := &Analyzer{}
	stateTable, err := tb.NewStateTable(conf.StateTableConf)
	if err != nil {
		return nil, err
	}
	synAnalyzer.stateTable = stateTable
	synAnalyzer.conf = conf
	return synAnalyzer, nil
}


func (az *Analyzer)ShowSyntaxTree() {
	az.syntaxTreeRoot.Show()
}

func (az *Analyzer)ParseSyntaxTree() error {
	return az.syntaxTreeRoot.Exec()
}
