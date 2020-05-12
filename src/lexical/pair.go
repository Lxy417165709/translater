package lexical


type pair struct {
	tip string
	kindFlag int
	value interface{}
}

func NewPair(tip string,kindFlag int,value interface{}) *pair{
	return &pair{tip,kindFlag,value}
}
