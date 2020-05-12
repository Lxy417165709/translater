package lexical

type trie struct {
	endFlag bool
	next    map[byte]*trie
}

func NewTrie() *trie {
	return &trie{false, make(map[byte]*trie)}
}
func (tr *trie) Insert(word string) {
	curTrie := tr
	for i := 0; i < len(word); i++ {
		if curTrie.next[word[i]] == nil {
			curTrie.next[word[i]] = NewTrie()
		}
		curTrie = curTrie.next[word[i]]
	}
	curTrie.endFlag = true
}
func (tr *trie) GetNextTrie(ch byte) *trie {
	return tr.next[ch]
}
func (tr *trie) NextTrieIsExist(ch byte) bool {
	return tr.next[ch] != nil
}
