### 问题
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

### 解答
二叉树的问题使用递归，递归方法一般包含三步

```
func result（root,tmp）{
	if root ==nil{
		return
	}
	if tmp + root.value == 满足条件{
		// 1. 判断存在性问题
		tmp == true
		return
		// 2. 统计所有满足的情况
		tmp ==append(tmp,当前)
	}
	if root.Left !=nil{
		result(root.Left,tmp)
	}
	if root.Right !=nil{
		result(root.Right,tmp)
	}
}

```