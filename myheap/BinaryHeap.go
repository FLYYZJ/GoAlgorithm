package myheap

import "fmt"

// Maxheap 最大堆
type Maxheap struct {
	heap []int
}

func (mh *Maxheap) MaxheapInsert(data int) {
	size := len(mh.heap) // 添加了一个元素后，heap的长度就是size
	mh.heap = append(mh.heap, data)
	mh.filterUp(size)
}

func (mh *Maxheap) MaxheapRemove(data int) int {
	if len(mh.heap) == 0 {
		return -1
	}
	index := getIndex(data, mh.heap)
	if index == -1 {
		return -1
	}
	size := len(mh.heap)
	mh.heap[index] = mh.heap[size-1]
	mh.heap = mh.heap[:size-1]

	if len(mh.heap) > 1 {
		mh.filterDown(index, len(mh.heap)-1)
	}
	return 0
}

func (mh *Maxheap) filterDown(start, end int) {
	c := start
	l := 2*c + 1
	tmp := mh.heap[c]

	for l <= end {
		if l < end && mh.heap[l] < mh.heap[l+1] {
			l++
		}
		if tmp >= mh.heap[l] {
			break
		} else {
			mh.heap[c] = mh.heap[l]
			c = l
			l = 2*l + 1
		}
	}
	mh.heap[c] = tmp

}

func (mh *Maxheap) filterUp(start int) {
	c := start
	p := (c - 1) / 2 // 父节点
	tmp := mh.heap[c]
	for c > 0 {
		if mh.heap[p] >= tmp {
			break
		} else {
			mh.heap[c] = mh.heap[p]
			c = p
			p = (p - 1) / 2
		}
	}
	mh.heap[c] = tmp
}

func (mh *Maxheap) PrintHeap() {
	for _, v := range mh.heap {
		fmt.Printf("%d  ", v)
	}
}

// Minheap 最小堆
type Minheap struct {
	heap []int
}

func (mh *Minheap) MinheapInsert(data int) {
	size := len(mh.heap) // 添加了一个元素后，heap的长度就是size
	mh.heap = append(mh.heap, data)
	mh.filterUp(size)
}

func (mh *Minheap) MinheapRemove(data int) int {
	if len(mh.heap) == 0 {
		return -1
	}
	index := getIndex(data, mh.heap)
	if index == -1 {
		return -1
	}
	size := len(mh.heap)
	mh.heap[index] = mh.heap[size-1]
	mh.heap = mh.heap[:size-1]

	if len(mh.heap) > 1 {
		mh.filterDown(index, len(mh.heap)-1)
	}
	return 0
}

func (mh *Minheap) filterUp(start int) {
	c := start
	p := (c - 1) / 2 // 父节点
	tmp := mh.heap[c]
	for c > 0 {
		if mh.heap[p] <= tmp {
			break
		} else {
			mh.heap[c] = mh.heap[p]
			c = p
			p = (p - 1) / 2
		}
	}
	mh.heap[c] = tmp
}

func (mh *Minheap) filterDown(start, end int) {
	c := start
	l := 2*c + 1
	tmp := mh.heap[c]

	for l <= end { // 取左右子节点中较小的一个
		if l < end && mh.heap[l] > mh.heap[l+1] {
			l++
		}
		if tmp <= mh.heap[l] {
			break
		} else { // 如果当前节点的值大于被选中的子节点，则进行交换
			mh.heap[c] = mh.heap[l]
			c = l
			l = 2*l + 1
		}
	}
	mh.heap[c] = tmp
}

func (mh *Minheap) PrintHeap() {
	for _, v := range mh.heap {
		fmt.Printf("%d  ", v)
	}
}

// 辅助函数
func getIndex(data int, heap []int) int {
	for i := 0; i < len(heap); i++ {
		if heap[i] == data {
			return i
		}
	}
	return -1
}
