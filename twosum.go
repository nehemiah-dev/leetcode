package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		needed := target - num
		if j, ok := seen[needed]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	// Test case 1
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Test 1:", result)

	// Test case 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Println("Test 2:", result2)

	// Test case 3
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Println("Test 3:", result3)
}
