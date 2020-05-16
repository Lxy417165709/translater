package stateMachine

const (
	repeatPlusSymbol = '@'
	repeatZeroSymbol = '$'
	RegexSplitString = "|"
	eps = 0
)

func NewEmptyNFA() *NFA {
	return &NFA{NewState(false), NewState(true)}
}

func NewNFA(char byte) *NFA {
	startState := NewState(false)
	endState := NewState(true)
	startState.LinkByChar(char,endState)
	return &NFA{startState, endState}
}

func getStatesToNext(states []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	hasExist := make(map[byte]map[*State]bool)
	for _, state := range states {
		for char, nextStates := range state.toNextState {
			for _,nextState := range nextStates{
				if hasExist[char]==nil{
					hasExist[char]=make(map[*State]bool)
				}
				if hasExist[char][nextState]{
					continue
				}
				hasExist[char][nextState]=true
				result[char] = append(result[char], nextState)
			}
		}
	}
	return result
}


func charToNFA(char byte) *NFA {
	if !GlobalRegexpsManager.CharIsSpecial(char) {
		return NewNFA(char)
	}

	regexp := GlobalRegexpsManager.GetRegexp(char)
	nfa := NewNFABuilder(regexp).BuildNFA()
	return nfa
}
