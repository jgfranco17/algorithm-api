package sequence

func longestCommonSubsequence(mainString string, subString string) (string, error) {
	m := len(mainString)
	n := len(subString)

	// Create a 2D slice to store the LCS lengths
	lcsLengths := make([][]int, m+1)
	for i := range lcsLengths {
		lcsLengths[i] = make([]int, n+1)
	}

	// Build the LCS lengths matrix
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				lcsLengths[i][j] = 0
			} else if mainString[i-1] == subString[j-1] {
				lcsLengths[i][j] = lcsLengths[i-1][j-1] + 1
			} else {
				lcsLengths[i][j] = max(lcsLengths[i-1][j], lcsLengths[i][j-1])
			}
		}
	}

	// Reconstruct the LCS from the matrix
	var lcs string
	i, j := m, n
	for i > 0 && j > 0 {
		if mainString[i-1] == subString[j-1] {
			lcs = string(mainString[i-1]) + lcs
			i--
			j--
		} else if lcsLengths[i-1][j] > lcsLengths[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcs, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
