package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part2() {
	n1, n2 := parseInput("input.txt")

	counts := make(map[int]int)
	for i := range n2 {
		counts[n2[i]]++
	}

	sum := 0
	for i := range n1 {
		if counts[n1[i]] > 0 {
			sum += n1[i] * counts[n1[i]]
		}
	}

	fmt.Println(sum)
}

func part1() {
	n1, n2 := parseInput("input.txt")
	sort.Ints(n1)
	sort.Ints(n2)

	sumOfDiffs := 0
	for i := range n1 {
		t := n2[i] - n1[i]
		if t < 0 {
			t = -t
		}
		sumOfDiffs += t
	}

	fmt.Println(sumOfDiffs)
}

func parseInput(inputFile string) ([]int, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var numbers1 []int
	var numbers2 []int

	// regex reading two numbers separated by a tab
	regex := regexp.MustCompile(`(\d+)[^\d]*(\d+)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if len(matches) != 3 {
			panic("Invalid input")
		}
		number1, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		number2, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		numbers1 = append(numbers1, number1)
		numbers2 = append(numbers2, number2)
	}

	return numbers1, numbers2
}
