package CodeTop

// 二叉树展开为链表，展开后的单链表应该与二叉树 先序遍历 顺序相同。
// 左子树的最下最右的节点，是右子树的父节点.
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		tmp := root.Left
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
