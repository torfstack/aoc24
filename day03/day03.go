package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	input := parseInput("input.txt")
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	do := regexp.MustCompile(`do\(\)`)
	dont := regexp.MustCompile(`don't\(\)`)
	digit := regexp.MustCompile(`\d+`)

	does := do.FindAllStringIndex(input, -1)
	donts := dont.FindAllStringIndex(input, -1)

	for _, d := range donts {
		nextDoIndex := sort.Search(
			len(does), func(i int) bool {
				return does[i][0] > d[0]
			},
		)
		if nextDoIndex == len(does) {
			input = input[:d[0]]
			break
		}
		input = input[:d[0]] + strings.Repeat("\\", does[nextDoIndex][0]-d[0]) + input[does[nextDoIndex][0]:]
	}
	input = strings.ReplaceAll(input, "\\", "")

	sum := 0
	matches := r.FindAllString(input, -1)
	for _, m := range matches {
		nums := digit.FindAllString(m, -1)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		sum += a * b
	}

	fmt.Println(sum)
}

func part1() {
	input := parseInput("input.txt")
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	d := regexp.MustCompile(`\d+`)

	sum := 0
	matches := r.FindAllString(input, -1)
	for _, m := range matches {
		nums := d.FindAllString(m, -1)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		sum += a * b
	}

	fmt.Println(sum)
}

func parseInput(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
