package tree

type BTree interface {
	Tree
	Left() BTree
	Right() BTree
}
