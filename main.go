package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	res := maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5)
	fmt.Printf("%v", res)
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}
	// 前k个进行排序，构成第一个最大值队列
	n, max_queue := len(nums), make(intSlice, 0)
	max_queue = append(max_queue, nums[:k]...)
	sort.Sort(sort.Reverse(max_queue))
	result = append(result, max_queue[0])
	for i := 1; i+k <= n; i++ {
		pos := len(max_queue)
		for j := pos - 1; j >= 0; j-- {
			if nums[i+k-1] > max_queue[j] {
				pos--
			} else {
				break
			}
		}
		max_queue = max_queue[0:pos]
		max_queue = append(max_queue, nums[i+k-1])
		if max_queue[0] == nums[i-1] {
			max_queue = max_queue[1:]
		}
		result = append(result, max_queue[0])
	}
	return result
}
