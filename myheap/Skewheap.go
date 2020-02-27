package myheap

import "fmt"

type skewNode struct {
	val    int
	lchild *skewNode
	rchild *skewNode
}

// SkewHeap 斜堆
type SkewHeap struct {
	root *skewNode
}

func (skewh *SkewHeap) GetMin() (int, error) {
	if skewh.root == nil {
		return -1, fmt.Errorf("当前斜堆为空")
	}
	return skewh.root.val, nil
}

func (skewh *SkewHeap) merge(x *skewNode, y *skewNode) *skewNode {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	// 将x指向较小堆，y指向较大堆，此时x为新堆的根节点
	if x.val > y.val {
		x, y = y, x
	}
	// 合并较小堆根节点的右孩子 和 较大堆，得到新的右孩子指向
	x.rchild, x.lchild = x.lchild, skewh.merge(x.rchild, y)
	return x
}

// Insert 斜堆插入新值
func (skewh *SkewHeap) Insert(data int) {
	node := &skewNode{val: data, lchild: nil, rchild: nil}
	skewh.root = skewh.merge(skewh.root, node)
}

// PopMin 斜堆取出优先级最高值
func (skewh *SkewHeap) PopMin() {
	if skewh.root == nil {
		fmt.Println("该斜堆为空")
	} else {
		l, r := skewh.root.lchild, skewh.root.rchild
		skewh.root = skewh.merge(l, r)
	}

}

// 中序遍历斜堆
func (skewh *SkewHeap) PrintHeap() {
	if skewh.root == nil {
		fmt.Println("左倾堆为空")
	} else {
		skewh.midVisit(skewh.root)
		fmt.Println()
	}
}

func (skewh *SkewHeap) midVisit(root *skewNode) {
	if root == nil {
		return
	}
	skewh.midVisit(root.lchild)
	fmt.Printf("%d  ", root.val)
	skewh.midVisit(root.rchild)

}
