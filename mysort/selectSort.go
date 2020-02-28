package mysort

// 选择排序的基本思想是：首先在未排序的数列中找到最小(or最大)元素，然后将其存放到数列的起始位置；
// 接着，再从剩余未排序的元素中继续寻找最小(or最大)元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
func SelectSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
}
