package main

import "fmt"

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
	var middle []int
	sum = sumOfLeftLeaves(*root, sum)
	fmt.Println(middle)
}

// 创建二叉树
func createBinaryTree(i int, nums []int) *TreeNode {
	root := &TreeNode{nums[i], nil, nil}
	//左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		root.Left = createBinaryTree(2*i+1, nums)
	}
	//右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		root.Right = createBinaryTree(2*i+2, nums)
	}
	return root
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	var LeftNum, RightNum, num []int
	num = make([]int, 0)
	if root == nil {
		return num
	}
	if root.Left != nil {
		fmt.Println(root.Left)
		LeftNum = inorderTraversal(root.Left)
		for _, value := range LeftNum {
			num = append(num, value)
		}
	}
	if string(root.Val) != "" {
		num = append(num, root.Val)
	}
	if root.Right != nil {
		RightNum = inorderTraversal(root.Right).Value
		for _, value := range RightNum {
			num = append(num, value)
			continue
		}
	}

	return num
}
