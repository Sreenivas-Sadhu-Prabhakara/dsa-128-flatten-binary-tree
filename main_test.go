package main

import (
	"reflect"
	"testing"
)

func chain(root *TreeNode) ([]int, bool) {
	out := []int{}
	for n := root; n != nil; n = n.Right {
		if n.Left != nil {
			return out, false // not a proper right-chain
		}
		out = append(out, n.Val)
	}
	return out, true
}

func TestFlatten(t *testing.T) {
	root := tn(1, tn(2, tn(3, nil, nil), tn(4, nil, nil)), tn(5, nil, tn(6, nil, nil)))
	Flatten(root)
	got, ok := chain(root)
	if !ok {
		t.Fatal("flattened tree still has left pointers")
	}
	want := []int{1, 2, 3, 4, 5, 6} // preorder of the original
	if !reflect.DeepEqual(got, want) {
		t.Errorf("chain = %v, want %v", got, want)
	}

	Flatten(nil) // must not panic

	single := tn(9, nil, nil)
	Flatten(single)
	if got, _ := chain(single); !reflect.DeepEqual(got, []int{9}) {
		t.Errorf("single = %v", got)
	}

	leftOnly := tn(1, tn(2, nil, nil), nil)
	Flatten(leftOnly)
	if got, _ := chain(leftOnly); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("leftOnly = %v", got)
	}
}
