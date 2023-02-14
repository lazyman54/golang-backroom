package collection

import (
	"testing"
)

func TestSliceData(t *testing.T) {
	sliceData(t)
}

func sliceData(t *testing.T) {
	nums := make([]int, 0, 8)
	nums = append(nums, 1, 2, 3, 4, 5)
	nums2 := nums[2:4]

	t.Logf("nums len:%d, cap:%d, val:%v", len(nums), cap(nums), nums)
	t.Logf("num2 len:%d, cap:%d, val:%v", len(nums2), cap(nums2), nums2)

	nums2 = append(nums2, 50, 60)

	t.Logf("nums len:%d, cap:%d, val:%v", len(nums), cap(nums), nums)
	t.Logf("num2 len:%d, cap:%d, val:%v", len(nums2), cap(nums2), nums2)
	nums = append(nums, 55)
	t.Logf("num2 len:%d, cap:%d, val:%v", len(nums2), cap(nums2), nums2)

}
