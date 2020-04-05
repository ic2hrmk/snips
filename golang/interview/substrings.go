package interview

/**
 * @param p: the given string
 * @return: the number of different non-empty substrings of p in the string s
 */
func findSubstringInWraproundString(p string) []string {
	var (
		result []string
		pr     = []rune(p)
		cache  = make(map[string]struct{})
	)

	//
	// Waving:
	//		ab
	//		ac -> offset
	//		bc
	//

	for i := range pr {
		result = append(result, string(pr[i]))
		for j := i + 1; j < len(pr); j++ {
			cache[string(pr[i:j+1])] = struct{}{}

		}
	}

	var tmpSubstring string
	for tmpSubstring = range cache {
		result = append(result, tmpSubstring)
	}

	return result
}
