/**
* 最大子序列和问题的4种算法
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//穷举 复杂度O(n^3)
func maxSumNum1(list []int) int {
	var maxSum, thisSum int
	maxSum = 0
	arrLen := len(list)
	for i := 0; i < arrLen; i++ {
		for j := i; j < arrLen; j++ {
			thisSum = 0
			for k := i; k <= j; k++ {
				thisSum += list[k]
			}
			if thisSum > maxSum {
				maxSum = thisSum
			}
		}
	}
	return maxSum
}

//穷举 复杂度O(n^2)
func maxSumNum2(list []int) int {
	var maxSum, thisSum int
	maxSum = 0
	arrLen := len(list)
	for i := 0; i < arrLen; i++ {
		thisSum = 0
		for j := i; j < arrLen; j++ {
			thisSum += list[j]
			if thisSum > maxSum {
				maxSum = thisSum
			}
		}
	}
	return maxSum
}

//递归法，复杂度是O(nlogn) 
func maxSumNum3(list []int) int {
	return maxSumRec(list, 0, len(list)-1)
}

func maxSumNum4(list []int) int {
	var thisSum, maxSum int
	maxSum = 0
	thisSum = 0
	arrLen := len(list)
	for i := 0; i < arrLen; i++ {
		thisSum += list[i]
		if thisSum > maxSum {
			maxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}
	return maxSum
}
func maxSumRec(list []int, left int, right int) int {
	if left == right {
		if list[left] > 0 {
			return list[left]
		} else {
			return 0
		}
	}
	center := (left + right) / 2
	maxLeftSum := maxSumRec(list, left, center)
	maxRightSum := maxSumRec(list, center+1, right)
	maxLeftBorderSum := 0
	leftBorderSum := 0
	for i := center; i >= left; i-- {
		leftBorderSum += list[i]
		if leftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = leftBorderSum
		}
	}
	maxRightBorderSum := 0
	rightBorderSum := 0
	for j := center + 1; j <= right; j++ {
		rightBorderSum += list[j]
		if rightBorderSum > maxRightBorderSum {
			maxRightBorderSum = rightBorderSum
		}
	}
	return max3(maxLeftSum, maxRightSum, maxLeftBorderSum+maxRightBorderSum)
}

func max3(a int, b int, c int) int {
	if a < b {
		a = b
	}
	if a > c {
		return a
	}
	return c
}

func main() {
	//list := []int{23, 12, -34, 56, -100, 5, 18, -92, 11, 20}
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000000; i++ {
		if r.Intn(2) == 1 {
			list = append(list, -r.Intn(100))
		} else {
			list = append(list, r.Intn(100))
		}
	}
	fmt.Println(len(list))
	//t1 := maxSumNum1(list)
	//t2 := maxSumNum2(list)
	//t3 := maxSumNum3(list)
	t4 := maxSumNum4(list)
	fmt.Println(t4)
}
