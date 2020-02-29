package slidingwindows

// MaxSlidingWindow 求长度为k的滑动窗口的最大值
// leetcode 面试题59，https://leetcode-cn.com/problems/sliding-window-maximum/
func MaxSlidingWindow(nums []int, k int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}
	if k == 1 {
		return nums
	}
	// 构造第一个滑窗的递减队列
	// 例如 [12,3,4,123,12,3]，5的第一个滑窗的递减队列为[123,12]
	var max_queue []int
	max, max_index, n := nums[0], 0, len(nums)
	max_queue = append(max_queue, max)
	for i := 1; i < len(nums[:k]); i++ {
		if max < nums[i] {
			max = nums[i]
			max_queue = max_queue[0:0]
			max_queue = append(max_queue, max)
			max_index = i
		} else {
			if i > max_index {
				pos := len(max_queue)
				for j := pos - 1; j >= 0; j-- {
					if nums[i] > max_queue[j] {
						pos--
					} else {
						break
					}
				}
				max_queue = max_queue[0:pos]
				max_queue = append(max_queue, nums[i])
			}
		}
	}
	result = append(result, max_queue[0])
	for i := 1; i+k <= n; i++ {
		if max_queue[0] == nums[i-1] {
			max_queue = max_queue[1:]
		}
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
		result = append(result, max_queue[0])
	}
	return result
}

// 这种解法和上述的想法类似，只是辅助队列存放的是下标
// 在每次添加值后根据下标的前后关系来确定是不是最大值
// 效率和速度更快
func maxSlidingWindow2(nums []int, k int) []int {
	stack := []int{}
	ret := make([]int, 0)
	for i, v := range nums {
		j := len(stack) - 1
		// 找到新下标的插入位置
		for j = len(stack) - 1; j >= 0; j-- {
			if nums[stack[j]] > v {
				break
			}
		}
		stack = stack[0 : j+1]
		stack = append(stack, i)
		// 要把已经跳出当前滑窗大小的下标移除出stack
		for j = 0; j < len(stack); j++ {
			if stack[j] >= i-k+1 {
				break
			}
		}
		stack = stack[j:]
		// 放入当前滑窗中的最大值
		if i+1 >= k {
			ret = append(ret, nums[stack[0]])
		}
	}
	return ret
}
