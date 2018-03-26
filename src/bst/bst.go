package bst

import (
	"fmt"
)

type BSTree struct {
	root *BST
}

type BST struct {
	key         int64
	value       interface{}
	left, right *BST
	size        int
}

func New() *BSTree {
	return &BSTree{}
}

func (b *BSTree) Size() int {
	if b == nil || b.root == nil {
		return 0
	}
	return b.root.size
}

func (root *BST) ssize() int {
	if root == nil {
		return 0
	}
	return root.size
}

func (b *BSTree) Get(key int64) interface{} {
	return b.root.get(key)
}

func (root *BST) get(key int64) interface{} {
	if root == nil {
		return nil
	}

	if root.key > key {
		return root.left.get(key)
	} else if root.key < key {
		return root.right.get(key)
	} else {
		return root.value
	}

	return nil
}

func (b *BSTree) Put(key int64, value interface{}) {
	if b == nil {
		return
	}
	b.root = put(b.root, key, value)
}

func put(b *BST, key int64, value interface{}) *BST {
	if b == nil {
		return newBST(key, value)
	}
	if b.size == 0 {
		b.key = key
		b.value = value
		b.size = 1
		return b
	}

	if b.key > key {
		b.left = put(b.left, key, value)
	} else if b.key < key {
		b.right = put(b.right, key, value)
	} else {
		b.value = value
	}
	b.size = b.left.ssize() + b.right.ssize() + 1
	return b
}

func (b *BSTree) Max() *BST {
	if b == nil {
		return nil
	}
	return b.root.max()
}

func (root *BST) max() *BST {
	if root == nil {
		return nil
	}
	if root.right == nil {
		return root
	}
	return root.right.max()
}

func (b *BSTree) Min() *BST {
	if b == nil {
		return nil
	}
	return b.root.min()
}

func (root *BST) min() *BST {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root
	}
	return root.left.min()
}

func (b *BSTree) DeleteMin() {
	if b == nil {
		return
	}
	b.root = deleteMin(b.root)
}

func deleteMin(root *BST) *BST {
	if root == nil {
		return nil
	}

	if root.left == nil {
		return root.right
	}

	root.left = deleteMin(root.left)
	root.size -= 1
	return root
}

func (b *BSTree) DeleteMax() {
	if b == nil {
		return
	}
	b.root = deleteMax(b.root)
}

func deleteMax(root *BST) *BST {
	if root == nil {
		return nil
	}

	if root.right == nil {
		return root.left
	}

	root.right = deleteMax(root.right)
	root.size -= 1
	return root
}

func (b *BSTree) Delete(key int64) {
	if b == nil {
		return
	}
	b.root = delete(b.root, key)
}

func delete(root *BST, key int64) *BST {
	if root == nil {
		return nil
	}
	if root.key > key {
		root.left = delete(root.left, key)
	} else if root.key < key {
		root.right = delete(root.right, key)
	} else {
		if root.right == nil {
			return root.left
		}
		if root.left == nil {
			return root.right
		}
		tmp := root
		root = tmp.min()
		root.left = tmp.left
		root.right = deleteMin(tmp.right)
	}
	root.size = root.right.ssize() + root.left.ssize() + 1
	return root
}

func (b *BST) Key() int64 {
	if b == nil {
		return 0
	}
	return b.key
}

func (b *BST) Value() interface{} {
	if b == nil {
		return nil
	}
	return b.value
}

func (b *BSTree) Traverse() {
	if b == nil {
		return
	}
	traverse(b.root)
}

func traverse(root *BST) {
	if root == nil {
		return
	}
	if root.left != nil {
		traverse(root.left)
	}
	fmt.Println(root.key, root.value)
	if root.right != nil {
		traverse(root.right)
	}
}

func newBST(key int64, value interface{}) *BST {
	return &BST{
		key:   key,
		value: value,
		size:  1,
	}
}
