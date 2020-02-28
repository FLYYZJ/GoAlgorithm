package mysort

// ShellSort 希尔排序（插入排序改进）
// 分组插入方法。它的基本思想是：对于n个待排序的数列，取一个小于n(整体长度)的整数gap(gap被称为步长)将待排序元素分成若干个组子序列，
// 所有距离为gap的倍数的记录放在同一个组中；然后，对各组内的元素进行直接插入排序。 这一趟排序完成之后，每一个组的元素都是有序的。
// 然后减小gap的值，并重复执行上述的分组和排序。重复这样的操作，当gap=1时，整个数列就是有序的。
func ShellSort1(a []int) {
	var i, j, gap int
	lena := len(a)
	for gap = lena / 2; gap > 0; gap /= 2 {
		// 拿gap之后每gap个元素插入gap之前的列表中
		for i = 0; i < gap; i++ {
			for j = gap + i; j < lena; j += gap {
				// 相当于i, i+gap, i+2gap, ... 组成了一个比较的列表
				// 如果a[j] < a[j-gap]，则寻找a[j]位置，并将后面数据的位置都后移。插入排序的变型
				if a[j] < a[j-gap] { // 小于a[j]的元素不断后移，挪出前面的位置给a[j]
					tmp, k := a[j], j-gap
					for k >= 0 && a[k] > tmp {
						a[k+gap], k = a[k], k-gap
					}
					a[k+gap] = tmp
				}
			}
		}
	}

}

// ShellSort2 希尔排序
func ShellSort2(a []int) {
	n := len(a)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := 0; i < gap; i++ {
			groupSort(a, i, gap)
		}
	}
}

func groupSort(a []int, i, gap int) {
	// i 为组的起始位置，gap组的步长， n 组的终点限制
	for j := i + gap; j < len(a); j += gap {
		if a[j] < a[j-gap] {
			tmp, k := a[j], j-gap
			for k >= 0 && a[k] > tmp {
				a[k+gap], k = a[k], k-gap
			}
			a[k+gap] = tmp
		}
	}
}
