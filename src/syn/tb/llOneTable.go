package tb

import (
	"env"
	"fmt"
	"unicode"
	"utils"
)


type LLOneTable struct {
	conf *env.LLOneTableConf
	productions []*Production

	first   map[string][]string
	follow  map[string][]string
	_select map[*Sentence][]string

	indexOfHandlingProduction         int
	indexOfHandlingProductionSentence int
	bufferOfSet                       map[string][]string // first,follow集 求取过程的缓存
}

func NewLLOneTable(conf *env.LLOneTableConf) (*LLOneTable, error) {
	table := &LLOneTable{}
	lines, err := utils.GetFileLines(conf.FilePath)
	if err != nil {
		return nil, err
	}
	for index, line := range lines {
		production, err := NewProduction(line, conf.DelimiterOfPieces, conf.DelimiterOfSentences, conf.DelimiterOfSymbols)
		if err != nil {
			return nil, fmt.Errorf("%s 路径，读取第 %d 行时发生错误，%s",
				conf.FilePath,
				index+1,
				err.Error(),
			)
		}
		table.productions = append(table.productions, production)
	}
	table.conf = conf
	table.formSelect()
	return table, nil
}

// 规定非大写字符的字符串为终结符
func (one *LLOneTable) isTerminator(symbol string) bool {
	for _,char := range symbol{
		if unicode.IsUpper(char){
			return false
		}
	}
	return true
}


func (one *LLOneTable) GetSelect() map[*Sentence][]string{
	return one._select
}
func (one *LLOneTable) GetSentenceToNonTerminator() map[*Sentence]string {
	sentenceToNonTerminator := make(map[*Sentence]string)
	for _, production := range one.productions {
		for _, sentence := range production.sentences {
			sentenceToNonTerminator[sentence] = production.nonTerminator
		}
	}
	return sentenceToNonTerminator
}

func (one *LLOneTable) GetTerminators() []string {
	result := make([]string, 0)
	hasAdded := make(map[string]bool)
	for _, production := range one.productions {
		for i:=0;i<len(production.sentences);i++{
			for t:=0;t<len(production.sentences[i].Symbols);t++{
				symbol := production.sentences[i].Symbols[t]
				if !one.isTerminator(symbol) {
					continue
				}
				if hasAdded[symbol]{
					continue
				}
				hasAdded[symbol] = true
				result = append(result, symbol)
			}
		}

	}
	return result
}
func (one *LLOneTable) GetNonTerminators() []string {
	result := make([]string, 0)
	hasAdded := make(map[string]bool)
	for _, production := range one.productions {
		nonTerminator := production.nonTerminator
		if hasAdded[nonTerminator] {
			continue
		}
		hasAdded[nonTerminator] = true
		result = append(result, nonTerminator)
	}
	return result
}



