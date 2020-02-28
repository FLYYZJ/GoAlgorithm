package mysort

import "fmt"

// BucketSort 桶排序
// 假设待排序的数组a中共有N个整数，并且已知数组a中数据的范围[0, MAX)。
// 在桶排序时，创建容量为MAX的桶数组r，并将桶数组元素都初始化为0；将容量为MAX的桶数组中的每一个单元都看作一个"桶"。
func BucketSort(a []int, max int) {
	if len(a) < 1 || max < 1 {
		fmt.Println("该列表不需要排序")
	}
	buckets := make([]int, max)
	for i := 0; i < len(a); i++ {
		buckets[a[i]]++
	}
	for i, j := 0, 0; i < max; i++ {
		for buckets[i] > 0 {
			buckets[i], a[j], j = buckets[i]-1, i, j+1
		}
	}
}
