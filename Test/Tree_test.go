package test

import (
	"GoAlgorithm/mytree"
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	var binarytree mytree.BinaryTree
	fmt.Println("测试二叉树插入")
	var h = [...]int{80, 50, 20, 70, 10, 30, 40, 60, 90}
	for _, v := range h {
		binarytree.Insert(v)
	}
	fmt.Println("二叉树插入完成")
	fmt.Println("获取二叉树的根节点")
	r, _ := binarytree.GetRoot()
	fmt.Println(r)
	fmt.Println("\n中序遍历")
	binarytree.MidVisit()
	fmt.Println("\n先序遍历")
	binarytree.PreVisit()
}

func TestAVLTree(t *testing.T) {
	var avltree mytree.AVLTree
	var h = [...]int{3, 2, 1, 4, 5, 6, 7, 16, 15, 14, 13, 12, 11, 10, 8, 9}
	fmt.Println("测试AVL树插入")
	for _, v := range h {
		avltree.Insert(v)
	}
	fmt.Println("中序遍历")
	avltree.MidVisit()
	fmt.Println("\n先序遍历")
	avltree.PreVisit()
	fmt.Println("\n打印AVL树")
	avltree.PrintTree()
	fmt.Println("\n测试AVL树删除")
	avltree.Remove(10)
	avltree.MidVisit()
	fmt.Println("\n获取最大值")
	maxval, err := avltree.GetMax()
	if err != nil || maxval != 16 {
		t.Errorf("期望值为90, 实际为%d", maxval)
	}

}

func TestRBTree(t *testing.T) {
	var rbtree mytree.RedBlackTree
	var h = [...]int{10, 40, 30, 60, 90, 70, 20, 50, 80}
	//var h = [...]int{10, 5, 9, 3, 6, 7, 19, 32, 24, 17}

	fmt.Println("测试红黑树插入")
	for _, v := range h {
		rbtree.Insert(v)
	}
	fmt.Println("中序遍历")
	rbtree.MidVisit()
	fmt.Println("\n先序遍历")
	rbtree.PreVisit()
	fmt.Printf("\n测试删除 %d\n ", 40)
	rbtree.Remove(40)
	fmt.Println("先序遍历")
	rbtree.PreVisit()
}
