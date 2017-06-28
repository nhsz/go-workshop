package binarytree

// Tree is a binary tree with integer values
type Tree struct {
	left  *Tree
	value int
	right *Tree
}

// GenerateBinaryTreeFromList returns a Tree generated from []int
func GenerateBinaryTreeFromList(list []int) *Tree {
	var aBinaryTree *Tree

	for _, v := range list {
		if aBinaryTree == nil {
			aBinaryTree = &Tree{nil, v, nil}
		} else {
			insertNode(aBinaryTree, v)
		}
	}

	return aBinaryTree
}

// Insert node (int value) in a Tree
func insertNode(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v < t.value {
		t.left = insertNode(t.left, v)
	} else {
		t.right = insertNode(t.right, v)
	}

	return t
}

// GenerateListFromBinaryTree returns a []int generated from a Tree
func GenerateListFromBinaryTree(t *Tree) []int {
	var list []int

	if t.left != nil {
		list = append(list, GenerateListFromBinaryTree(t.left)...)
	}

	list = append(list, t.value)

	if t.right != nil {
		list = append(list, GenerateListFromBinaryTree(t.right)...)
	}

	return list
}
