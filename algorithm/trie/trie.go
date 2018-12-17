package main

import "fmt"

//前缀树，借助leetcode的题加以理解

//基础题
/**
Implement Trie (Prefix Tree),
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
*/
type Trie struct {
	isLeaf bool

	childs [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	trie := this
	for _, char := range word {
		isLeaf := false
		if trie.childs[char-97] == nil {
			trie.childs[char-97] = &Trie{isLeaf: isLeaf}
		}
		trie = trie.childs[char-97]
	}
	trie.isLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	trie := this
	for _, char := range word {
		if trie.childs[char-97] == nil {
			return false
		}
		trie = trie.childs[char-97]
	}
	return trie.isLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	trie := this
	for _, char := range prefix {
		if trie.childs[char-97] == nil {
			return false
		}
		trie = trie.childs[char-97]
	}
	return true
}

func TestTrie() {
	trie := Constructor()
	trie.Insert("app")
	trie.Insert("apple")
	//trie.Insert("bear")
	//trie.Insert("add")
	//trie.Insert("jam")
	//trie.Insert("rental")
	fmt.Println(trie.Search("app"))
	trie.Search("apple") // returns true

}

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

/**
654. Maximum Binary Tree
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

1. The root is the maximum number in the array.
2. The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
3. The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
4. Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1

读懂题就是个很尴尬的问题。看了别人的解法才读懂题..就是首先找出最大值，然后在当前数组中，以这个值分隔成左右两个数组，分别作为左右子树，完了再继续找最大值..
递归求法
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[max] < nums[i] {
			max = i
		}
	}
	node := &TreeNode{Val: nums[max]}
	if max > 0 {
		node.Left = constructMaximumBinaryTree(nums[:max])
	}
	if max < len(nums)-1 {
		node.Right = constructMaximumBinaryTree(nums[max+1:])
	}
	return node
}

func (node *TreeNode) print() {
	fmt.Print(node.Val)
	fmt.Print("{")
	if node.Left != nil {
		node.Left.print()
	}
	if node.Right != nil {
		fmt.Print(",")
		node.Right.print()
	}
	fmt.Print("}")

}

/**
AVL树，平衡二叉树，左右子树的高度差不超过1。值的大小 lefts < root.val <= rights
平衡二叉树的插入，如果左右子树的高度差不超过1，就要进行旋转
旋转一次：    6       5    旋转两次:        6       6     5
           /       / \                  /       /     / \
          5       4   6                4       5     4   6
         /                              \     /
        4                                5   4
*/
func (node *TreeNode) insertAVLTree(val int) *TreeNode {
	root := node
	direction := ""
	if val < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			node.Left.insertAVLTree(val)
			direction = "left"
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			node.Right.insertAVLTree(val)
			direction = "right"
		}
	}
	//旋转位置,左右子树高度差2(没得右(-1),左高度为1)
	if direction == "left" && height(node.Left)-height(node.Right) == 2 {
		if node.Left.Val > val {
			//单旋转
			root = singleRotateWithLeft(node)
		} else {
			//双旋转
			node.Left = singleRotateWithLeft(node.Left)
			root = singleRotateWithLeft(node)
		}
	}

	if direction == "right" && height(node.Right)-height(node.Left) == 2 {
		if node.Right.Val > val {
			//单旋转
			root = singleRotateWithRight(node)
		} else {
			//双旋转
			node.Right = singleRotateWithRight(node.Right)
			root = singleRotateWithRight(node)
		}
	}

	//+1是因为自己比左右高一层
	node.Height = max(height(node.Left), height(node.Right)) + 1
	return root
}

func singleRotateWithLeft(node *TreeNode) *TreeNode {
	root := node.Left
	root.Right = node
	root.Right.Left = nil
	root.Height = max(height(root.Left), height(root.Right)) + 1
	root.Right.Height = max(height(root.Left), height(root.Right)) + 1
	return root
}

func singleRotateWithRight(node *TreeNode) *TreeNode {
	root := node.Right
	root.Left = node
	root.Left.Right = nil
	root.Height = max(height(root.Left), height(root.Right)) + 1
	root.Right.Height = max(height(root.Left), height(root.Right)) + 1
	return root
}

func max(number1, number2 int) int {
	if number1 > number2 {
		return number1
	}
	return number2
}

func height(node *TreeNode) int {
	if node == nil {
		return -1
	}
	return node.Height
}

func main() {
	//TestTrie()
	//constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}).print()

	root := &TreeNode{Val: 6}
	root = root.insertAVLTree(5)
	root = root.insertAVLTree(4)
	root.print()
}
