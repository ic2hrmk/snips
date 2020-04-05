package adjacent

import "strings"

func alternatingCharacters(s string) int32 {
	if len(s) == 0 {
		return 0
	}

	s = strings.ToUpper(s)

	var (
		runned = []rune(s)

		curChar  rune
		prevChar = runned[0]

		deletionCount int32
	)

	for _, curChar = range runned[1:] {
		if curChar == prevChar {
			deletionCount += 1
		} else {
			prevChar = curChar
		}
	}

	return deletionCount
}
