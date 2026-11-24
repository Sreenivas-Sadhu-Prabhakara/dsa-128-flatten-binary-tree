package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Flatten rearranges the tree in place into a right-leaning chain in
// preorder. Visiting in REVERSE preorder (right, left, root) and threading
// each node to the previously visited one builds the chain back to front.
func Flatten(root *TreeNode) {
	var prev *TreeNode
	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		walk(n.Right)
		walk(n.Left)
		n.Right = prev
		n.Left = nil
		prev = n
	}
	walk(root)
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, tn(3, nil, nil), tn(4, nil, nil)), tn(5, nil, tn(6, nil, nil)))
	Flatten(root)
	for n := root; n != nil; n = n.Right {
		fmt.Print(n.Val, " ")
	}
	fmt.Println()
}
