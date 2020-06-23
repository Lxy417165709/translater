package er


type State struct {
	isEnd             bool
	next         map[byte][]*State
}

func NewState(isEnd bool) *State{
	return &State{
		isEnd:isEnd,
		next:make(map[byte][]*State),
	}
}

func (s *State) link (char byte,aimState *State) {
	s.next[char ] = append(s.next[char],aimState)
}

func (s *State) GetIsEnd() bool{
	return s.isEnd
}
