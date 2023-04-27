/*
Level: Hard

Given an
 - integer k
 - string s
find the length of the longest substring that contains at most k distinct characters and it also can returns an array of longest substrings
.

For example, given:
 - s = "abcba"
 - k = 2,
the longest substring with k distinct characters is "bcb".


ValorK = 3
Texto = "abcba"

*/

package main

import (
	"fmt"
)

func main() {

	subs, length := LengthOfLongestSubstringKDistinct("viiiiiiiiqva", 2)

	ExibirValor(subs, length)

}

func ExibirValor(s []string, length int) {

	fmt.Printf("Longest substrings : %v\n", s)
	fmt.Printf("Length: %v\n", length)
}

func LengthOfLongestSubstringKDistinct(s string, k int) ([]string, int) {

	stringSize := len(s)

	if stringSize*k == 0 {
		return nil, 0
	}

	if stringSize == 1 {
		return []string{s}, 1
	}

	freq := make(map[byte]int)
	maxLen := 1
	left, right := 0, 0
	longestSubstring := []string{}

	for right < stringSize {

		freq[s[right]]++

		for len(freq) > k {

			freq[s[left]]--

			if freq[s[left]] == 0 {
				delete(freq, s[left])
			}

			left++

		}
		if len(freq) <= k && right-left+1 >= maxLen {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
				longestSubstring = []string{s[left : right+1]}
			} else {
				longestSubstring = append(longestSubstring, s[left:right+1])
			}
		}

		right++
	}

	return longestSubstring, maxLen

}
