package machine

import (
	"grammar"
)

type state struct {
	isEnd             bool
	next         map[byte][]*state
	specialChar byte
}

func NewState(isEnd bool) *state{
	return &state{
		isEnd:isEnd,
		next:make(map[byte][]*state),
		specialChar:grammar.Eps,
	}
}

func (s *state) linkByOrdinaryChar(ordinaryChar byte,aimState *state) {
	s.next[ordinaryChar] = append(s.next[ordinaryChar],aimState)
}
func (s *state) linkByEpsChar(aimState *state) {
	s.next[grammar.Eps] = append(s.next[grammar.Eps],aimState)
}

func (s *state) getAllNextStates() []*state {
	result := make([]*state, 0)
	for char := range s.next {
		result = append(result, s.getNextStates(char)...)
	}
	return result
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
