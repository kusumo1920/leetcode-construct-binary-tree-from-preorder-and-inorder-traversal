package main

import "fmt"

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	output := buildTree(preorder, inorder)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	mapper := make(map[int]int)
	preidx := 0

	var recursiveFn func(int, int) *TreeNode
	recursiveFn = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := preorder[preidx]
		preidx++
		index := mapper[rootVal]
		root := &TreeNode{Val: rootVal}
		root.Left = recursiveFn(left, index-1)
		root.Right = recursiveFn(index+1, right)
		return root
	}

	for i := 0; i < len(inorder); i++ {
		mapper[inorder[i]] = i
	}

	return recursiveFn(0, len(inorder)-1)
}
