package main

import "fmt"

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	if max > min && max > 1 {
		for i := 0; i < len(nums); i++ {
			m := (max + min) / 2
			if nums[m] == target {
				return m
			} else if nums[m] > target {
				max = m - 1
			} else if nums[m] < target {
				min = m + 1
			}
		}
	} else if len(nums) == 2 {
		if nums[0] == target {
			return 0

		} else if nums[1] == target {
			return 1
		}
	} else if len(nums) == 1 {
		if nums[0] == target {
			return 0

		}
	}
	return -1
}
func main() {
	fmt.Println("search -", search([]int{4, 6}, 5))
}
