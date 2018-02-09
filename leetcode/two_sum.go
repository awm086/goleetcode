package leetcode

// all ints that sum to target
func two_sum(arr []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
		  if arr[i] + arr[j] == target {
		  	t := []int{i,j}
			  result = append(result, t)
			  break
		  }
		}
	}
	return result
}

// only one pair
func TwoSum(nums []int, target int) []int {
	results := []int{}
	for i := 0; i < len(nums) ; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				results = append(results, i,j)
				return results
			}

		}
	}
	return results;
}

func TwoSumMap(nums []int, target int) []int {

	m := make(map[int]int)

	for key, val := range(nums) {
		complement := target - nums[key]
		if _, ok := m[complement]; ok {
			return []int{m[complement], key }
		}
		m[val] = key
	}

	return []int{}
}

