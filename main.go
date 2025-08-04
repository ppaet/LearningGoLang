package main

import (
	"fmt"
	"os"
)

func main() {
	fruits := []int{0, 0, 1, 1}
	totalFruit(fruits)
	os.Exit(0)
}

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		// 组合数公式
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}

func generate(numRows int) [][]int {
	var ret [][]int = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		for range numRows - i - 1 {
			fmt.Print("  ")
		}
		ret[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ret[i][j] = 1
			} else {
				ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
			}
			fmt.Print(ret[i][j], "   ")
		}
		fmt.Println()
	}
	return ret
}

func totalFruit(fruits []int) int {
	t1, t2, t3 := -1, -1, -1
	now, n2, n3 := 0, 0, 0
	maxN := 0
	for i := 0; i < len(fruits); i++ {
		if t3 == fruits[i] {
			n3++
			if t3 == t1 {
				now++
			} else {
				now = 1
			}

		} else if t2 == fruits[i] {
			n2++
			if t2 == t1 {
				now++
			} else {
				now = 1
			}
		} else {
			if n2+n3 > maxN {
				maxN = n2 + n3
			}
			if t1 == t3 && t1 != -1 {
				t2 = fruits[i]
				n2 = 1
				n3 = now
			} else {
				t3 = fruits[i]
				n3 = 1
				n2 = now
			}
			now = 1
		}
		t1 = fruits[i]
	}
	if n2+n3 > maxN {
		maxN = n2 + n3
	}
	fmt.Println("Max fruits:", maxN)
	return maxN
}
