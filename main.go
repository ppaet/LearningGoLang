package main

import (
	"fmt"
	"os"
)

func main() {
	var numRows int
	fmt.Print("Enter the number of rows for Pascal's Triangle: ")
	fmt.Scanln(&numRows)
	generate(numRows)
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
		for range (numRows - i - 1) {
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
