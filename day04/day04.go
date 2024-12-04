package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	letters := parseInput(filename)

	c := 0
	for i := 1; i < len(letters)-1; i++ {
		for j := 1; j < len(letters[i])-1; j++ {
			if checkAdjacent(letters, i, j) {
				c++
			}
		}
	}

	fmt.Println(c)
}

func checkAdjacent(letters [][]string, i int, j int) bool {
	word := letters[i-1][j-1]
	word += letters[i-1][j+1]
	word += letters[i][j]
	word += letters[i+1][j-1]
	word += letters[i+1][j+1]

	return word == "MMASS" ||
		word == "MSAMS" ||
		word == "SSAMM" ||
		word == "SMASM"
}

func part1(filename string) {
	letters := parseInput(filename)

	c := 0
	xmasRegex := regexp.MustCompile(`XMAS`)
	for i := range letters {
		leftToRight := ""
		for j := range letters[i] {
			leftToRight += letters[i][j]
		}
		c += count(leftToRight, xmasRegex)
		c += count(reverse(leftToRight), xmasRegex)

		topToBottom := ""
		for j := range letters {
			topToBottom += letters[j][i]
		}
		c += count(topToBottom, xmasRegex)
		c += count(reverse(topToBottom), xmasRegex)

		topLeftToBottomRightVertical := ""
		topLeftToBottomRightHorizontal := ""
		for j := 0; j < len(letters)-i; j++ {
			topLeftToBottomRightVertical += letters[i+j][j]
			topLeftToBottomRightHorizontal += letters[j][i+j]
		}
		c += count(topLeftToBottomRightVertical, xmasRegex)
		c += count(reverse(topLeftToBottomRightVertical), xmasRegex)
		if i != 0 {
			c += count(topLeftToBottomRightHorizontal, xmasRegex)
			c += count(reverse(topLeftToBottomRightHorizontal), xmasRegex)
		}

		topRightToBottomLeftVertical := ""
		topRightToBottomLeftHorizontal := ""
		for j := 0; j < len(letters)-i; j++ {
			topRightToBottomLeftVertical += letters[i+j][len(letters)-1-j]
			topRightToBottomLeftHorizontal += letters[j][len(letters)-1-i-j]
		}
		c += count(topRightToBottomLeftVertical, xmasRegex)
		c += count(reverse(topRightToBottomLeftVertical), xmasRegex)
		if i != 0 {
			c += count(topRightToBottomLeftHorizontal, xmasRegex)
			c += count(reverse(topRightToBottomLeftHorizontal), xmasRegex)
		}
	}

	fmt.Println(c)
}

func reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func count(s string, r *regexp.Regexp) int {
	return len(r.FindAllString(s, -1))
}

func parseInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var letters [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var lineLetters []string
		for _, letter := range line {
			lineLetters = append(lineLetters, string(letter))
		}
		letters = append(letters, lineLetters)
	}

	return letters
}
