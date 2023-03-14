package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	steps := 2

	rightRotate(nums, steps)
}

// output:
// [6 1 2 3 4 5]
// [5 6 1 2 3 4]
func rightRotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		var last int = nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
		fmt.Println(nums)
	}
}
