package mysort

// InsertSort 插入排序思路
// 把n个待排序的元素看成为一个有序表和一个无序表。开始时有序表中只包含1个元素，无序表中包含有n-1个元素，
// 排序过程中每次从无序表中取出第一个元素，将它插入到有序表中的适当位置，使之成为新的有序表，重复n-1次可完成排序过程。
func InsertSort(a []int) {
	var i, j, k int
	lena := len(a)
	for i = 1; i < lena; i++ { // i-1 为已排序好的列表的大小
		for j = i - 1; j >= 0; j-- { // 寻找合适插入a[i]的位置
			if a[j] < a[i] {
				break
			}
		}
		if j != i-1 { // 如果不是排到最大位置，则需要将原来已排序好的i-1个
			tmp := a[i]
			for k = i - 1; k > j; k-- {
				a[k+1] = a[k]
			}
			a[k+1] = tmp
		}
	}
}
