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
		// first take...
		rev = rev + int(float64(i%10) *  math.Pow(10, math.Floor(math.Log10(float64(i)))))
	}
	return rev * (x/int(t))
}

// better math.
func reverse2(x int) int {
	if (x > math.MaxInt32) {
		return 0
	}
	num := 0
	for i :=x ;i !=0; i = i/10{
		digit := i % 10
		num = num * 10 + digit
	}
	return num
}

// Could go recursive?
func ReverseR(x int, accum int) int {
	if x == 0 {
		return accum
	}
	return ReverseR(x/10, accum + int(float64(x%10)*math.Pow(10, math.Floor(math.Log10(float64(x))))))
}

func ReverseR2(x int, accum int) int {
	if (accum > math.MaxInt32) { return 0 }
	if x == 0 { return accum }
	return ReverseR(x/10, ((accum * 10) + x %10))
}


func Reverse(x int) int {
	//return ReverseR(x,0)
	return ReverseR(x,0)
}