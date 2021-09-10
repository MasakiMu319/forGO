package main

import (
	. "fmt"
)

func main()  {
	Println(chalkReplacer([]int{5,1,5}, 22))
}

func chalkReplacer(chalk []int, k int) (ans int) {
	n, sum := len(chalk), 0
	for i := 0; i < n; i++ {
		sum += chalk[i]
	}
	k -= sum*(k/sum)
	for i := 0; i < n; i++ {
		if k < chalk[i] {
			ans = i
			break
		}
		k -= chalk[i]
	}
	return ans
}