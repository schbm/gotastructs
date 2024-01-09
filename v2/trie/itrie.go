package trie

type Trie interface {
	Insert(string)
	Get() []int
	Remove(string)
}
