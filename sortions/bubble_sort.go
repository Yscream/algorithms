package main

import "fmt"

func main() {
	nums := []int{8, 3, 2, 11, 25, 18, 1}
	bubbleSort1(nums)
	bubbleSort2(nums)
}

// via "while" loop
// output:
// [3 8 2 11 25 18 1]
// [3 2 8 11 25 18 1]
// [3 2 8 11 25 18 1]
// [3 2 8 11 25 18 1]
// [3 2 8 11 18 25 1]
// [3 2 8 11 18 1 25]
// [2 3 8 11 18 1 25]
// [2 3 8 11 18 1 25]
// [2 3 8 11 18 1 25]
// [2 3 8 11 18 1 25]
// [2 3 8 11 1 18 25]
// [2 3 8 11 1 18 25]
// [2 3 8 11 1 18 25]
// [2 3 8 11 1 18 25]
// [2 3 8 11 1 18 25]
// [2 3 8 1 11 18 25]
// [2 3 8 1 11 18 25]
// [2 3 8 1 11 18 25]
// [2 3 8 1 11 18 25]
// [2 3 8 1 11 18 25]
// [2 3 1 8 11 18 25]
// [2 3 1 8 11 18 25]
// [2 3 1 8 11 18 25]
// [2 3 1 8 11 18 25]
// [2 3 1 8 11 18 25]
// [2 1 3 8 11 18 25]
// [2 1 3 8 11 18 25]
// [2 1 3 8 11 18 25]
// [2 1 3 8 11 18 25]
// [2 1 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
// [1 2 3 8 11 18 25]
func bubbleSort1(nums []int) []int {
	var isSorted bool

	for !isSorted {
		isSorted = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				isSorted = false
			}
			fmt.Println(nums)
		}
	}

	return nums
}

// via 2 "for" loop
// output:
// [3 2 8 11 18 1 25]
// [2 3 8 11 1 18 25]
// [2 3 8 1 11 18 25]
// [2 3 1 8 11 18 25]
// [2 1 3 8 11 18 25]
// [1 2 3 8 11 18 25]
func bubbleSort2(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		fmt.Println(nums)
	}

	return nums
}