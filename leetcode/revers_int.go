package leetcode

import (
	"math"
)

func reverse(x int) int {
	var rev int
	t := int(math.Abs(float64(x)))
	if (t == 0 || t > math.MaxInt32) {
		return 0
	}
	for i := t ; i >0;  i = i/10{
		rev = rev + int(float64(i%10) *  math.Pow(10, math.Floor(math.Log10(float64(i)))))
	}
	return rev * (x/int(t))
}

func reverse2(x int) int {
	result := 0
	for ;x != 0; x = x/10 {
		t := x % 10
		newResult := result*10 + t
		if ((newResult-t)/10 != result) {
			return 0
		}
		result = newResult;
	}
	return result
}

func ReverseR(x int, accum int) int {
	if x == 0 {
		return accum
	}
	return ReverseR(x/10, accum + int(float64(x%10)*math.Pow(10, math.Floor(math.Log10(float64(x))))))
}


func Reverse(x int) int {
	//return ReverseR(x,0)
	return reverse2(x)
}