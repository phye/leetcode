package leetcode

func lengthOfLongestSubString(s string) int {
	longest := 0
	buf := []byte(s)
	lm := make(map[byte]int)
	for i := 0; i < len(buf); i++ {
		if prev, ok := lm[buf[i]]; !ok {
			lm[buf[i]] = i
			continue
		} else {
			if len(lm) > longest {
				longest = len(lm)
			}
			lm = make(map[byte]int)
			for j := prev + 1; j <= i; j++ {
				lm[buf[j]] = j
			}
		}
	}

	// Incase there's no any repeating charaters until end of string
	if len(lm) > longest {
		longest = len(lm)
	}
	return longest
}
