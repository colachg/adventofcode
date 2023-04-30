package day11

import (
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id         int
	items      []int
	opt        func(int) int
	divisor    int
	test       func(int) bool
	throwTrue  int
	throwFalse int
	activity   int
}

func parseItems(line string) []int {
	itemsStr := strings.Split(line, ":")[1]
	items := make([]int, 0, len(itemsStr))
	for _, itemStr := range strings.Split(itemsStr, ",") {
		item, _ := strconv.Atoi(strings.TrimSpace(itemStr))
		items = append(items, item)
	}
	return items
}

// parseOperation get the operation func
func parseOperation(line string) func(int) int {
	operation := strings.Split(line, ": new = old ")[1]
	ops := strings.Split(operation, " ")

	if ops[1] == "old" {
		switch ops[0] {
		case "+":
			return func(old int) int { return old + old }
		case "*":
			return func(old int) int { return old * old }
		default:
			return nil
		}
	}

	num, _ := strconv.Atoi(ops[1])

	switch ops[0] {
	case "+":
		return func(old int) int { return old + num }
	case "*":
		return func(old int) int { return old * num }
	default:
		return nil
	}
}

// parseTest get the test func
func parseTest(line string) func(i int) bool {
	divisor := strings.Split(line, " divisible by ")[1]
	d, _ := strconv.Atoi(divisor)

	return func(i int) bool {
		return i%d == 0
	}
}

// parseMonkey get the monkey id
func parseDivisor(line string) int {
	s := strings.Split(line, " divisible by ")
	m, _ := strconv.Atoi(s[1])
	return m
}

// parseMonkey get the monkey id
func parseMonkey(line string) int {
	monkey := strings.Split(line, " throw to monkey ")[1]
	m, _ := strconv.Atoi(monkey)
	return m
}

// setupMonkeys read from the input to create monkeys
func setupMonkeys(input string) []monkey {
	lines := strings.Split(input, "\n")
	monkeys := make([]monkey, len(lines)/7+1)
	for i, line := range lines {
		switch i % 7 {
		case 0:
			monkeys[i/7].id = i / 7
		case 1:
			monkeys[i/7].items = parseItems(line)
		case 2:
			monkeys[i/7].opt = parseOperation(line)
		case 3:
			monkeys[i/7].test = parseTest(line)
			monkeys[i/7].divisor = parseDivisor(line)
		case 4:
			monkeys[i/7].throwTrue = parseMonkey(line)
		case 5:
			monkeys[i/7].throwFalse = parseMonkey(line)
		}
	}
	return monkeys
}

// inspect to inspect all items and divided by 3
func (m *monkey) inspect(base int) []int {
	var res []int
	for _, item := range m.items {
		if base == 1 {
			res = append(res, m.opt(item)/3)
		} else {
			res = append(res, m.opt(item)%base)
		}
		m.activity++
	}
	return res
}

func part1(input string, round int) int {
	monkeys := setupMonkeys(input)
	for i := 0; i < round; i++ {
		for j := range monkeys {
			m := &monkeys[j]
			for _, v := range m.inspect(1) {
				if m.test(v) {
					monkeys[m.throwTrue].items = append(monkeys[m.throwTrue].items, v)
				} else {
					monkeys[m.throwFalse].items = append(monkeys[m.throwFalse].items, v)
				}
			}
			m.items = []int{}
		}
	}

	var res []int
	for _, m := range monkeys {
		res = append(res, m.activity)
	}
	sort.Ints(res)
	return res[len(res)-1] * res[len(res)-2]
}

func part2(input string, round int) int {
	base := 1
	monkeys := setupMonkeys(input)

	for _, m := range monkeys {
		base *= m.divisor
	}
	for i := 0; i < round; i++ {
		for j := range monkeys {
			m := &monkeys[j]
			for _, v := range m.inspect(base) {
				if m.test(v) {
					monkeys[m.throwTrue].items = append(monkeys[m.throwTrue].items, v)
				} else {
					monkeys[m.throwFalse].items = append(monkeys[m.throwFalse].items, v)
				}
			}
			m.items = []int{}
		}
	}

	var res []int
	for _, m := range monkeys {
		res = append(res, m.activity)
	}
	sort.Ints(res)
	return res[len(res)-1] * res[len(res)-2]
}
