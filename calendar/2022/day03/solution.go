package day03

import "bytes"

func generate() map[rune]int {
	m := make(map[rune]int, 52)
	// a-z
	for ch, i := 'a', 1; ch <= 'z'; ch, i = ch+1, i+1 {
		m[ch] = i
	}
	//A-Z
	for ch, i := 'A', 27; ch <= 'Z'; ch, i = ch+1, i+1 {
		m[ch] = i
	}
	return m
}

func part1(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})
	var result []rune
	var sum int

outer:
	for _, line := range lines {
		for i := 0; i < len(line)/2; i++ {
			for j := len(line) / 2; j < len(line); j++ {
				if line[i] == line[j] {
					result = append(result, rune(line[i]))
					continue outer
				}
			}
		}
	}

	mp := generate()
	for _, ch := range result {
		sum += mp[ch]
	}
	return sum
}

func part2(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})

	var result []rune
	var sum int

	for i := 0; i < len(lines); i += 3 {
		m1 := make(map[byte]int)
		m2 := make(map[byte]int)
		m3 := make(map[byte]int)

		for j, ch := range lines[i] {
			m1[ch] = j
		}
		for j, ch := range lines[i+1] {
			m2[ch] = j
		}
		for j, ch := range lines[i+2] {
			m3[ch] = j
		}
		for k := range m1 {
			_, ok := m2[k]
			_, ok2 := m3[k]
			if ok && ok2 {
				result = append(result, rune(k))
				break
			}
		}
	}

	mp := generate()
	for _, ch := range result {
		sum += mp[ch]
	}
	return sum
}
