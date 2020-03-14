package main

import "fmt"

func main() {
	res := lengthOfLIS([]int{2, 3, 4, 5, 1, 2, 3, 4, 1, 2, 3, 4, 5, 1, 2, 7, 8, 10, 1, 2, 11, 12})
	fmt.Printf("%v\n", res)
}

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := []int{}
	dp = append(dp, nums[0])
	for i := 1; i < len(nums); i++ {
		index := binSearch(dp, nums[i])
		if index < len(dp) {
			dp[index] = nums[i]
		} else {
			dp = append(dp, nums[i])
		}
	}
	return len(dp)
}

func binSearch(nums []int, target int) int { // 返回坐标
	l, r := 0, len(nums)
	mid := (l + r) / 2
	for l < r {
		if nums[mid] < target {
			if mid+1 < r && nums[mid+1] > target {
				return mid + 1
			}
			l = mid + 1
		} else if nums[mid] > target {
			if mid-1 >= 0 && nums[mid-1] < target {
				return mid
			}
			r = mid
		} else {
			return mid
		}
		mid = (l + r) / 2
	}
	return mid
}
