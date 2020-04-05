package sum

func diagonalDifference(arr [][]int32) int32 {
	var sum int32

	for i, n := 0, len(arr); i < n; i++ {
		sum += arr[i][i] - arr[i][n-i-1]
	}

	if sum < 0 {
		sum = -1 * sum
	}

	return sum
}
