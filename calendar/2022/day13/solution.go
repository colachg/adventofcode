package day13

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Pair [][]interface{}

func parsePackage(line string) []interface{} {
	var pack []interface{}
	err := json.Unmarshal([]byte(line), &pack)
	if err != nil {
		panic(fmt.Errorf("%q is not unmarshable", line))
	}
	return pack
}

func parseInput(input string) []Pair {
	var output []Pair
	for _, pairStr := range strings.Split(strings.TrimSpace(input), "\n\n") {
		lines := strings.Split(pairStr, "\n")
		pair := Pair{parsePackage(lines[0]), parsePackage(lines[1])}
		output = append(output, pair)
	}
	return output
}

func toInterfaceList(v interface{}) []interface{} {
	var result []interface{}

	if numList, ok := v.([]float64); ok {
		for _, f := range numList {
			result = append(result, f)
		}
	} else if otherList, ok2 := v.([]interface{}); ok2 {
		result = append(result, otherList...)
	}
	return result
}

func isCorrectOrder(e1, e2 interface{}) (correct bool, isDecided bool) {

	num1, isNum1 := e1.(float64)
	num2, isNum2 := e2.(float64)

	// rule 1
	if isNum1 && isNum2 {
		if num1 == num2 {
			return false, false
		}
		return num1 < num2, true
	}

	// rule 3
	if isNum1 && !isNum2 {
		return isCorrectOrder([]float64{num1}, e2)
	}
	if isNum2 && !isNum1 {
		return isCorrectOrder(e1, []float64{num2})
	}

	// rule 2 (both lists)
	list1 := toInterfaceList(e1)
	list2 := toInterfaceList(e2)

	for i := 0; i < len(list1) && i < len(list2); i++ {
		result, isDifferent := isCorrectOrder(list1[i], list2[i])
		if isDifferent {
			return result, true
		}
	}
	if len(list1) == len(list2) {
		return false, false
	}
	return len(list1) < len(list2), true
}

func part1(input string) int {
	pairs := parseInput(input)

	sum := 0
	for i, pair := range pairs {
		if isCorrect, _ := isCorrectOrder(pair[0], pair[1]); isCorrect {
			sum += i + 1
		}
	}

	return sum
}

func part2(input string) int {
	pairs := parseInput(input)
	var packages [][]interface{}

	for _, pair := range pairs {
		packages = append(packages, pair[0], pair[1])
	}
	div1 := []interface{}{[]float64{2}}
	div2 := []interface{}{[]float64{6}}
	packages = append(packages, div1, div2)

	sort.Slice(packages, func(i, j int) bool {
		less, _ := isCorrectOrder(packages[i], packages[j])
		return less
	})

	result := 1
	for i, pack := range packages {
		if _, isDifferent := isCorrectOrder(pack, div1); !isDifferent {
			result *= 1 + i
		} else if _, isDifferent = isCorrectOrder(pack, div2); !isDifferent {
			result *= 1 + i
		}
	}
	return result
}
