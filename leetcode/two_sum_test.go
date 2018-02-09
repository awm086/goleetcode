package leetcode

import "testing"

func Test_twoSum(t *testing.T) {
	r := TwoSum([]int{2, 7, 11, 15}, 9)
	if len(r) != 2 || r[0] != 0 || r[1] != 1 {
		t.Fatal(r)
	}
}