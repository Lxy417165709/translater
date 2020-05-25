package machine

type WordPair struct {
	specialChar byte
	word string
}

func (wp *WordPair) GetSpecialChar() byte{
	return wp.specialChar
}
func (wp *WordPair) GetWord() string{
	return wp.word
}

func (wp *WordPair) SetWord(value string) {
	wp.word = value
}
