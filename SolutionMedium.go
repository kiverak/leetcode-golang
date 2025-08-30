package main

func lengthOfLongestSubstring(s string) int {
	const N = 256
	last := [N]int{}
	for i := 0; i < N; i++ {
		last[i] = -1
	}

	left, maxLen := 0, 0
	for right := 0; right < len(s); right++ {
		b := s[right] // байт строки
		if last[b] >= left {
			left = last[b] + 1
		}
		last[b] = right

		if curr := right - left + 1; curr > maxLen {
			maxLen = curr
		}
	}
	return maxLen
}

func lengthOfLongestSubstringRunes(s string) int {
	last := make(map[rune]int, 256)

	left, maxLen := 0, 0
	i := 0 // индекс руны
	for _, r := range s {
		if idx, ok := last[r]; ok && idx >= left {
			left = idx + 1
		}
		last[r] = i

		if curr := i - left + 1; curr > maxLen {
			maxLen = curr
		}
		i++
	}
	return maxLen
}
