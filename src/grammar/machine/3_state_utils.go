package machine


func (s *State) getNextStates(char byte) []*State {
	return s.next[char]
}
func (s *State) getAllNextStates() []*State {
	result := make([]*State, 0)
	for char := range s.next {
		result = append(result, s.getNextStates(char)...)
	}
	return result
}
