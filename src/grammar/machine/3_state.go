package machine


const Eps = byte(0)

type State struct {
	isEnd             bool
	next         map[byte][]*State
	specialChar byte



}



func NewState(isEnd bool) *State{
	return &State{
		isEnd:isEnd,
		next:make(map[byte][]*State),
		specialChar:Eps,
	}
}
func (s *State) GetSpecialChar() byte{
	return s.specialChar
}
func (s *State) GetIsEnd() bool{
	return s.isEnd
}
func (s *State) GetNext() map[byte][]*State{
	return s.next
}

func (s *State) linkByOrdinaryChar(ordinaryChar byte,aimState *State) {
	s.next[ordinaryChar] = append(s.next[ordinaryChar],aimState)
}
func (s *State) linkByEpsChar(aimState *State) {
	s.next[Eps] = append(s.next[Eps],aimState)
}




func getStatesToNext(States []*State) map[byte][]*State {
	result := make(map[byte][]*State)
	hasExist := make(map[byte]map[*State]bool)
	for _, stat := range States {
		for char, nextStates := range stat.next {
			for _, nextState := range nextStates {
				if hasExist[char] == nil {
					hasExist[char] = make(map[*State]bool)
				}
				if hasExist[char][nextState] {
					continue
				}
				hasExist[char][nextState] = true
				result[char] = append(result[char], nextState)
			}
		}
	}
	return result
}
