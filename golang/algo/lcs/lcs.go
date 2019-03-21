package main

import (
	"fmt"
)

func lcs(X, Y string) int {
	return lcsRecursive([]rune(X), []rune(Y), len([]rune(X)), len([]rune(Y)))
}

func lcsRecursive(X, Y []rune, m, n int) int {

	if m == 0 || n == 0 {
		return 0

	} else if X[m-1] == Y[n-1] {
		return 1 + lcsRecursive(X, Y, m-1, n-1)
	}

	return max(
		lcsRecursive(X, Y, m-1, n),
		lcsRecursive(X, Y, m, n-1),
	)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(lcs("AGGTAB", "GXTXAYB"))
}
