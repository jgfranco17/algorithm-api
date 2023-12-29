package sequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		mainString string
		subString  string
		expected   string
	}{
		{"ABCBDAB", "BDCAB", "BDAB"},
		{"AGGTAB", "GXTXAYB", "GTAB"},
		{"ABCD", "ACDF", "ACD"},
		{"", "XYZ", ""},
		{"ABC", "", ""},
		{"ABCD", "DCBA", "D"},
	}

	for _, test := range tests {
		result, err := longestCommonSubsequence(test.mainString, test.subString)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != test.expected {
			t.Errorf("For mainString=%s, subString=%s, expected %s but got %s", test.mainString, test.subString, test.expected, result)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 3},
		{5, 1, 5},
		{-2, 0, 0},
		{0, -1, 0},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := max(test.a, test.b)
		if result != test.expected {
			t.Errorf("For a=%d, b=%d, expected %d but got %d", test.a, test.b, test.expected, result)
		}
	}
}
