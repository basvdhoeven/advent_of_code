package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("2024/day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	// count horizontal
	countHorizontal := countXMASinStrings(lines)
	fmt.Printf("Count horizontal: %d\n", countHorizontal)

	// count vertical
	linesVertical := make([]string, len(lines[0]))
	for _, line := range lines {
		for j, letter := range line {
			linesVertical[j] += string(letter)
		}
	}
	countVertical := countXMASinStrings(linesVertical)
	fmt.Printf("Count vertical: %d\n", countVertical)

	// count diagonal direction /
	nLines := len(lines)
	nCharacter := len(lines[0])
	diagonalLines := nLines + nCharacter - 1
	linesDiagonal := make([]string, diagonalLines)  // direction /
	linesDiagonal2 := make([]string, diagonalLines) // direction \
	for i, line := range lines {
		for j, letter := range line {
			linesDiagonal[i+j] += string(letter)
			linesDiagonal2[i+(nCharacter-j-1)] += string(letter)
		}
	}
	countDiagonal := countXMASinStrings(linesDiagonal) + countXMASinStrings(linesDiagonal2)
	fmt.Printf("Count diagonal: %d\n", countDiagonal)

	fmt.Printf("Total count: %d\n", countHorizontal+countVertical+countDiagonal)

	// part 2
	letters := make([][]string, nLines)
	for i, line := range lines {
		letters[i] = make([]string, nCharacter)
		for j, letter := range line {
			letters[i][j] = string(letter)
		}
	}

	var xShapedMasCount int
	for i := 1; i < nLines-1; i++ {
		for j := 1; j < nCharacter-1; j++ {
			if letters[i][j] == "A" {
				diagonalLetters1 := letters[i-1][j-1] + letters[i+1][j+1]
				diagonalLetters2 := letters[i-1][j+1] + letters[i+1][j-1]
				if (diagonalLetters1 == "MS" || diagonalLetters1 == "SM") && (diagonalLetters2 == "MS" || diagonalLetters2 == "SM") {
					xShapedMasCount += 1
				}
			}
		}
	}

	fmt.Printf("Total x-shaped MAS count: %d\n", xShapedMasCount)
}

func countXMASinStrings(lines []string) int {
	count := 0
	for _, line := range lines {
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}
	return count
}
