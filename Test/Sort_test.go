package test

import (
	"GoAlgorithm/mysort"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var h = []int{29, 10, 40, 20, 11, 80, 23, 30, 50, 60, 70, 90, 1}
	mysort.BubbleSort(h)
	for _, v := range h {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

func TestInsertSort(t *testing.T) {
	var h = []int{29, 10, 40, 20, 11, 80, 23, 30, 50, 60, 70, 90, 1}
	mysort.InsertSort(h)
	for _, v := range h {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

func TestShellSort(t *testing.T) {
	var h1 = []int{29, 10, 40, 20, 11, 80, 23, 30, 50, 60, 70, 90, 1}
	mysort.ShellSort1(h1)
	for _, v := range h1 {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
	var h2 = []int{12, 3, 2, 1, 31, 22, 34, 13, 98, 8, 7, 77, 65, 43, 87}
	mysort.ShellSort1(h2)
	for _, v := range h2 {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

func TestSelectSort(t *testing.T) {
	var h = []int{29, 10, 40, 20, 11, 80, 23, 30, 50, 60, 70, 90, 1}
	mysort.SelectSort(h)
	for _, v := range h {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

func TestBucketSort(t *testing.T) {
	var h = []int{29, 10, 40, 20, 11, 80, 23, 30, 50, 60, 70, 90, 1}
	mysort.BucketSort(h, 100)
	for _, v := range h {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}