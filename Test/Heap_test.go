package test

import (
	binaryheap "GoAlgorithm/myheap"
	"fmt"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	var mMax binaryheap.Maxheap
	var h = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for _, v := range h {
		mMax.MaxheapInsert(v)
	}
	fmt.Println("最大堆测试")
	mMax.PrintHeap()
	mMax.MaxheapRemove(20)
	fmt.Println("\n删除最大堆中的元素20")
	mMax.PrintHeap()

	var mMin binaryheap.Minheap
	for _, v := range h {
		mMin.MinheapInsert(v)
	}
	fmt.Println("\n最小堆测试")
	mMin.PrintHeap()
	mMin.MinheapRemove(10)
	fmt.Println("\n删除最小堆中的元素10")
	mMin.PrintHeap()
}

func TestLeftLeanHeap(t *testing.T) {
	var lMin binaryheap.LLheap
	var h = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for _, v := range h {
		lMin.Insert(v)
	}
	lMin.PrintHeap()
	var min int
	min, _ = lMin.FindMin()
	if min != 10 {
		t.Errorf("期望值为 10， 实际值为 %d", min)
	}
	lMin.PopMin()
	min, _ = lMin.FindMin()
	if min != 20 {
		t.Errorf("期望值为 20， 实际值为 %d", min)
	}
}

func TestSkewLeanHeap(t *testing.T) {
	var skewheap binaryheap.SkewHeap
	var h = [...]int{50, 30, 10, 20, 40, 60, 70, 80, 90}
	for _, v := range h {
		skewheap.Insert(v)
	}
	fmt.Println("中序遍历斜堆")
	skewheap.PrintHeap()
	fmt.Println("\n获取优先级最高的元素值")
	min, err := skewheap.GetMin()
	if err != nil {
		t.Errorf("获取优先级最高元素值失败，空堆")
	} else {
		if min != 10 {
			t.Errorf("期望值为 10 实际值为 %d", min)
		}
	}
	fmt.Println("测试删除")
	skewheap.PopMin()
	min, err = skewheap.GetMin()
	if err != nil {
		t.Errorf("获取优先级最高元素值失败，空堆")
	} else {
		if min != 20 {
			t.Errorf("期望值为 20 实际值为 %d", min)
		}
	}
}
