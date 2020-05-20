package stateMachine

import "grammar"

func (nfa *NFA) GetAllFiltratedTokens() []*Token{
	unFiltratedTokens := nfa.startState.GetAllTokensFromHere("",make(map[*State]bool))
	filtratedTokens := make([]*Token,0)
	variableCharHasAdded := make(map[byte]bool)
	for i:=0;i<len(unFiltratedTokens);i++{

		if grammar.GetRegexpsManager().IsVariableChar(unFiltratedTokens[i].specialChar){
			if variableCharHasAdded[unFiltratedTokens[i].specialChar]{
				continue
			}
			variableCharHasAdded[unFiltratedTokens[i].specialChar]=true

			unFiltratedTokens[i].value = ""
		}
		filtratedTokens = append(filtratedTokens,unFiltratedTokens[i])
	}
	return filtratedTokens
}
