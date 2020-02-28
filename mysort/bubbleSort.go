package mysort

// BubbleSort 冒泡排序
func BubbleSort(a []int) {
	lena := len(a)
	for i := lena-1; i > 0; i-- {
		flag := 0
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1], flag = a[j+1], a[j], 1
			}
		}
		if flag == 0 { // 提前结束冒泡排序
			break
		}
	}
}
