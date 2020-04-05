package special

func substrCount(n int32, s string) int64 {
	var (
		sr = []rune(s)

		hits    = int64(0)
		offset  = int32(0)
		repeats = int32(0)
	)

	for i := int32(0); i < n; i++ {
		offset = 1
		repeats = 0

		for ; i-offset >= 0 &&
			i+offset < n &&
			sr[i-offset] == sr[i-1] &&
			sr[i+offset] == sr[i-1];
		{
			hits += 1
			offset += 1
		}

		for ; i+1 < n && sr[i] == sr[i+1]; {
			i += 1
			repeats += 1
		}

		hits += int64(repeats * (repeats + 1) / 2)
	}

	return hits + int64(n)
}

/*
func substrCount(_ int32, s string) int64 {
	var (
		index               = make(map[rune][]int64)
		rIncome             = []rune(s)
		specialStringsCount = int64(0) // Each symbol is palindrome itself
	)

	// 1. Phase, single char spec symbols + indexing

	for i := range rIncome {
		index[rIncome[i]] = append(index[rIncome[i]], int64(i))
	}

	// 2. Special string counting

	lookupSpecial := func(positions []int64) int64 {
		if len(positions) == 1 {
			return 1 // Exclude single char
		}



		return 0
	}

	for i := range index {
		specialStringsCount += lookupSpecial(index[i])
	}

	return specialStringsCount
}
*/

//printByPosition := func(data []rune, positions []int) {
//	var (
//		toPrint = make([]rune, 0, len(positions))
//	)
//
//	for i := range positions {
//		toPrint = append(toPrint, data[positions[i]])
//	}
//
//	log.Println(string(toPrint))
//}
