package myheap

import (
	"fmt"
)

type llnode struct {
	element int
	npl     int
	lc      *llnode
	rc      *llnode
}

// LLheap 左倾堆
type LLheap struct {
	root *llnode
}

func (llh *LLheap) merge(l1 *llnode, l2 *llnode) *llnode {
	// 非空左倾堆 合并 空左倾堆， 直接返回空
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 比较两个根节点，取较小的根节点作为新根，合并较小根节点的右子堆和较大根节点堆
	if l1.element < l2.element {
		return llh.merge1(l1, l2)
	}
	return llh.merge1(l2, l1)
}

func (llh *LLheap) merge1(l1 *llnode, l2 *llnode) *llnode {
	if l1.lc == nil {
		l1.lc = l2
	} else {
		// 合并较小根节点的右子堆和较大根节点堆
		l1.rc = llh.merge(l1.rc, l2)
		//
		if l1.lc.npl < l1.rc.npl {
			l1.rc, l1.lc = l1.lc, l1.rc
		}
		l1.npl = l1.rc.npl + 1
	}
	return l1
}

func (llh *LLheap) Insert(val int) {
	n := llnode{element: val, lc: nil, rc: nil, npl: 0}
	llh.root = llh.merge(llh.root, &n)
}

func (llh *LLheap) FindMin() (int, error) {
	if llh.root != nil {
		return llh.root.element, nil
	}
	return 0, fmt.Errorf("the left lean heap is empty")
}

func (llh *LLheap) PopMin() (int, error) {
	if llh.root == nil {
		return 0, fmt.Errorf("左倾堆为空")
	}
	l := llh.root.lc
	r := llh.root.rc
	val := llh.root.element
	llh.root = llh.merge(l, r)
	return val, nil
}

func (llh *LLheap) PrintHeap() {
	if llh.root == nil {
		fmt.Println("左倾堆为空")
	} else {
		llh.midVisit(llh.root)
		fmt.Println()
	}
}

func (llh *LLheap) midVisit(root *llnode) {
	if root == nil {
		return
	}
	llh.midVisit(root.lc)
	fmt.Printf("%d  ", root.element)
	llh.midVisit(root.rc)

}
