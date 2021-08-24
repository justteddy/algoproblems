package tree_test

import (
	"testing"

	"leetcode/data_structures/tree"

	"github.com/stretchr/testify/assert"
)

func Test_TraverseBFS(t *testing.T) {
	assert.Equal(t, []int{20, 12, 32, 6, 16, 25, 35, 18, 40}, tree.TraverseBFS(treeExample()))
}

func Test_TraverseDFSPreOrderRecursive(t *testing.T) {
	assert.Equal(t, []int{20, 12, 6, 16, 18, 32, 25, 35, 40}, tree.TraverseDFSPreOrderRecursive(treeExample()))
}

func Test_TraverseDFSInOrderRecursive(t *testing.T) {
	assert.Equal(t, []int{6, 12, 16, 18, 20, 25, 32, 35, 40}, tree.TraverseDFSInOrderRecursive(treeExample()))
}

func Test_TraverseDFSPostOrderRecursive(t *testing.T) {
	assert.Equal(t, []int{6, 18, 16, 12, 25, 40, 35, 32, 20}, tree.TraverseDFSPostOrderRecursive(treeExample()))
}

/*
		 [20]
		/    \
     [12]    [32]
     /  \    /   \
  [6]  [16] [25]  [35]
         \           \
         [18]        [40]
*/
func treeExample() *tree.BinaryNode {
	return &tree.BinaryNode{
		Val: 20,
		Left: &tree.BinaryNode{
			Val: 12,
			Left: &tree.BinaryNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.BinaryNode{
				Val:  16,
				Left: nil,
				Right: &tree.BinaryNode{
					Val:   18,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &tree.BinaryNode{
			Val: 32,
			Left: &tree.BinaryNode{
				Val:   25,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.BinaryNode{
				Val:  35,
				Left: nil,
				Right: &tree.BinaryNode{
					Val:   40,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
}
