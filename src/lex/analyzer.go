package lex

import (
	"bytes"
	"dto"
	"env"
	"fmt"
	"lex/er"
	"lex/tb"
)

type Analyzer struct {
	symbolTable *tb.SymbolTable
	nfaBuilder  *er.NFABuilder

	nfas []*er.NFA

	readingPosition int
	text            []byte
	preEndNFA       *er.NFA
	nfaQueue        []*er.NFA
	bufferOfChars   bytes.Buffer
	finalTokens     []*dto.Token

	nfaToSpecialChar map[*er.NFA]byte
}

func NewAnalyzer(conf *env.LexAnalyzerConf) (*Analyzer, error) {
	var err error
	var symbolTable *tb.SymbolTable
	var nfaBuilder *er.NFABuilder
	analyzer := &Analyzer{}
	if symbolTable, err = tb.NewSymbolTable(conf.SymbolTableConf); err != nil {
		return nil, err
	}
	analyzer.symbolTable = symbolTable
	if nfaBuilder, err = er.NewNFABuilder(conf.NFABuilderConf); err != nil {
		return nil, err
	}
	analyzer.nfaBuilder = nfaBuilder

	analyzer.initNFAs()
	return analyzer, nil

}
func (az *Analyzer) initNFAs() {
	az.nfaToSpecialChar = make(map[*er.NFA]byte)
	for _, specialChar := range az.symbolTable.GetSpecialCharsOfCreatingNFA() {
		nfa := az.nfaBuilder.BuildNFABySpecialChar(specialChar)
		az.nfas = append(az.nfas, nfa)
	}
}


func (az *Analyzer) GetTokens(text []byte) ([]*dto.Token, error) {
	az.text = text
	if err := az.parseTextToFinalTokens(); err != nil {
		return nil, err
	}
	return az.finalTokens, nil
}
func (az *Analyzer) parseTextToFinalTokens() error {
	az.getFinalTokensInit()
	for az.readingIsNotOver() {
		//fmt.Println(az.nfaQueue,az.nfaQueue[0].GetInner()[0].GetIsEnd())
		az.updatePreEndNFA()
		az.expandNFAQueue()
		if az.isReachParseBoundary() {
			if err := az.handleParseBoundary(); err != nil {
				return err
			}
		} else {
			az.handleParseProcess()
		}
	}
	return nil
}

func (az *Analyzer) getFinalTokensInit() {
	az.bufferOfChars = bytes.Buffer{}
	az.finalTokens = make([]*dto.Token, 0)
	az.nfaQueue = make([]*er.NFA, 0)
	az.nfaQueue = append(az.nfaQueue, az.nfas...)
	az.readingPosition = 0
	az.text = append(az.text, ' ') // a+b 这种情况，最后的字符无法识别，所以这里添加一个空白符，辅助分析
}
func (az *Analyzer) readingIsNotOver() bool {
	return az.readingPosition != len(az.text)
}
func (az *Analyzer) updatePreEndNFA() {
	az.preEndNFA = getFirstEndNFA(az.nfaQueue)
}

func (az *Analyzer) ResetNFAs() {
	for _, nfa := range az.nfas {
		nfa.ResetInnerStates()
	}
}

func (az *Analyzer) expandNFAQueue() {
	readingChar := az.text[az.readingPosition]
	tmpQueue := make([]*er.NFA, 0)
	for _, nfa := range az.nfaQueue {
		if nfa.CanChangeInnerStates(readingChar) {
			nfa.ChangeInnerStates(readingChar)
			tmpQueue = append(tmpQueue, nfa)
		}
	}
	az.nfaQueue = tmpQueue
}
func (az *Analyzer) isReachParseBoundary() bool {
	return len(az.nfaQueue) == 0
}
func (az *Analyzer) handleParseProcess() {
	az.writeReadingCharToBuffer()
	az.readNextOne()
}
func (az *Analyzer) writeReadingCharToBuffer() {
	readingChar := az.text[az.readingPosition]
	az.bufferOfChars.WriteByte(readingChar)
}
func (az *Analyzer) readNextOne() {
	az.readingPosition++
}

func (az *Analyzer) handleParseBoundary() error {
	if az.isPreEndNfaNotExist() {
		if az.readingCharIsNotBlank() {
			return fmt.Errorf("非法字符: %s, 索引位置为: %d", string(az.getReadingChar()), az.readingPosition)
		}
		az.readNextOne()
	} else {
		az.generateToken()
	}
	az.reParse()
	return nil
}

func (az *Analyzer) isPreEndNfaNotExist() bool {
	return az.preEndNFA == nil
}

func (az *Analyzer) readingCharIsNotBlank() bool {
	return !isBlank(az.getReadingChar())
}
func (az *Analyzer) reParse() {
	az.ResetNFAs()
	az.bufferOfChars = bytes.Buffer{}
	az.nfaQueue = nil
	az.nfaQueue = append(az.nfaQueue, az.nfas...)
}
func (az *Analyzer) generateToken() {

	specialChar := az.preEndNFA.GetSpecialChar()
	word := az.bufferOfChars.String()
	az.finalTokens = append(az.finalTokens, &dto.Token{
		Value :word,
		Symbol: az.symbolTable.GetTerminator(specialChar,word),
	})
}
func (az *Analyzer) getReadingChar() byte {
	return az.text[az.readingPosition]
}
func getFirstEndNFA(nfas []*er.NFA) *er.NFA {
	for _, nfa := range nfas {
		if nfa.IsEnd() {
			return nfa
		}
	}
	return nil
}
func isBlank(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}
