package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMedianSortedArrays1([]int{}, []int{}))
	fmt.Println(findMedianSortedArrays1([]int{1, 3, 5}, []int{2, 4, 6}))
}

// merge 从小到大排序，两个切片诸位比较，小的数值添加到结果切片中
func merge(nums1, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	i, j := 0, 0
	res := make([]int, 0, l1+l2)
	for i < l1 && j < l2 {
		switch {
		case nums1[i] > nums2[j]:
			res = append(res, nums2[j])
			j++
		case nums1[i] < nums2[j]:
			res = append(res, nums1[i])
			i++
		default:
			res = append(res, nums1[i], nums2[j])
			i++
			j++
		}
	}
	if i < l1 {
		res = append(res, nums1[i:]...)
	}
	if j < l2 {
		res = append(res, nums2[j:]...)
	}
	return res
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	res := merge(nums1, nums2)
	n := len(res)
	if n == 0 {
		return -1
	}
	if n%2 == 0 {
		return float64(res[n/2-1]+res[n/2]) / 2
	}
	return float64(res[n/2])
}

func findKth(nums1, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2 = n2, n1
		nums1, nums2 = nums2, nums1
	}

	if n1 == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return int(math.Min(float64(nums1[0]), float64(nums2[0])))
	}

	k1 := int(math.Min(float64(k/2), float64(n1)))
	k2 := k - k1
	switch {
	case nums1[k1-1] < nums2[k2-1]:
		return findKth(nums1[k1:], nums2, k2)
	case nums1[k1-1] > nums2[k2-1]:
		return findKth(nums1, nums2[k2:], k1)
	default:
		return nums1[k1-1]
	}
}
