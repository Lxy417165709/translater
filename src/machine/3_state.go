package machine

type state struct {
	isEnd             bool
	next         map[byte][]*state
	specialChar byte
}

func NewState(isEnd bool) *state{
	return &state{
		isEnd:isEnd,
		next:make(map[byte][]*state),
		specialChar:Eps,
	}
}

func (s *state) linkByOrdinaryChar(ordinaryChar byte,aimState *state) {
	s.next[ordinaryChar] = append(s.next[ordinaryChar],aimState)
}
func (s *state) linkByEpsChar(aimState *state) {
	s.next[Eps] = append(s.next[Eps],aimState)
}




func getStatesToNext(states []*state) map[byte][]*state {
	result := make(map[byte][]*state)
	hasExist := make(map[byte]map[*state]bool)
	for _, stat := range states {
		for char, nextStates := range stat.next {
			for _, nextState := range nextStates {
				if hasExist[char] == nil {
					hasExist[char] = make(map[*state]bool)
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
