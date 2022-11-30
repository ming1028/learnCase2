package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	idxTarMap := make(map[int]int, len(nums))
	for idx, val := range nums {
		diffVal := target - val
		idxTarget, ok := idxTarMap[diffVal]
		if ok {
			return []int{idxTarget, idx}
		}
		idxTarMap[val] = idx
	}
	return nil
}
