package trie

type Trie interface {
	Insert(text string)
	Search(text string) []string
}

func New() Trie {
	return newCompressed()
}

func (t *compressedTrie) Insert(text string) {
	t.insert(text)
}

func (t *compressedTrie) Search(text string) []string {
	return t.search(text)
}

func NewStandard() Trie {
	return newStandard()
}

func (t *trie) Insert(text string) {
	t.insert(text)
}

func (t *trie) Search(text string) []string {
	return t.search(text)
}
