package main

import (
	"sort"
)

func sortUniq(s []string) []string {
	if len(s) == 0 {
		return s
	}

	sort.Strings(s)

	i := 0
	for j := 1; j < len(s); j++ {
		if i < len(s)-1 && s[i] != s[j] {
			i++
		}
		s[i] = s[j]
	}

	return s[:i+1]
}
