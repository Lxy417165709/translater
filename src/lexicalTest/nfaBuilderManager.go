package lexicalTest

var GlobalNFAManager = NewNFAManager()


type NFAManager struct {
	identityToNFABuilder map[byte]string
}



func NewNFAManager() *NFAManager {
	return &NFAManager{make(map[byte]string)}
}

func (nm *NFAManager) Add(identity byte, regexp string) {
	nm.identityToNFABuilder[identity] = regexp
}
func (nm *NFAManager)Get(identity byte) string{
	return nm.identityToNFABuilder[identity]
}
func (nm *NFAManager)IdentityIsExist(identity byte) bool{
	return nm.identityToNFABuilder[identity]!=""
}


