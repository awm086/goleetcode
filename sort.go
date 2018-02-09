package main

import "fmt"

/**
 * bubble sort keep looping over an array of ints making necessary
 * swaps until there are no more swaps to do.
 */
func bsort(unsorted []int) []int {
	fmt.Println("Sorting:")
	sorted := false
	checks := 0
	for !sorted {
		swapped := false
		for i := 0; i < len(unsorted)-1; i++ {
			checks++
			if unsorted[i] > unsorted[i+1] {
				unsorted[i], unsorted[i+1] = unsorted[i+1], unsorted[i];
				swapped = true
			}
		}
		if swapped == false {
			sorted = true
		}
	}
	fmt.Println("NUmber of checks: ", checks)
	return unsorted;
}

func bsort2(unsorted []int) []int {
	fmt.Println("Sorting:")
	sorted := false
	n := len(unsorted)
	checks := 0
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			checks++;
			if unsorted[i] > unsorted[i+1] {
				unsorted[i], unsorted[i+1] = unsorted[i+1], unsorted[i];
				swapped = true
			}
		}
		// when the inner loop is done, the ith element is
		// in its correct poss
		n = n-1
		if swapped == false {
			sorted = true
		}

	}

	fmt.Println("NUmber of checks: ", checks)
	return unsorted;
}

func main() {

	unsorted := []int{3, 2, 1, 10, 15,4,1}
	for _, i := range unsorted {
		fmt.Println(i);
	}
	fmt.Println("------------")

	bsort(unsorted)

	for _, i := range unsorted {
		fmt.Println(i);
	}
}
