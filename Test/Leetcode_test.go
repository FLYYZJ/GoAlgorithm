package test

import (
	slidingwindows "GoAlgorithm/leetcode/SlidingWindows"
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	res := slidingwindows.MaxSlidingWindow([]int{7, 2, 4}, 2)
	fmt.Printf("%v", res)
}
