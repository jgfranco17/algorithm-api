package data

import (
	"github.com/jgfranco17/algorithm-api/service/pkg/env"
)

var About = AboutInfo{
	Name:        "Algorithm API",
	Author:      "Joaquin Franco",
	Repository:  "https://github.com/jgfranco17/algorithm-api",
	Version:     "0.0.1",
	Environment: env.GetApplicationEnv(),
	License:     "MIT",
	Languages:   []string{"Go"},
	Algorithms: map[string][]string{
		"array": {
			"MaxSubArray",
			"TwoSum",
		},
		"sequence": {
			"LongestCommonSubsequence",
			"Fibonacci",
		},
		"palindrome": {
			"palindrome",
		},
	},
}
