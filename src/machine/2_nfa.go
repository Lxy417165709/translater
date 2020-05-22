package machine

type NFA struct {
	// 构建NFA相关的成员变量
	startState            *state
	endState              *state
	specialChar byte
}

func NewNFA(ordinaryChar byte) *NFA {
	startState:=NewState(false)
	endState:=NewState(true)
	startState.next[ordinaryChar] = append(startState.next[ordinaryChar],endState)
	return &NFA{
		startState:startState,
		endState:endState,
	}
}
func NewEmptyNFA() *NFA{
	return &NFA{
		startState:NewState(false),
		endState:NewState(true),
	}
}


func (nfa *NFA) setSpecialChar(specialChar byte){

}

















