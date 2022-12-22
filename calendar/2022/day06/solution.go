package day06

func part1(input string) int {
	var i int
	var marker = 4
	for ; i < len(input); i++ {
		tmp := make(map[byte]int, marker)
		for j := 0; j < marker; j++ {
			if _, ok := tmp[input[i+j]]; ok {
				break
			} else {
				tmp[input[i+j]]++
			}
		}
		if len(tmp) == marker {
			break
		}
	}
	return i + marker
}

func part2(input string) int {
	var i int
	var marker = 14
	for ; i < len(input); i++ {
		tmp := make(map[byte]int, marker)
		for j := 0; j < marker; j++ {
			if _, ok := tmp[input[i+j]]; ok {
				break
			} else {
				tmp[input[i+j]]++
			}
		}
		if len(tmp) == marker {
			break
		}
	}
	return i + marker
}
