package token

func (tp *TokenParser) GetTokens(text []byte) []*Token{
	wordPairs := tp.finalNFA.GetWordPairs(text)
	tokens := make([]*Token,0)
	for _,wordPair := range wordPairs{
		tokens = append(tokens,tp.wordPairToToken(wordPair))
	}
	return tokens
}
