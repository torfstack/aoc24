package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	testFilename := "testinput.txt"
	part1(testFilename)
	part2(testFilename)

	filename := "input.txt"
	part1(filename)
	part2(filename)
}

func part2(filename string) {
	rules, updates := parseInput(filename)
	before := make(map[int][]int)
	for _, rule := range rules {
		before[rule.After] = append(before[rule.After], rule.Before)
	}
	sum := 0
	for _, update := range updates {
		if !isCorrect(update, before) {
			slices.SortFunc(
				update, func(i, j int) int {
					if slices.Contains(before[i], j) {
						return -1
					} else if slices.Contains(before[j], i) {
						return 1
					}
					return 0
				},
			)
			sum += update[len(update)/2]
		}
	}

	fmt.Println(sum)
}

func part1(filename string) {
	rules, updates := parseInput(filename)
	before := make(map[int][]int)
	for _, rule := range rules {
		before[rule.After] = append(before[rule.After], rule.Before)
	}
	sum := 0
	for _, update := range updates {
		if isCorrect(update, before) {
			sum += update[len(update)/2]
		}
	}

	fmt.Println(sum)
}

func isCorrect(update Update, before map[int][]int) bool {
	res := true
	for pi, pageNumber := range update {
		befores, ok := before[pageNumber]
		if !ok {
			continue
		}
		for _, b := range befores {
			if slices.Contains(update[pi+1:], b) {
				res = false
				break
			}
		}
	}
	return res
}

func parseInput(filename string) ([]Rule, []Update) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanUpdates := false
	var rules []Rule
	var updates []Update
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanUpdates = true
			continue
		}
		if scanUpdates {
			var update Update
			pageNumbers := strings.Split(line, ",")
			for _, pageNumber := range pageNumbers {
				n, err := strconv.Atoi(pageNumber)
				if err != nil {
					panic(err)
				}
				update = append(update, n)
			}
			updates = append(updates, update)
		} else {
			parts := strings.Split(line, "|")
			before, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			after, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			rules = append(rules, Rule{Before: before, After: after})
		}
	}
	return rules, updates
}

type Rule struct {
	Before int
	After  int
}

type Update []int
