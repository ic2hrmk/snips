package anagrams

import (
	"math"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	if len(a) == 0 {
		return int32(len(b))
	}

	if len(b) == 0 {
		return int32(len(a))
	}

	getSet := func(data []rune) map[rune]int {
		var (
			resultSet = make(map[rune]int)
			hits      int
		)

		for i := range data {
			hits = resultSet[data[i]]
			resultSet[data[i]] = hits + 1
		}

		return resultSet
	}

	var (
		incomeA, incomeB = []rune(a), []rune(b)

		dicA              = getSet(incomeA)
		dicB              = getSet(incomeB)
		examined          = make(map[rune]struct{})
		unnecessaryExtras int32

		currChar       rune
		hitsA, hitsB   int
		isFoundInOther bool
	)

	for currChar, hitsA = range dicA {
		hitsB, isFoundInOther = dicB[currChar]
		if isFoundInOther {
			unnecessaryExtras += int32(math.Abs(float64(hitsA - hitsB)))
			examined[currChar] = struct{}{}
		} else {
			unnecessaryExtras += int32(hitsA)
		}
	}

	for currChar, hitsB = range dicB {
		_, isFoundInOther = examined[currChar]
		if !isFoundInOther {
			unnecessaryExtras += int32(hitsB)
		}
	}

	return unnecessaryExtras
}
