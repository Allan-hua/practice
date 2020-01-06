package main

import "fmt"

/*
You are given the root of a binary tree,
along with two nodes, A and B.
Find and return the lowest common ancestor of A and B.
For this problem, you can assume that each node also has a pointer to its parent,
along with its left and right child.

class TreeNode:
def init(self, val):
self.left = None
self.right = None
self.parent = None
self.val = val


def lowestCommonAncestor(root, a, b):
# Fill this in.

# a
# / \
# b c
# / \
# d* e*
root = TreeNode('a')
root.left = TreeNode('b')
root.left.parent = root
root.right = TreeNode('c')
root.right.parent = root
a = root.right.left = TreeNode('d')
root.right.left.parent = root.right
b = root.right.right = TreeNode('e')
root.right.right.parent = root.right

*/

type node struct {
	left   *node
	right  *node
	parent *node
	value  string
}

func main() {
	root := &node{value: "a"}
	root.left = &node{value: "b"}
	root.left.parent = root
	root.right = &node{value: "c"}
	root.right.parent = root
	a := &node{value: "d"}
	root.right.left = a
	root.right.left.parent = root.right
	b := &node{value: "e"}
	root.left.right = b
	root.left.right.parent = root.left
	fmt.Println(lowestCommonAncestor(root, a, b))
}

type queue []*node

func lowestCommonAncestor(root, a, b *node) string {
	var q queue
	q = append(q, a.parent)
	for a != root {
		a = q[0]
		q = q[1:]
		if a == b {
			return q[len(q)-1].value
		}
		if a.left != nil {
			q = append(q, a.left)
		}
		if a.right != nil {
			q = append(q, a.right)
		}
		if a.parent != nil {
			q = append(q, a.parent)
		}
	}
	return q[len(q)-1].value
}
