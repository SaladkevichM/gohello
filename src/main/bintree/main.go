package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	left  *Tree
	value int
	right *Tree
}

func Walk(t *Tree) {

	fmt.Printf("value=%v", t.value)

	if t.right != nil {
		Walk(t.right)
	}
	if t.left != nil {
		Walk(t.left)
	}

}

func randomTree(n int) *Tree {
	var t *Tree
	for _, value := range rand.Perm(n) {
		t = insert(t, value)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.value {
		t.left = insert(t.left, v)
		return t
	}
	t.right = insert(t.right, v)
	return t
}

func main() {
	tree := randomTree(10)
	Walk(tree)
}
