package main

import "fmt"

// 定义二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//0代表节点为空
	num := []int{1, 0, 2, 0, 0, 3, 0}
	tree := createBinaryTree(0, num)
	//二叉树数据
	//var middle []int
	isSymmetric(*tree)
	//fmt.Println(middle)
}

// 创建二叉树
func createBinaryTree(i int, nums []int) *TreeNode {
	tree := &TreeNode{nums[i], nil, nil}
	//左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.Left = createBinaryTree(2*i+1, nums)
	}
	//右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.Right = createBinaryTree(2*i+2, nums)
	}
	return tree
}

// 中序遍历
func isSymmetric(root TreeNode) bool {
	fmt.Println("test")
	if root.Left.Val != root.Right.Val {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	return isSymmetric(*root.Left) && isSymmetric(*root.Right)
}
