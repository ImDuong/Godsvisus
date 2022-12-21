package unittestsample_test

import (
	"godsvisus"
	"reflect"
	"testing"
)

// moveZeroes moves all zeroes to the end of slice
func moveZeroes(nums []int) {
	curP := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[curP] = nums[i]
			curP++
		}
	}
	for i := curP; i < len(nums); i++ {
		nums[i] = 0
	}
}

type (
	arrayTestV2 struct {
		arg1     []int
		expected []int
	}
)

var testcase283 = []arrayTestV2{
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{[]int{0}, []int{0}},
	{[]int{1, 2, 0, 3}, []int{1, 2, 3, 0}},
	{[]int{1, 0, 0, 0, 3, 2}, []int{1, 3, 2, 0, 0, 0}},
	{[]int{0, 0, 0, 3, 2}, []int{3, 2, 0, 0, 0}},
	{[]int{0, 1, 0, 3, 0, 12}, []int{1, 3, 12, 0, 0, 0}},
}

func Test283(t *testing.T) {
	godsvisus.Init()
	// this won't work when running unit test because fyne app only run in main goroutine
	defer godsvisus.Run()
	for idx, test := range testcase283 {
		moveZeroes(test.arg1)
		if !reflect.DeepEqual(test.arg1, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, test.arg1)
			godsvisus.CompareArrays([][]int{
				test.arg1,
				test.expected,
			})
		}
	}
}
