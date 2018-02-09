package main

import (
	 leetcode "github.com/awm086/algos/leetcode"
	"fmt"
)

func main () {

	//https://leetcode.com/problems/two-sum/description/
	fmt.Println("Given nums = [2, 7, 11, 15], target = 9 the indices are:")
	fmt.Println(leetcode.TwoSum([]int{2, 7, 11, 15},  9))

	fmt.Println("add up 2 lists of number 1->3->4->5 (5431 + 5431)")
	// https://leetcode.com/problems/add-two-numbers/description/
	leetcode.CreatAndAddTwoList()

	//fmt.Println()
	//fmt.Println(leetcode.Reverse(189))
	fmt.Println("")
	fmt.Println(leetcode.Reverse(90000))

}
