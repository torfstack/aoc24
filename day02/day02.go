package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	input := parseInput("input.txt")
	safeCount := 0
	for _, row := range input.Values {
		if isSafePart2(row) {
			safeCount++
		}
	}
	println(safeCount)
}

func isSafePart2(row []int) bool {
	diffs := make([]int, len(row)-1)
	for i := 0; i < len(row)-1; i++ {
		diffs[i] = row[i+1] - row[i]
		if diffs[i] == 0 || diffs[i] < -3 || 3 < diffs[i] {
			return isSafeRemovingIndex(row, i)
		}
		if i == 0 {
			continue
		}
		if diffs[i]*diffs[i-1] < 0 {
			return isSafeRemovingIndex(row, i)
		}
	}
	return true
}

func isSafeRemovingIndex(row []int, i int) bool {
	if i == 0 {
		return isSafePart1(slices.Concat(row[:i], row[i+1:])) ||
			isSafePart1(slices.Concat(row[:i+1], row[i+2:]))
	}
	return isSafePart1(slices.Concat(row[:i-1], row[i:])) ||
		isSafePart1(slices.Concat(row[:i], row[i+1:])) ||
		isSafePart1(slices.Concat(row[:i+1], row[i+2:]))
}

func part1() {
	input := parseInput("input.txt")
	safeCount := 0
	for _, row := range input.Values {
		if isSafePart1(row) {
			safeCount++
		}
	}
	println(safeCount)
}

func isSafePart1(row []int) bool {
	diffs := make([]int, len(row)-1)
	for i := 0; i < len(row)-1; i++ {
		diffs[i] = row[i+1] - row[i]
		if diffs[i] == 0 || diffs[i] < -3 || 3 < diffs[i] {
			return false
		}
		if i == 0 {
			continue
		}
		if diffs[i]*diffs[i-1] < 0 {
			return false
		}
	}
	return true
}

func parseInput(filename string) Input {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values [][]int
	for scanner.Scan() {
		var row []int
		numbers := strings.Split(scanner.Text(), " ")
		for _, number := range numbers {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		values = append(values, row)
	}

	return Input{
		Values: values,
	}
}

type Input struct {
	Width  int
	Height int
	Values [][]int
}
