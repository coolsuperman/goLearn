package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/trapping-rain-water/
*/
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("%d\n", trap(height))
}

func trap(height []int) int {
	var stack []int
	ret := 0
	for n, v := range height {
		count := len(stack)
		if count == 0 {
			stack = append(stack, n)
		} else if count == 1 {
			stop := stack[count-1]
			if height[stop] < v {
				stack = []int{}
			}
			stack = append(stack, n)
		} else {
			for count > 1 {
				stop := stack[count-1]
				if height[stop] < v {
					sindex := stack[count-2]
					ret += (min(height[sindex], v) - height[stop]) * (n - sindex - 1)
					stack = stack[:count-1]
					count -= 1
				} else {
					stack = append(stack, n)
					count += 1
					break
				}
			}
			if count == 1 {
				if height[stack[count-1]] < v {
					stack = []int{n}
				} else {
					stack = append(stack, n)
				}
			}
		}
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
