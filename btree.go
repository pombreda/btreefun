package btree

//  The point of this library is to implement a simple binary tree
//  that can be used generically. A la a btree template in C++.
//

type Node struct {
	Left, Right *Node
	Item
}

// convenience pointer to the root node of the tree
type Tree struct {
	Count int
	Root  *Node
}

// to use this
type Item interface {
	Less(than Item) bool
}

// make a new tree and return a pointer to it
func NewTree() *Tree {
	tree := new(Tree)
	root := new(Node)
	tree.Root = root
	tree.Count = 1
	return tree
}

func newNode(item Item) *Node {
    node := new(Node)
    node.Item = item
    return node
}

//Insert
func (t *Tree) Insert(thing Item) Item {
    
}

//Has
//Remove
//Inorder
