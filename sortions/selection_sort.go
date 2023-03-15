package main

import "fmt"

func main() {
	nums := []int{8, 3, 2, 11, 25, 18, 1}
	selectionSort(nums)
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[minIdx], nums[i] = nums[i], nums[minIdx]
		fmt.Println(nums)
	}
}
