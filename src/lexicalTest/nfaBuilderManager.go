package lexicalTest

var GlobalNFAManager = NewNFAManager()


type NFAManager struct {
	identityToNFABuilder map[byte]*NFABuilder
}



func NewNFAManager() *NFAManager {
	return &NFAManager{make(map[byte]*NFABuilder)}
}

func (nm *NFAManager) Add(identity byte, nfaBuilder *NFABuilder) {
	nm.identityToNFABuilder[identity] = nfaBuilder
}
func (nm *NFAManager)Get(identity byte) *NFABuilder{
	return nm.identityToNFABuilder[identity]
}
func (nm *NFAManager)IdentityIsExist(identity byte) bool{
	return nm.identityToNFABuilder[identity]!=nil
}


