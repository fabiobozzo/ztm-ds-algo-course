package trees

type BSTNode struct {
	Value int      `json:"value"`
	Left  *BSTNode `json:"left"`
	Right *BSTNode `json:"right"`
}

type BST interface {
	GetRoot() *BSTNode
	Insert(value int)
	Remove(value int)
	Lookup(value int) *BSTNode
}

type binarySearchTree struct {
	root *BSTNode
}

func NewBinarySearchTree() BST {
	return &binarySearchTree{}
}

func (b *binarySearchTree) GetRoot() *BSTNode {
	return b.root
}

func (b *binarySearchTree) Insert(value int) {
	newNode := BSTNode{
		Value: value,
	}

	if b.root == nil {
		b.root = &newNode
	}

	curNode := b.root
	for curNode != nil {
		if value < curNode.Value {
			if curNode.Left == nil {
				curNode.Left = &newNode
				return
			}
			curNode = curNode.Left
		} else if value > curNode.Value {
			if curNode.Right == nil {
				curNode.Right = &newNode
				return
			}
			curNode = curNode.Right
		} else {
			return
		}
	}

	return
}

func (b *binarySearchTree) Lookup(value int) *BSTNode {
	if b.root == nil {
		return nil
	}

	currentNode := b.root
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		} else if currentNode.Value < value {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}

	return nil
}

//func lookup(value int, node *BSTNode) *BSTNode {
//	if node == nil || node.Value == value {
//		return node
//	}
//
//	if value > node.Value {
//		return lookup(value, node.Right)
//	}
//
//	return lookup(value, node.Left)
//}

func (b *binarySearchTree) Remove(_ int) {
	// https://replit.com/@aneagoie/Data-Structures-Trees
}
