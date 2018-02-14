package leetcode

import (
	"fmt"
	"strings"
	"bytes"
)

//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring,
// "pwke" is a subsequence and not a substring.

func LengthOfLongestSubstring(str string) []string {
	var s []string
	fmt.Println(len(s))
	i , j := 0,0
	max :=0
	for i = 0; i <= len(str); i++ {
		fmt.Println("i:", i)
		//s = append(s,string(str[i]))
		for j = i + 1; j <= len(str); j++ {
			fmt.Println("i,j", i, j)
			// if we are at the end of the string
			// Otherwise check to see if the 2 consective letters are equal.
			if ( str[i] == str[j])  {
				// If they are append  j number of letters starting from i.
				// example: aab when i =0 and j = 1 we only append 1 = j letters starting
				// from i = 0. When
				// example 2: abcabc when i =0 and j =3, we extract 3 letters from 0.
				//s = append(s, str[i:j])
				if (max < len(str[i:j])) {
					max = len(str[i:j])
				}
				fmt.Println("str[", i, ",", j, "]", ":", str[i:j])
				break
			}
		}
	}

	//fmt.Println(s)
	fmt.Println("i,j", i, j)
	fmt.Println(max)
	return s
}


func LengthOfLongestSubstring2(str string) []string {
	fmt.Println(str)
	var s []string
	var t bytes.Buffer
	max := 0
	for i := 0; i < len(str); i++ {
		fmt.Println("i in the begining:", i)
		if !strings.Contains(t.String(), string(str[i])) {
			t.WriteByte(byte(str[i]))
			if max <  t.Len() {
				max = t.Len()
			}
			fmt.Println("t :", t.String())
		}else {
			fmt.Println(t.Len(), " i ", i)
			i = i - t.Len()
			fmt.Println("new i", i)
			t.Truncate(0)

		}
	}

	//fmt.Println(t.String(), "t")
	fmt.Println(max)
	return s
}


