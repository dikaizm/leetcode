func lengthOfLongestSubstring(s string) int {
    if len(s) <= 1 {

		return len(s)

	}

	var longest int

	var start int

	charIndex := map[byte]int{}

	var char byte

	for i := 0; i < len(s); i++ {

		char = s[i]

		if index, found := charIndex[char]; found && index >= start {

			start = index + 1

		}

		charIndex[char] = i

		currentLength := i - start + 1

		if currentLength > longest {

			longest = currentLength

		}

	}

	return longest
}