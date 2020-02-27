package mytree

import (
	"fmt"
)

type node struct {
	val    int
	lchild *node
	rchild *node
}

// BinaryTree 二叉搜索树
type BinaryTree struct {
	root *node
}

// Insert 二叉搜索树插入
func (bt *BinaryTree) Insert(data int) {
	if bt.root == nil {
		bt.root = &node{val: data, lchild: nil, rchild: nil}
	} else {
		bt.insert(data)
	}
}

func (bt *BinaryTree) insert(data int) {
	p := bt.root
	for p != nil {
		if p.val > data {
			if p.lchild != nil {
				p = p.lchild
			} else {
				p.lchild = &node{val: data, lchild: nil, rchild: nil}
				return
			}
		} else {
			if p.rchild != nil {
				p = p.rchild
			} else {
				p.rchild = &node{val: data, lchild: nil, rchild: nil}
				return
			}
		}
	}
}

// Remove 二叉搜索树删除
func (bt *BinaryTree) Remove(data int) (int, error) {
	c := bt.root
	for c != nil {
		if c.lchild.val == data {
			p := c.lchild.rchild // 搜索右子树的最小值
			for p != nil && p.lchild != nil {
				p = p.lchild
			}
			c.val = p.val
			p = nil
			return c.val, nil
		} else if c.rchild.val == data {
			p := c.rchild.lchild // 找左子树的最大值
			for p != nil && p.rchild != nil {
				p = p.rchild
			}
			c.val = p.val
			p = nil
			return c.val, nil
		} else {
			if data > c.val {
				c = c.rchild
			} else {
				c = c.lchild
			}
		}
	}
	return -1, fmt.Errorf("未找到删除值 %d", data)
}

// Find 二叉搜索树查找
func (bt *BinaryTree) Find(data int) (int, error) {
	p := bt.root
	for p != nil {
		if p.val == data {
			return p.val, nil
		}
		if data > p.val {
			p = p.rchild
		} else {
			p = p.lchild
		}
	}
	return 0, fmt.Errorf("无法找到键 %d", data)
}

// MidVisit 中序遍历二叉搜索树
func (bt *BinaryTree) MidVisit() {
	if bt.root == nil {
		fmt.Println("二叉搜索树为空")
		return
	}
	bt.midVisit(bt.root)
}
func (bt *BinaryTree) midVisit(n *node) {
	if n == nil {
		return
	}
	bt.midVisit(n.lchild)
	fmt.Printf("%d  ", n.val)
	bt.midVisit(n.rchild)
}

// PreVisit 先序遍历二叉搜索树
func (bt *BinaryTree) PreVisit() {
	if bt.root == nil {
		fmt.Println("二叉搜索树为空")
		return
	}
	bt.preVisit(bt.root)
}
func (bt *BinaryTree) preVisit(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%d  ", n.val)
	bt.midVisit(n.lchild)
	bt.midVisit(n.rchild)
}

// GetRoot 获取二叉树的根
func (bt *BinaryTree) GetRoot() (int, error) {
	if bt.root == nil {
		return 0, fmt.Errorf("树为空")
	}
	return bt.root.val, nil
}
