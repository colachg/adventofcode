package day07

import (
	_ "embed"
	"fmt"
	"path"
	"strings"
)

func parse(input string) map[string]int {
	fs, cd := map[string]int{}, ""
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		var size int
		var name string

		if strings.HasPrefix(s, "$ cd") {
			cd = path.Join(cd, strings.Fields(s)[2])
		} else if _, err := fmt.Sscanf(s, "%d %s", &size, &name); err == nil {
			for d := cd; d != "/"; d = path.Dir(d) {
				fs[d] += size
			}
			fs["/"] += size
		}
	}
	return fs
}

func part1(fs map[string]int) int {
	var result int
	for _, size := range fs {
		if size < 100000 {
			result += size
		}
	}
	return result
}

func part2(fs map[string]int) int {
	result := fs["/"] // this is important
	for _, size := range fs {
		if 70000000-fs["/"]+size >= 30000000 && size < result {
			result = size
		}
	}
	return result
}
