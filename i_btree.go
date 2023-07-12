package gotastructs

type BTree interface {
	Tree
	Left() BTree
	Right() BTree
}
