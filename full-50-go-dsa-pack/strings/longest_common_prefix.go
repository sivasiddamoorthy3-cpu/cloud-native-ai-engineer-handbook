package strings

import "strings"

func LongestCommonPrefix(strs []string) string {
	// Handle edge cases
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// Start with the first string as the initial prefix
	prefix := strs[0]

	// Iterate through the rest of the strings
	for i := 1; i < len(strs); i++ {
		// Reduce the prefix until it matches the start of the current string
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
