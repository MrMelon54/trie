package trie

type dummy struct{}

type TrieKey struct {
	Trie[dummy]
}

func (t *TrieKey) PutString(prefix string) {
	t.Trie.PutString(prefix, dummy{})
}

func (t *TrieKey) Put(newPrefix []byte) {
	t.Trie.Put(newPrefix, dummy{})
}

func (t *TrieKey) Keys(prefix string) []string {
	return t.GetAllKeys([]byte(prefix))
}
