package machine


func (s *state) getNextStates(char byte) []*state {
	return s.next[char]
}
