package machine


func (s *state) getNextStates(char byte) []*state {
	return s.next[char]
}
func (s *state) getAllNextStates() []*state {
	result := make([]*state, 0)
	for char := range s.next {
		result = append(result, s.getNextStates(char)...)
	}
	return result
}
