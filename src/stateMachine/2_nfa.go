package stateMachine

type NFA struct {
	// 构建NFA相关的成员变量
	startState            *State
	endState              *State
	specialChar byte

	// parseToken相关的成员变量
	tokens  []*Token
	readingPosition int
	stateQueue []*State
	firstEndState *State
	parsedPattern []byte
	buffer string
}

func NewNFA() *NFA {
	startState :=NewState(false)
	endState :=NewState(true)
	return &NFA{
		startState:startState,
		endState:endState,
	}
}













